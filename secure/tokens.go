// Package secure
package secure

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AccessToken struct {
	ID  uuid.UUID
	Sub uuid.UUID
	Iat int64
	Exp int64
	Ref uuid.UUID // id refreshToken
}

type RefreshToken struct {
	ID        uuid.UUID
	IP        string
	UserAgent string
}

// GenerateAccessToken предоставляет валидацию токена с учетом возможного протухания exp.
func GenerateAccessToken(guid uuid.UUID, refID uuid.UUID) string {
	key := os.Getenv("API_SECRET")
	expiredTime, _ := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_EXPIRED"), 10, 0)
	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":  uuid.New(),
		"sub": guid,
		"iat": now.Unix(),
		"exp": now.Add(time.Duration(expiredTime) * time.Second).Unix(),
		"ref": refID,
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		panic(err)
	}
	return tokenString
}

// DecodeAccessToken декодирует токен в структуру AccessToken
func DecodeAccessToken(secret, tokenString string) (*AccessToken, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Alg()}),
		jwt.WithoutClaimsValidation())
	if err != nil {
		return nil, fmt.Errorf("can't parse token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("claim not valid")
	}
	accessToken, err := parseClaims(claims)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

// parseClaims парсит calims в структуру AccessToken
func parseClaims(claims jwt.MapClaims) (*AccessToken, error) {
	idStr, ok := claims["id"].(string)
	if !ok {
		return nil, fmt.Errorf("id token not valid")
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, fmt.Errorf("id token not valid")
	}
	subStr, ok := claims["sub"].(string)
	if !ok {
		return nil, fmt.Errorf("sub token not valid")
	}
	sub, err := uuid.Parse(subStr)
	if err != nil {
		return nil, fmt.Errorf("sub token not valid")
	}
	var iat int64
	switch v := claims["iat"].(type) {
	case float64:
		iat = int64(v)
	case int64:
		iat = v
	default:
		return nil, fmt.Errorf("iat token not valid")
	}
	var exp int64
	switch v := claims["exp"].(type) {
	case float64:
		exp = int64(v)
	case int64:
		exp = v
	default:
		return nil, fmt.Errorf("exp token not valid")
	}
	refStr, ok := claims["ref"].(string)
	if !ok {
		return nil, fmt.Errorf("ref token not valid")
	}
	ref, err := uuid.Parse(refStr)
	if err != nil {
		return nil, fmt.Errorf("ref token not valid")
	}

	return &AccessToken{id, sub, iat, exp, ref}, nil
}

// GenerateRefreshToken кодирует токен в base64
func GenerateRefreshToken(id string, ip string, userAgent string) string {
	payload := strings.Join([]string{id, ip, userAgent}, "|")
	token := []byte(payload)
	return base64.StdEncoding.EncodeToString(token)
}

// DecodeRefreshToken декодирует base64 в RefreshToken
func DecodeRefreshToken(token string) (*RefreshToken, error) {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(token)))
	base64.StdEncoding.Decode(base64Text, []byte(token))
	return validateByteRefreshToken(base64Text)
}

func validateByteRefreshToken(base64Text []byte) (*RefreshToken, error) {
	parseErr := fmt.Errorf("parse err")

	elems := strings.Split(string(base64Text), "|")
	if len(elems) != 3 {
		return nil, fmt.Errorf("refresh token not valid")
	}
	ref, err := uuid.Parse(elems[0])
	if err != nil {
		return nil, parseErr
	}
	ip := elems[1]
	userAgent := elems[2]

	return &RefreshToken{ref, ip, userAgent}, nil
}

// HashRefreshToken хеширует refresh токен в bcrypt
func HashRefreshToken(token string) string {
	if len(token) > 72 {
		token = token[:72]
	}
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		panic(err)

	}
	return string(hashedToken)
}

func ValidateAccessToken(secret, tokenString string) (*AccessToken, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Alg()}))
	if err != nil {
		return nil, fmt.Errorf("token not valid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("token claims not valid")
	}

	accessToken, err := parseClaims(claims)
	if err != nil {
		return nil, fmt.Errorf("validate error")
	}
	return accessToken, nil
}
