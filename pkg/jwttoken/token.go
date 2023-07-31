package jwttoken

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

type Token struct {
	secretKey []byte
}

func New(secretKet []byte) *Token {
	return &Token{
		secretKey: secretKet,
	}
}

func (t Token) GenerateJWT(id int64, ttl time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(ttl)
	claims["authorized"] = true
	claims["user_id"] = id
	tokenString, err := token.SignedString(t.secretKey)
	if err != nil {
		return "Signing Error", err
	}

	return tokenString, nil
}

func (t Token) VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			tokenString := request.Header["Token"][0]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodECDSA)
				if !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("YOUR_SECRET_KEY"), nil // Replace "YOUR_SECRET_KEY" with your actual secret key used for signing the tokens
			})

			// Handle parsing errors
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("You're Unauthorized due to error parsing the JWT"))
				if err != nil {
					return
				}
			}

			// Check token validity and expiration
			if token.Valid {
				claims, ok := token.Claims.(jwt.MapClaims)
				if !ok {
					writer.WriteHeader(http.StatusUnauthorized)
					_, err := writer.Write([]byte("Invalid token claims"))
					if err != nil {
						return
					}
				}

				expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
				if time.Now().After(expirationTime) {
					writer.WriteHeader(http.StatusUnauthorized)
					_, err := writer.Write([]byte("Token has expired"))
					if err != nil {
						return
					}
				} else {
					endpointHandler(writer, request)
				}
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("You're Unauthorized due to invalid token"))
				if err != nil {
					return
				}
			}
		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			_, err := writer.Write([]byte("You're Unauthorized due to No token in the header"))
			if err != nil {
				return
			}
		}
	})
}

func (t Token) ExtractClaims(_ http.ResponseWriter, request *http.Request) (string, error) {
	if request.Header["Token"] != nil {
		tokenString := request.Header["Token"][0]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("there's an error with the signing method")
			}
			return t.secretKey, nil
		})
		if err != nil {
			return "Error Parsing Token: ", err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			// Claims fields
			username := claims["username"].(string)
			return username, nil
		}
	}

	return "unable to extract claims", nil
}
