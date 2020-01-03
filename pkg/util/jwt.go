package util

import (
    "time"

    jwt "github.com/dgrijalva/jwt-go"

    "rwplus-api/pkg/setting"
)

var rwplusSecret =  []byte(setting.RwplusSecret)

type Claims struct {
    Name string `json:"name"`
    Password string `json:"password"`
    jwt.StandardClaims
}

func GenerateToken(name, password string) (string, error) {
    nowTime := time.Now()
    expireTime := nowTime.Add(3 * time.Hour)

    claims := Claims {
        name,
        password,
        jwt.StandardClaims {
            ExpiresAt : expireTime.Unix(),
            Issuer : "rwplus-api",
        },
    }

    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := tokenClaims.SignedString(rwplusSecret)

    return token, err
}

func ParseToken(token string) (*Claims, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return rwplusSecret, nil
    })

    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }

    return nil, err

}