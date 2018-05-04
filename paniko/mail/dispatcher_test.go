// +build debug

package mail

import (
	"testing"

	"github.com/CJ-Jackson/paniko/paniko/expiration"
	"github.com/golang/mock/gomock"
)

func TestDispatcher(t *testing.T) {
	type Mocks struct {
		mailer     *MockMailer
		expiration *expiration.MockExpiration
	}

	let := func(t *testing.T) (Mocks, Dispatcher) {
		ctrl := gomock.NewController(t)

		mocks := Mocks{
			mailer:     NewMockMailer(ctrl),
			expiration: expiration.NewMockExpiration(ctrl),
		}

		subject := Dispatcher{
			mailer:     mocks.mailer,
			expiration: mocks.expiration,
		}

		return mocks, subject
	}

	t.Run("It did not expire therefore did not call dispatch", func(t *testing.T) {
		mocks, subject := let(t)

		mocks.expiration.EXPECT().Expired().Return(false)

		subject.CheckExpirationAndDispatch()
	})

	t.Run("It did expire therefore did call dispatch", func(t *testing.T) {
		mocks, subject := let(t)

		mocks.expiration.EXPECT().Expired().Return(true)
		mocks.mailer.EXPECT().Dispatch().Times(1)

		subject.CheckExpirationAndDispatch()
	})
}
