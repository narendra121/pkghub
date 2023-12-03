package auth

type Login struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginBuilder struct {
	login Login
}

func NewLoginBuilder() *LoginBuilder {
	return &LoginBuilder{}
}

func (lb *LoginBuilder) AddUserName(userName string) *LoginBuilder {
	lb.login.UserName = userName
	return lb
}

func (lb *LoginBuilder) AddPassword(password string) *LoginBuilder {
	lb.login.UserName = password
	return lb
}

func (lb *LoginBuilder) Build() Login {
	return lb.login
}
