package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/Makinde1034/budget-app/helpers"
)

func Authentication(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
		authorization := r.Header.Get("Authorization")
		fmt.Println(authorization)
		token := strings.Split(authorization," ")[1]

		claims, msg := helpers.VerifyToken(token)

		if msg != "" {
			fmt.Println("An error occured")
			errMsg := struct{
				Message string `json:"message"`
			}{"Invalid token"}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errMsg)
			
			return
		}
		userId := r.WithContext(context.WithValue(r.Context(), "id", claims.Uid))
		next.ServeHTTP(w,userId)
	})
}