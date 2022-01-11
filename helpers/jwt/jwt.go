package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	UserId     uint64
	Username   string
	BufferTime int64
	jwt.StandardClaims
}

var (
	// std is the name of the standard
	std              = New()
	TokenExpired     = errors.New("Token is expired ")
	TokenNotValidYet = errors.New("Token not active yet ")
	TokenMalformed   = errors.New("That's not even a token ")
	TokenInvalid     = errors.New("Couldn't handle this token: ")
)

func New() *JWT {
	return newJWT()
}

func UseJWT() *JWT {
	return std
}

type JWT struct {
	SigningKey []byte
}

func newJWT() *JWT {
	return &JWT{
		[]byte("mq+ZeGafL+b1xdC0u9vSVg=="),
	}
}

func CreateNewToken(userId uint64, username string) (string, error) {
	cc := CustomClaims{
		UserId:     userId,
		Username:   username,
		BufferTime: 86400,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,    // 签名生效时间
			ExpiresAt: time.Now().Unix() + 86400*7, // 过期时间 7天  配置文件
			Issuer:    "thh",                       // 签名的发行者
		},
	}
	return UseJWT().CreateToken(cc)
}

func (itself *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(itself.SigningKey)
}

// ParseToken 解析 token
func (itself *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
		// 防止bug从我做起
		func([]byte) func(token *jwt.Token) (i interface{}, e error) {
			return func(token *jwt.Token) (i interface{}, e error) {
				return itself.SigningKey, nil
			}
		}(itself.SigningKey))

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch ve.Errors {
			case jwt.ValidationErrorMalformed:
				return nil, TokenMalformed
			case jwt.ValidationErrorExpired:
				return nil, TokenExpired
			case jwt.ValidationErrorNotValidYet:
				return nil, TokenNotValidYet
			default:
				return nil, TokenInvalid
			}
		}
	}
	return nil, TokenInvalid
}
