package constant

type ErrorBase struct {
	msg string
}

func (e *ErrorBase) Error() string {
	return e.msg
}

var (
	HashErr            = &ErrorBase{"password hash err"}
	CaptchaErr         = &ErrorBase{"验证码错误"}
	EmailExistErr      = &ErrorBase{"邮箱已注册"}
	UserNotMatchGroup  = &ErrorBase{"用户不在群组内"}
	UnLoginErr         = &ErrorBase{"登录失败"}
	UnLoginPwdErr      = &ErrorBase{"密码错误"}
	TokenLayoutErr     = &ErrorBase{"Token格式错误"}
	NotGroupRecordErr  = &ErrorBase{"NoGroupFound"}
	GroupGetMembersErr = &ErrorBase{"GroupGetMembers"}
	TokenAuthErr       = &ErrorBase{"token鉴权失败"}
)
