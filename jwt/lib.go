package jwt

import (
	"fmt"
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignWithRSA(rsaKey *rsa.PrivateKey,notBefore time.Time,ttl time.Duration,data interface{})(token string,expireAt time.Time,err error){
	expireAt = notBefore.Add(ttl)
	token,err = jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"dat":    data,
		"exp" : expireAt.Unix(),
		"iat": time.Now().Unix(),
		"nbf" : notBefore.Unix(),
	}).SignedString(rsaKey)

	return token, expireAt, err
}

func ValidateWithRSA(pubKey *rsa.PublicKey,token string)(interface{}, error){
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return pubKey, nil
	})
	if err != nil{
		return  nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	return claims["dat"], nil
}
