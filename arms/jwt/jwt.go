package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type CustomClaims struct {
	UserId     uint64
	BufferTime int64
	jwt.RegisteredClaims
}

var (
	// std is the name of the standard
	stdJwt           = New()
	TokenExpired     = errors.New("Token is expired ")
	TokenNotValidYet = errors.New("Token not active yet ")
	TokenMalformed   = errors.New("That's not even a token ")
	TokenInvalid     = errors.New("Couldn't handle this token: ")
)

func New() *JWT {
	return newJWT()
}

func Std() *JWT {
	return stdJwt
}

type JWT struct {
	SigningKey []byte
}

func newJWT() *JWT {
	return &JWT{
		[]byte("mq+ZeGafL+b1xdC0u9vSVg=="),
	}
}

func VerifyTokenWithFresh(tokenStr string) (userId uint64, newToken string, err error) {
	claims, err := Std().ParseToken(tokenStr)
	if err != nil {
		return 0, "", err
	}

	if !claims.VerifyExpiresAt(time.Now().Add(time.Second*86400*1), true) {
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(expiresTime))
		tokenStr, _ = Std().CreateToken(*claims)
	}
	return claims.UserId, tokenStr, err
}

func VerifyToken(tokenStr string) (userId uint64, err error) {
	claims, err := Std().ParseToken(tokenStr)
	if err != nil {
		return 0, err
	}
	return claims.UserId, err
}

var expiresTime = time.Second * 86400 * 7

func CreateNewToken(userId uint64, expireTime time.Duration) (string, error) {
	cc := CustomClaims{
		UserId:     userId,
		BufferTime: 86400,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-10)),        // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)), // 过期时间 7天  配置文件
			Issuer:    "thh",                                          // 签名的发行者
		},
	}
	return Std().CreateToken(cc)
}

// CreateToken 创建一个token
func (itself *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(itself.SigningKey)
}

// ParseToken 解析 token
func (itself *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
		func(token *jwt.Token) (i any, e error) {
			return itself.SigningKey, nil
		},
	)

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
		} else {
			return nil, err
		}
	}
	return nil, TokenInvalid
}
