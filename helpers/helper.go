package helpers

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)




func ConvertByteToString(value []uint8) string{
	str := string(value)
	return str
}




type SignedDetails struct {
	Email      string
	Uid        interface{}
	jwt.StandardClaims
}

func GenerateToken(email string, uid interface{}) (signedToken string,  err error) {
	claims := &SignedDetails{
		Email:      email,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, err
}

func VerifyToken(signedToken string) (claims *SignedDetails, msg string){
	token, err := jwt.ParseWithClaims(  
        signedToken,
        &SignedDetails{},
        func(token *jwt.Token) (interface{}, error) {
            return []byte("secret"), nil
        },
    )

    if err != nil {
        msg = err.Error()
        return
    }

	claims, ok := token.Claims.(*SignedDetails)
    if !ok {
        msg = fmt.Sprintf("the token is invalid")
        msg = err.Error()
        return
    }

    if claims.ExpiresAt < time.Now().Local().Unix() {
        msg = fmt.Sprintf("token is expired")
        msg = err.Error()
        return
    }

    return claims, msg
}