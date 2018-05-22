// +build debug

package login

import (
	"testing"

	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/CJ-Jackson/paniko/paniko/security"
	"github.com/golang/mock/gomock"
)

func TestLoginController(t *testing.T) {
	type Mocks struct {
		userController *security.MockUserController
		loginView      *MockLoginView
	}

	let := func(t *testing.T) (Mocks, LoginController) {
		ctrl := gomock.NewController(t)
		mocks := Mocks{
			userController: security.NewMockUserController(ctrl),
			loginView:      NewMockLoginView(ctrl),
		}

		subject := LoginController{
			userController: mocks.userController,
			loginView:      mocks.loginView,
		}

		return mocks, subject
	}

	t.Run("It's fails to logs in", func(t *testing.T) {
		mocks, subject := let(t)

		form := LoginForm{
			Uri:         "Uri",
			Username:    "Username",
			UsernameErr: nil,
			Password:    "Password",
			PasswordErr: nil,
			Attempt:     true,
			Valid:       true,
		}

		mocks.userController.EXPECT().Login(nil, "Username", "Password").Return(false)
		mocks.loginView.EXPECT().LoginTemplate(nil, form).Times(1)

		subject.DoLogin(nil, form)
	})

	t.Run("It's logs in", func(t *testing.T) {
		mocks, subject := let(t)

		form := LoginForm{
			Uri:         "Uri",
			Username:    "Username",
			UsernameErr: nil,
			Password:    "Password",
			PasswordErr: nil,
			Attempt:     true,
			Valid:       true,
		}

		mocks.userController.EXPECT().Login(nil, "Username", "Password").Return(true)

		defer func() {
			r := recover().(common.HttpRedirectError)
			if r.Location != "Uri" {
				t.Error("Should be 'Uri'")
			}
		}()

		subject.DoLogin(nil, form)
	})
}
