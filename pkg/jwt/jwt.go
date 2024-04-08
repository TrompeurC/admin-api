package jwt

import (
	"admin-api/api/entity"
	"admin-api/constant"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// 过期时间
const Token_EXPIRE_DUARTION = 24 * time.Hour

// token秘钥
var Secret = []byte("admin-api")

var (
	// 令牌不存在
	ErrAbsent = "token absent"
	// 令牌无效
	ErrInvalid = "token invalid"
)

type userStdClaims struct {
	entity.JwtAdmin
	jwt.StandardClaims
}

// GenerateTokenByAdmin 根据用户信息生成token
func GenerateTokenByAdmin(admin entity.SysAdmin) (string, error) {
	jwtAdmin := entity.JwtAdmin{
		ID:       admin.ID,
		Username: admin.Username,
		Nickname: admin.Nickname,
		Icon:     admin.Icon,
		Phone:    admin.Phone,
		Note:     admin.Note,
	}
	c := userStdClaims{
		jwtAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(Token_EXPIRE_DUARTION).Unix(),
			Issuer:    "backstage",
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// ValidateToken 解析JWT
func ValidateToken(tokenString string) (*entity.JwtAdmin, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	clams := userStdClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &clams, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	return &clams.JwtAdmin, nil
}

// 返回id
func GetAdminId(c *gin.Context) (uint, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return 0, errors.New("can't get user id")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if !ok {
		return admin.ID, nil
	}
	return 0, errors.New("can't convert to id struct")
}

// 返回用户名
func GetAdminName(c *gin.Context) (string, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return "", errors.New("can't get user name")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if !ok {
		return admin.Username, nil
	}
	return "", errors.New("can't convert to name struct")
}

// 返回用户
func GetAdmin(c *gin.Context) (*entity.JwtAdmin, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return nil, errors.New("can't get api")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin, nil
	}
	return nil, errors.New("can't convert to api struct")
}
