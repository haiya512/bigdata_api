package controllers

import (
	// "crypto"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/bigdata_api/models"
	"github.com/bigdata_api/utils"
	// "log"
	"strings"
	"time"
)

var (
	ErrContentBlank     = ErrResponse{422001, "微信朋友圈内容为空,图片也视为空"}
	ErrNicknameIsBlank  = ErrResponse{422002, "微信昵称不可为空"}
	ErrWxid				= ErrResponse{422003, "微信ID不可为空"}
)

type MessageController struct {
	BaseController
}

// @Title 添加新message
// @Description 添加微信朋友圈
// @Param	content		formData 	string	true 		"用户手机号"
// @Param	nickname	formData 	string	true		"微信昵称"
// @Param	wxid		formData 	string	true		"微信ID"
// @Success 200 {object}
// @Failure 403 参数错误：缺失或格式错误
// @Faulure 422 已被注册
// @router /reg [post]
func (this *MessageController) AddNewMessage() {
	content  := this.GetString("content")
	nickname := this.GetString("nickname")
	wxid 	 := this.GetString("wxid")
	pic_urls := this.GetString("pic_urls")

	valid := validation.Validation{}
	//表单验证
	//valid.Required(content, "content").Message("朋友圈内容不可为空")
	valid.Required(nickname, "nickname").Message("用户昵称不可为空")
	valid.Required(wxid, "wxid").Message("wxid不可为空")
	//valid.MinSize(nickname, 2, "nickname").Message("用户名最小长度为 2")
	//valid.MaxSize(nickname, 40, "nickname").Message("用户名最大长度为 40")
	//valid.Length(password, 32, "password").Message("密码格式不对")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			this.Ctx.ResponseWriter.WriteHeader(403)
			this.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
			this.ServeJSON()
			return
		}
	}
	//if models.CheckMessageContent(content) {
	//	this.Ctx.ResponseWriter.WriteHeader(422)
	//	this.Data["json"] = ErrContentBlank
	//	this.ServeJSON()
	//	return
	//}
	if models.CheckNickname(nickname) {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = ErrNicknameIsBlank
		this.ServeJSON()
		return
	}

	if models.CheckWxid(wxid) {
		this.Ctx.ResponseWriter.WriteHeader(422)
		this.Data["json"] = ErrWxid
		this.ServeJSON()
		return
	}

	mess := models.Message{
		Content:    content,
		Nickname:	nickname,
		Wxid: 		wxid,
		PicUrls: 	pic_urls,
	}
	this.Data["json"] = Response{0, "success.", models.MessageAdd(mess)}
	this.ServeJSON()

}

// @Title 登录
// @Description 账号登录
// @Success 200 {object}
// @Failure 404 no enough input
// @Failure 401 No Admin
// @router /login [post]
func (this *UserController) Login() {
	nickname := this.GetString("nickname")
	password := this.GetString("password")

	user, ok := models.CheckUserAuth(nickname, password)
	if !ok {
		this.Data["json"] = ErrNicknameOrPasswd
		this.ServeJSON()
		return
	}

	et := utils.EasyToken{
		Username: user.Nickname,
		Uid:      user.Id,
		Expires:  time.Now().Unix() + 3600,
	}

	token, err := et.GetToken()
	if token == "" || err != nil {
		this.Data["json"] = ErrResponse{-0, err}
	} else {
		this.Data["json"] = Response{0, "success.", LoginToken{user, token}}
	}

	this.ServeJSON()
}

// @Title 认证测试
// @Description 测试错误码
// @Success 200 {object}
// @Failure 401 unauthorized
// @router /auth [get]
func (this *UserController) Auth() {
	et := utils.EasyToken{}
	authtoken := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))
	valido, err := et.ValidateToken(authtoken)
	if !valido {
		this.Ctx.ResponseWriter.WriteHeader(401)
		this.Data["json"] = ErrResponse{-1, fmt.Sprintf("%s", err)}
		this.ServeJSON()
		return
	}

	this.Data["json"] = Response{0, "success.", "is login"}
	this.ServeJSON()
}
