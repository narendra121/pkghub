package auth

type LoginFactory interface {
	AuthenticateUser(userName, password string) string
}

func NewLoginFactory(loginType interface{}) LoginFactory {
	switch loginType.(type) {
	case Login:
		return loginType.(*Login)
	default:
		return nil
	}
}

func (l *Login) AuthenticateUser(userName, password string) string {

	return ""
}
