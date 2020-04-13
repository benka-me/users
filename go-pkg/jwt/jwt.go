package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/status"
	"time"
)

const (
	Jwtkey       = "com.hive-and-cell.com" //TODO !!! Change that
	TokenExpired = 34
	InvalidToken = 35
)

func CheckJwt(tokenString string) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Jwtkey), nil
	})
	if err != nil && "Token is expired" == err.Error() {
		return status.Error(TokenExpired, err.Error())
	} else if err != nil {
		return status.Error(InvalidToken, err.Error())
	}
	return nil
}

func GenerateToken(username, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":  username,
		"email": email,
		"exp":   time.Now().Add(time.Hour * time.Duration(48)).Unix(),
		"iat":   time.Now().Unix(),
	})
	return token.SignedString([]byte(Jwtkey))

}
