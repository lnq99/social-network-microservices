package services

import (
	"app/auth"
	"app/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"html"
	"net/http"
	"strconv"
	"time"
)

const tokenName = "jwt"

var (
	authRepo  repository.Repo
	TokenAuth *jwtauth.JWTAuth
)

func NewAuthService(repo repository.Repo) Service {
	authRepo = repo

	service := Service{
		Info: ServiceInfo{
			Name: "Auth",
			Addr: "",
			Path: "",
		},
		Endpoints: []Endpoint{
			{"POST", "login", loginHandler},
			{"POST", "register", registerHandler},
			{"DELETE", "logout", logoutHandler},
		},
	}
	return service
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//type loginResponse struct {
//	Token string          `json:"token"`
//	User  ProfileResponse `json:"user"`
//}
//
//type ProfileResponse struct {
//	Id        int    `json:"id"`
//	Name      string `json:"name"`
//	Gender    string `json:"gender"`
//	Birthdate string `json:"birthdate"`
//	Created   string `json:"created"`
//	Intro     string `json:"intro"`
//	AvatarS   string `json:"avatars"`
//	AvatarL   string `json:"avatarl"`
//	//PostCount  string `json:"postCount"`
//	//PhotoCount string `json:"photoCount"`
//}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "loginHandler, %q", html.EscapeString(r.URL.Path))

	var payload LoginBody
	var err error

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(ErrInvalidJson)
		return
	}

	//fmt.Printf("%+v\n", payload)

	account, err := authRepo.GetAccountByEmail(r.Context(), payload.Email)
	fmt.Printf("%+v\n", account)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(ErrWrongEmailOrPassword)
		return
	} else if account.Email != payload.Email {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = auth.CheckPassword(payload.Password, account.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(ErrWrongEmailOrPassword)
		return
	}

	_, tokenString, err := TokenAuth.Encode(map[string]interface{}{
		"ID":    account.ID,
		"email": account.Email,
		"role":  account.Role.String,
	})

	cookie := http.Cookie{
		Name: tokenName,
		Path: "/",
		//Domain: "",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 15),
		//Secure:   true,
		HttpOnly: true,
	}
	//fmt.Printf("%+v\n", cookie.Value)
	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	res, err := CallServiceWithCircuitBreaker(
		profilesCb, "GET",
		ProfilesServiceAddr+"/api/v1/profile/"+strconv.FormatInt(int64(account.ID), 10),
		r.Header, nil)

	var reqBody interface{}
	_ = json.NewDecoder(res.body).Decode(&reqBody)
	_ = r.Body.Close()

	w.WriteHeader(res.status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": tokenString,
		"user":  reqBody,
	})

	//processResponse[interface{}](w, req.status, req.body, err)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logoutHandler, %q", html.EscapeString(r.URL.Path))

	cookie := &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

type RegisterBody struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "registerHandler, %q", html.EscapeString(r.URL.Path))

	var payload RegisterBody
	var err error

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(ErrInvalidJson)
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	authRepo.CreateAccount(r.Context(), repository.CreateAccountParams{
		Email:    payload.Email,
		Role:     pgtype.Text{String: payload.Password},
		Password: hashedPassword,
	})

	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(payload)

	req, err := CallServiceWithCircuitBreaker(
		profilesCb, "POST", r.URL.String(), r.Header, &buf)

	processResponse[int](w, req.status, req.body, err)
}

//func Authenticator(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		token, _, err := FromContext(r.Context())
//
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusUnauthorized)
//			return
//		}
//
//		if token == nil || jwt.Validate(token) != nil {
//			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
//			return
//		}
//
//		// Token is authenticated, pass it through
//		next.ServeHTTP(w, r)
//	})
//}

//func AuthMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		articleID := chi.URLParam(r, "articleID")
//		article, err := dbGetArticle(articleID)
//		if err != nil {
//			http.Error(w, http.StatusText(404), 404)
//			return
//		}
//		ctx := context.WithValue(r.Context(), "article", article)
//		next.ServeHTTP(w, r.WithContext(ctx))
//
//		id := 0
//
//		token, err := r.Cookie("token")
//		if err == nil {
//			id, err = a.ParseTokenId(token.String())
//			if err == nil {
//				ctx := context.WithValue(r.Context(), "ID", id)
//				next.ServeHTTP(w, r.WithContext(ctx))
//				// log.Println(id, token, err)
//				return
//			}
//		}
//
//		err = a.TokenValid(c.Request)
//		// log.Println(err)
//
//		if err != nil {
//			c.AbortWithStatus(http.StatusUnauthorized)
//		}
//	})
//}

//func AuthMiddleware(a *auth.Manager) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		id := 0
//
//		token, err := r.Cookie("token")
//		if err == nil {
//			id, err = a.ParseTokenId(token.String())
//			if err == nil {
//				http.SetCookie(w)
//				c.Set("ID", id)
//				// log.Println(id, token, err)
//				return
//			}
//		}
//
//		err = a.TokenValid(c.Request)
//		// log.Println(err)
//
//		if err != nil {
//			c.AbortWithStatus(http.StatusUnauthorized)
//		}
//	}
//}

//func ArticleCtx(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		articleID := chi.URLParam(r, "articleID")
//		article, err := dbGetArticle(articleID)
//		if err != nil {
//			http.Error(w, http.StatusText(404), 404)
//			return
//		}
//		ctx := context.WithValue(r.Context(), "article", article)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}
