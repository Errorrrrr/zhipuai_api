package utils

import (
	"strings"
	"time"
	"zhipuai_api/internal/common/enum"
	"zhipuai_api/internal/common/jwt"
)

func GenToken() (string, error) {
	key := strings.Split(enum.ApiKey, ".")
	id, secret := key[0], key[1]
	payload := jwt.CustomClaims{
		ApiKey:    id,
		Exp:       time.Now().UnixNano()/int64(time.Millisecond) + int64(3600)*1000,
		Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
	}
	token, err := jwt.CreateToken([]byte(secret), payload)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetTitleFromContent(content string) string {
	if len(content) >= 13 {
		return content[:10] + "..."
	}
	return content
}
