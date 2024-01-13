package Service

import (
	"HiringSystem/Utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	secretKey = []byte("your_secret_key") // 替换为实际的密钥
)

func createToken(Id uint, UserType string) (string, error) {
	// 设置JWT过期时间为10分钟
	expirationTime := time.Now().Add(10 * time.Minute)

	// 创建JWT令牌，将用户信息加入负载
	claims := Utils.Claims{
		UserType: UserType,
		UserId:   Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
