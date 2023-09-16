package services

import (
	"context"
	"net/url"
	"os"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

var (
	jwtValidator = NewJwtValidator()
)

//func IsAuthenticated(c *fiber.Ctx) error {
//	bearerToken := c.Get(AuthHeader)
//	tokenList := strings.Split(bearerToken, "Bearer ")
//
//	token := ""
//	if len(tokenList) == 2 {
//		token = tokenList[1]
//	}
//
//	validToken, err := jwtValidator.ValidateToken(c.Context(), token)
//	if err != nil {
//		c.Status(http.StatusUnauthorized)
//		return c.JSON(ToErrResponse(err))
//	}
//
//	validatedClaims := validToken.(*validator.ValidatedClaims)
//
//	claims, ok := validatedClaims.CustomClaims.(*CustomClaims)
//
//	if !ok {
//		c.Status(http.StatusUnauthorized)
//		return c.JSON(ErrorResponse{"parse claim failed"})
//	}
//
//	c.Set(UsernameHeader, claims.Username)
//
//	return c.Next()
//}

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Username string `json:"username"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func NewJwtValidator() *validator.Validator {
	issuerURL, _ := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, _ := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)

	return jwtValidator
}
