package services

import (
	"app/auth"
	"app/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "loginHandler, %q", html.EscapeString(r.URL.Path))

	var err error
	var account repository.Account

	_, claims, err := jwtauth.FromContext(r.Context())
	fmt.Printf("%+v\n", claims)

	if err == nil {
		//token, err := TokenAuth.Decode(cookie.String())
		fmt.Printf("%+v\n", claims)

		idStr, ok1 := claims["ID"]
		email, ok2 := claims["email"]
		role, ok3 := claims["role"]

		if ok1 && ok2 && ok3 {
			account.ID = int32(idStr.(float64))
			account.Email = email.(string)
			account.Role = pgtype.Text{role.(string), true}
		}
	} else {
		fmt.Println(err)
	}

	if account.ID == 0 {
		var payload LoginBody

		err = json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write(ErrInvalidJson)
			return
		}

		//fmt.Printf("%+v\n", payload)

		account, err = authRepo.GetAccountByEmail(r.Context(), payload.Email)

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
	}

	_, tokenString, err := TokenAuth.Encode(map[string]interface{}{
		"ID":    account.ID,
		"email": account.Email,
		"role":  account.Role.String,
	})

	cookie := &http.Cookie{
		Name: tokenName,
		Path: "/",
		//Domain: "",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 15),
		//Secure:   true,
		HttpOnly: true,
	}
	//fmt.Printf("%+v\n", cookie.Value)
	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	r.AddCookie(cookie)

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
		"role":  account.Role.String,
	})

	//processResponse[interface{}](w, req.status, req.body, err)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

type RegisterBody struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "registerHandler, %q", html.EscapeString(r.URL.Path))

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
		Role:     pgtype.Text{String: payload.Password, Valid: true},
		Password: hashedPassword,
	})

	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(payload)

	req, err := CallServiceWithCircuitBreaker(
		profilesCb, "POST", r.URL.String(), r.Header, &buf)

	processResponse[int](w, req.status, req.body, err)
}
