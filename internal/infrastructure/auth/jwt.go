package auth

import (
	"Hermes/internal/config"
	"Hermes/internal/usecase/services"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	cnf *config.JWTConfig
}

func NewJWTService(conf config.JWTConfig) *JWTService {
	return &JWTService{
		cnf: &conf,
	}
}

type InsertClaims struct {
	services.CustomClaims
	jwt.RegisteredClaims
}

// Attempts to generate token with provided claims.
func (jwtServ *JWTService) GenerateToken(claims services.CustomClaims) (string, error) {
	expDate := time.Now().Add(jwtServ.cnf.ExpTime)

	genClaims := InsertClaims{
		CustomClaims: claims,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    jwtServ.cnf.Issuer,
			Audience:  jwt.ClaimStrings{jwtServ.cnf.Audience},
			ExpiresAt: jwt.NewNumericDate(expDate),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, genClaims)

	signedToken, err := token.SignedString(jwtServ.cnf.Secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Verifies validity of provided token.
func (jwtServ *JWTService) VerifyToken(token string) (*services.CustomClaims, error) {
	genClaims := InsertClaims{}
	resToken, err := jwt.ParseWithClaims(token, &genClaims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signinig method: %v", t.Header["alg"])
		}

		return []byte(jwtServ.cnf.Secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token session expired")
		}
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token signature")
		}

		return nil, fmt.Errorf("Token validation failed: %w", err)
	}

	if claims, ok := resToken.Claims.(*InsertClaims); ok && resToken.Valid {
		if claims.Issuer != jwtServ.cnf.Issuer {
			return nil, errors.New("invlaid token issuer")
		}
		if len(claims.Audience) == 0 || claims.Audience[0] != jwtServ.cnf.Audience {
			return nil, errors.New("invlaid token audience")
		}

		return &claims.CustomClaims, nil
	}
	return nil, errors.New("invalid token claims mapping")
}
