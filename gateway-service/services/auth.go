package services

import (
	"fmt"
	"html"
	"net/http"
)

func NewAuthService() Service {
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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "loginHandler, %q", html.EscapeString(r.URL.Path))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logoutHandler, %q", html.EscapeString(r.URL.Path))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "registerHandler, %q", html.EscapeString(r.URL.Path))
}

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
