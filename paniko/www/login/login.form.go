package login

type LoginForm struct {
	Uri         string
	Username    string
	UsernameErr error
	Password    string
	PasswordErr error
	Attempt     bool
	Valid       bool
}
