package auth

import (
	"base_app/common/authentication"
	"base_app/config"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type ContextKey int

const (
	Claims ContextKey = iota
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newRelicTxn := newrelic.FromContext(r.Context())
		defer newRelicTxn.End()

		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			next.ServeHTTP(w, r.WithContext(r.Context()))
		} else {
			jwtToken := authHeader
			function := func(token *jwt.Token) (interface{}, error) {
				signingMethod1, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				if signingMethod1.Name == "HS512" {
					return config.AuthSecretV2(), nil
				}
				return []byte(config.AuthSecret()), nil
			}
			token, err := jwt.Parse(jwtToken, function, jwt.WithoutClaimsValidation())

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized, please provide valid auth token to continue"))
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), Claims, claims)
				// Access context values in handlers like this
				// claims, _ := r.Context().Value(auth.Claims).(jwt.MapClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized, please provide valid auth token to continue"))
			}
		}
	})
}

func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newRelicTxn := newrelic.FromContext(r.Context())
		defer newRelicTxn.End()

		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			next.ServeHTTP(w, r.WithContext(r.Context()))
			return
		}
		if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			authHeader = strings.TrimSpace(authHeader[7:]) // Remove "Bearer " prefix
		}

		ctx := r.Context()
		userDetail, err := GetAdminUserIdAndRequestIdFromContext(authHeader)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized, please provide valid auth token to continue"))
			return
		}

		ctx = context.WithValue(ctx, "userDetail", userDetail)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetAdminUserIdAndRequestIdFromContext(authHeader string) (*authentication.UserDetail, error) {
	userDetail, err := getAdminPanelUser(&authHeader)
	if err != nil {
		return nil, err
	}

	return userDetail, nil
}

func getAdminPanelUser(token *string) (*authentication.UserDetail, error) {
	if token == nil {
		return nil, errors.New("Authorization token is missing")
	}

	biFrostAuthKeys := config.GetBifrostAuth()
	userDetail, err := GetAuthenticatedUserDetails(*token, biFrostAuthKeys.Key, biFrostAuthKeys.NewKey)
	if err != nil {
		return nil, err
	}

	return userDetail, nil
}

func GetAuthenticatedUserDetails(tokenString, bifrostSecretKey, bifrostSecretKeyNew string) (*authentication.UserDetail, error) {
	decodedSecretKey, err := base64.StdEncoding.DecodeString(bifrostSecretKey)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return decodedSecretKey, nil
	})

	if err != nil {
		decodedSecretKeyNew, err2 := base64.StdEncoding.DecodeString(bifrostSecretKeyNew)
		if err2 != nil {
			return nil, fmt.Errorf("error decoding fallback secret key: %w", err2)
		}
		claims = jwt.MapClaims{}
		_, err2 = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return decodedSecretKeyNew, nil
		})
		if err2 != nil {
			return nil, errors.New(authentication.UnauthorisedErrorMessage)
		}
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return nil, errors.New(authentication.UnauthorisedErrorMessage)
	}
	emailId, ok := claims["emailId"].(string)
	if !ok {
		return nil, errors.New(authentication.UnauthorisedErrorMessage)
	}

	return &authentication.UserDetail{
		UserId:  userId,
		EmailId: emailId,
	}, nil
}
