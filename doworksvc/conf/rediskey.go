package conf

//用户注册 邮箱 对应的验证码
const (
	user_register_email_key = "user_register_email_key"
)

func GenUserRegisterEmailKey(email string) string {
	return user_register_email_key + "." + email
}
