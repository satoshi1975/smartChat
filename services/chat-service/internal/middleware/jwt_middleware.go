package middleware

import (
	"context"
	"github.com/satoshi1975/smartChat/common/auth"
	"net/http"

	"strings"

	"github.com/julienschmidt/httprouter"
)

type JWTMiddleware struct {
	jwtService auth.JWTService
}

func NewJWTMiddleware(jwtService auth.JWTService) *JWTMiddleware {
	return &JWTMiddleware{jwtService: jwtService}
}

func (mw *JWTMiddleware) RequireAuth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := mw.jwtService.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Извлекаем базовую информацию о пользователе из токена и добавляем в контекст запроса
		claims, ok := token.Claims.(*auth.CustomClaims)
		if !ok || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Добавляем информацию о пользователе в контекст запроса
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		next(w, r.WithContext(ctx), ps)
	}
}
