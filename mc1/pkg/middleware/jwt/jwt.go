package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	algorithm       jwt.SigningMethod
	duration        time.Duration
	publicRoutesMap map[string]bool
	secretKey       []byte
}

func New(signingKeyEnv string, algorithm string, duration int) (*JWTService, error) {
	signingMethod := jwt.GetSigningMethod(algorithm)
	if signingMethod == nil {
		return nil, fmt.Errorf("%s is not recognized as a valid jwt signing method", algorithm)
	}

	secretKey := os.Getenv(signingKeyEnv)
	if secretKey == "" {
		return nil, fmt.Errorf("env variable %s not defined", signingKeyEnv)
	}

	return &JWTService{
		algorithm:       signingMethod,
		duration:        time.Duration(duration) * time.Minute,
		publicRoutesMap: publicRoutesMap(),
		secretKey:       []byte(secretKey),
	}, nil
}

// GenerateToken generates new JWT token
func (j *JWTService) GenerateToken() (string, string, error) {
	expire := time.Now().Add(j.duration)

	token := jwt.NewWithClaims((j.algorithm), jwt.MapClaims{
		"authorized": true,
		"user":       "Pratik",
		"exp":        expire.Unix(),
	})

	tokenString, err := token.SignedString(j.secretKey)

	return tokenString, expire.Format(time.RFC3339), fmt.Errorf("JWT generation failed: %s", err)
}

// MWFunc makes JWT implement the Middleware interface.
func (j *JWTService) MWFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isPublic := j.publicRoutesMap[r.RequestURI]; isPublic {
			h.ServeHTTP(w, r)
			return
		}

		token, err := j.ParseToken(r)
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "Not Authorized\n")
			return
		}

		h.ServeHTTP(w, r)
	})
}

// ParseToken parses token from Authorization header
func (j *JWTService) ParseToken(r *http.Request) (*jwt.Token, error) {
	headerTokenValue := r.Header["Authorization"]

	if headerTokenValue == nil {
		return nil, errors.New("authorization header is missing")
	}

	token, err := jwt.Parse(headerTokenValue[0], j.tokenParserFunc)
	if err != nil {
		return nil, errors.New("unable to parse token")
	}

	return token, nil
}

func (j *JWTService) tokenParserFunc(token *jwt.Token) (interface{}, error) {
	if j.algorithm != token.Method {
		return nil, fmt.Errorf("There was an error")
	}

	return j.secretKey, nil
}

func publicRoutesMap() map[string]bool {
	return map[string]bool{
		"/v1/login/":        true,
		"/v1/registration/": true,
	}
}
