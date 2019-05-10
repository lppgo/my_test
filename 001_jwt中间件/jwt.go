package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "hlms_yeeuu"
)

//...
type CustomClaims struct {
	User      string `form:"phone"`
	Level     string `form:"type"`
	HotelId   string `form:"hotel_id"`
	HotelName string `form:"hotelName"`
	UserName  string `form:"username"`
	jwt.StandardClaims
}

//...
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("claims")
		if err != nil {
			c.Redirect(302, "/login")
			return
		}
		j := NewJWT()
		// token := t.(string)
		// claims, err := j.ParseToken(token)
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				if token, err = j.RefreshToken(token); err == nil {
					c.JSON(http.StatusOK, gin.H{"error": 0, "message": "refresh token", "token": token})
					return
				}
			}
			c.Redirect(302, "/login")
			return
		}
		c.Set("claims", claims)
		c.Next()
		return
	}
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}
func GetSignKey() string {
	return SignKey
}
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

//...
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//-------jwt验证中间件------
// ------登录成功（将claims转为toekn，并将token放入cookie）----
	//  根据登录参数生成claims
	// 	jwt := middleware.NewJWT()
	// 	token, err := jwt.CreateToken(claims)
	// 	c.SetCookie("claims", token, 80000, "/", "", false, true)
// ------在将token解析为claims------
