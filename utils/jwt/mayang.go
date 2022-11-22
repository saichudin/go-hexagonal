package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var JWT_LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

type JwtMayang struct {
	secretKey string
	ClaimsMayang
}

type ClaimsMayang struct {
	Id       string `json:"id"`
	CabangId string `json:"cabang_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewJwtMayang(merchantId string, outletId string, owner string) JwtMayang {
	claims := ClaimsMayang{
		Id:       merchantId,
		CabangId: outletId,
		Username: owner,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(JWT_LOGIN_EXPIRATION_DURATION).Unix(),
		},
	}

	return JwtMayang{
		secretKey:    os.Getenv("JWT_SECRET"),
		ClaimsMayang: claims,
	}
}

func (jwtMayang *JwtMayang) CreateTokenMerchant() (string, error) {
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		jwtMayang.ClaimsMayang,
	)

	signedtoken, err := token.SignedString([]byte(jwtMayang.secretKey))
	if err != nil {
		return "", err
	}

	return signedtoken, nil
}
