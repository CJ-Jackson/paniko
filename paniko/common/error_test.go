// +build debug

package common

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestErrorService(t *testing.T) {
	type Mocks struct {
		log *MockLogger
	}

	let := func(t *testing.T) (Mocks, errorService) {
		ctrl := gomock.NewController(t)
		mocks := Mocks{
			log: NewMockLogger(ctrl),
		}

		subject := errorService{
			log: mocks.log,
		}

		return mocks, subject
	}

	t.Run("Does not panic", func(t *testing.T) {
		_, subject := let(t)

		subject.CheckErrorAndPanic(nil)
	})

	t.Run("Does panic", func(t *testing.T) {
		mocks, subject := let(t)

		err := errors.New("I am error")
		mocks.log.EXPECT().Panic(err).Times(1)

		subject.CheckErrorAndPanic(err)
	})

	t.Run("Does not log", func(t *testing.T) {
		_, subject := let(t)

		subject.CheckErrorAndLog(nil)
	})

	t.Run("Does log", func(t *testing.T) {
		mocks, subject := let(t)

		err := errors.New("I am error")
		mocks.log.EXPECT().Print(err).Times(1)

		subject.CheckErrorAndLog(err)
	})
}
