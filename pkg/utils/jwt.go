package utils

import (
	"errors"
	"github.com/Mutezebra/tiktok/consts"
	"github.com/Mutezebra/tiktok/pkg/log"
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	UserName                  string
	ID                        int64
	AccessToken, RefreshToken string
	jwt.StandardClaims
}

// CheckToken 用于jwt中间件里检查token
func CheckToken(aToken, rToken string) (claim *Claims, err error, count int) {
	claim = new(Claims)

	aClaims, err, aValid := ParseToken(aToken)
	if err != nil {
		log.LogrusObj.Println("atoken", err)
		return
	}
	rClaims, err, rValid := ParseToken(rToken)
	if err != nil {
		log.LogrusObj.Println("rtoken", err)
		return
	}

	if aClaims.ID != rClaims.ID || aClaims.UserName != rClaims.UserName {
		log.LogrusObj.Error("fake token")
		return
	}
	claim.ID = aClaims.ID
	claim.UserName = aClaims.UserName

	// 如果两者都没过期就不更新
	if aValid && rValid {
		claim.AccessToken = aToken
		claim.RefreshToken = rToken
		count = 0
		return
	}

	// 如果两者都过期
	if !aValid && !rValid {
		err = errors.New("token expired,please login again")
		return
	}

	// 如果a过期但是r没过期就只更新a
	if !aValid && rValid {
		claim.AccessToken, err = GenerateAccessToken(aClaims.UserName, aClaims.ID)
		claim.RefreshToken = rToken
		count = 1
		return
	}

	// 如果aToken没过期但是rToken过期，两者都更新
	claim.AccessToken, err = GenerateAccessToken(aClaims.UserName, aClaims.ID)
	if err != nil {
		return
	}
	claim.RefreshToken, err = GenerateRefreshToken(rClaims.UserName, rClaims.ID)
	if err != nil {
		return
	}
	count = 2
	return

}

// GenerateToken 登陆时签发Token
func GenerateToken(userName string, id int64) (accessToken, refreshToken string, err error) {
	accessToken, err = GenerateAccessToken(userName, id)
	if err != nil {
		return "", "", err
	}
	refreshToken, err = GenerateRefreshToken(userName, id)
	if err != nil {
		return "", "", err
	}
	return
}

// GenerateAccessToken 签发AccessToken
func GenerateAccessToken(userName string, id int64) (accessToken string, err error) {
	timeNow := time.Now()
	accessTokenExpireTime := timeNow.Add(consts.AccessTokenExpireTime).Unix()
	claims := &Claims{
		UserName: userName,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Mutezebra",
			Subject:   userName,
			ExpiresAt: accessTokenExpireTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(consts.JwtSecret))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

// GenerateRefreshToken 签发AccessToken
func GenerateRefreshToken(userName string, id int64) (refreshToken string, err error) {
	timeNow := time.Now()
	refreshTokenExpireTime := timeNow.Add(consts.RefreshTokenExpireTime).Unix()
	claims := &Claims{
		UserName: userName,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Mutezebra",
			Subject:   userName,
			ExpiresAt: refreshTokenExpireTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err = token.SignedString([]byte(consts.JwtSecret))
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

// ParseToken 解析token并判断其有没有过期
func ParseToken(token string) (*Claims, error, bool) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtSecret), nil
	})

	if err != nil {
		return nil, err, false
	}
	claims, ok := tokenClaims.Claims.(*Claims)
	if ok && tokenClaims.Valid {
		return claims, nil, IsValid(tokenClaims)
	}
	return nil, err, false
}

func IsValid(token *jwt.Token) bool {
	return token.Valid
}
