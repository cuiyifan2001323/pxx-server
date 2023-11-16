// @Author Bing
// @Date 2023-11-14 20:52:00
// @Desc
package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"pxx-server/common"
	"pxx-server/domain"
	"pxx-server/service"
	"strconv"
)

const (
	mobileReg = "^1[3456789][0-9]{9}$"
)

type HandleUser struct {
	svc       *service.UserService
	mobileReg *regexp.Regexp
}

func NewUserHandle(svc *service.UserService) *HandleUser {
	return &HandleUser{
		svc:       svc,
		mobileReg: regexp.MustCompile(mobileReg, regexp.None),
	}
}
func (u *HandleUser) RegisterRoutes(g *gin.Engine) {
	group := g.Group("/user")
	group.POST("/signup", u.SignUp)
}

func (u *HandleUser) SignUp(context *gin.Context) {

	params := new(domain.User)
	err := context.ShouldBind(params)
	var result common.Result
	if err != nil {
		result = common.Err(500, "参数错误")
	}
	isMobile, err := u.mobileReg.MatchString(strconv.Itoa(params.Mobile))
	if !isMobile || err != nil {
		result = common.Err(500, "参数错误")
	}
	err = u.svc.Signup(context, params)
	switch err {
	case nil:
		result = common.Success("创建成功", "成功")
	default:
		result = common.Err(500, "用户注册失败")
	}
	context.JSON(200, result)
}
