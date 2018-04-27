// +build debug

package security

import (
	"testing"

	"net/http"

	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/golang/mock/gomock"
)

func TestUserController(t *testing.T) {
	type Mocks struct {
		password     *common.MockPassword
		errorService *common.MockErrorService
		cookieHelper *MockCookieHelper
	}

	let := func(t *testing.T) (Mocks, userController) {
		ctrl := gomock.NewController(t)
		mocks := Mocks{
			password:     common.NewMockPassword(ctrl),
			errorService: common.NewMockErrorService(ctrl),
			cookieHelper: NewMockCookieHelper(ctrl),
		}

		subject := userController{
			passwordLocation: "location",
			password:         mocks.password,
			users:            map[string]string{"abc": "abc"},
			errorService:     mocks.errorService,
			cookieHelper:     mocks.cookieHelper,
		}

		return mocks, subject
	}

	t.Run("It's check cookie, but fail to get cookie", func(t *testing.T) {
		mocks, subject := let(t)

		mocks.cookieHelper.EXPECT().Get(nil, userCookieName).Return(nil)

		if _, ok := subject.CheckCookie(nil).(Guest); !ok {
			t.Error("Should be guest")
		}
	})

	t.Run("It's check cookie, but cookie value is not properly formatted", func(t *testing.T) {
		mocks, subject := let(t)

		mocks.cookieHelper.EXPECT().Get(nil, userCookieName).Return(&http.Cookie{Value: "abc"})

		if _, ok := subject.CheckCookie(nil).(Guest); !ok {
			t.Error("Should be guest")
		}
	})

	t.Run("It's check cookie, got cookie, password does not match", func(t *testing.T) {
		mocks, subject := let(t)

		mocks.cookieHelper.EXPECT().Get(nil, userCookieName).Return(&http.Cookie{Value: "abc:def"})

		if _, ok := subject.CheckCookie(nil).(Guest); !ok {
			t.Error("Should be guest")
		}
	})

	t.Run("It's check cookie, got cookie, password does match", func(t *testing.T) {
		mocks, subject := let(t)

		mocks.cookieHelper.EXPECT().Get(nil, userCookieName).Return(&http.Cookie{Value: "abc:abc"})

		if _, ok := subject.CheckCookie(nil).(User); !ok {
			t.Error("Should be user")
		}
	})

	t.Run("It's fails to log in", func(t *testing.T) {
		mocks, subject := let(t)

		mocks.password.EXPECT().CheckPassword("def", "abc").Return(false)

		if subject.Login(nil, "abc", "def") != false {
			t.Error("Should be false")
		}
	})

	t.Run("It's logs in no problem", func(t *testing.T) {
		mocks, subject := let(t)

		mocks.password.EXPECT().CheckPassword("abc", "abc").Return(true)
		mocks.cookieHelper.EXPECT().Set(nil, gomock.Any()).Times(1)
		mocks.password.EXPECT().SaltPassword("abc").Return("abc")

		if subject.Login(nil, "abc", "abc") != true {
			t.Error("Should be true")
		}
	})
}
