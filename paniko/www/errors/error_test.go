package errors

import (
	"testing"

	"github.com/CJ-Jackson/ctx"
	"github.com/golang/mock/gomock"
)

func TestErrorController(t *testing.T) {
	type Mocks struct {
		view *MockErrorView
	}

	let := func(t *testing.T) (Mocks, ErrorController) {
		ctrl := gomock.NewController(t)
		mocks := Mocks{
			view: NewMockErrorView(ctrl),
		}

		subject := ErrorController{
			production: false,
			view:       mocks.view,
		}

		return mocks, subject
	}

	t.Run("It's in production", func(t *testing.T) {
		mocks, subject := let(t)

		mocks.view.EXPECT().
			ErrorTemplate(nil, int(200), "OK", gomock.Any()).Do(
			func(context ctx.Context, code int, title string, data ErrorTemplateData) {
				if len(data.StackTrace) > 0 {
					t.Error("Should not have stack trace")
				}
			},
		)

		(&subject).production = true
		subject.ShowError(nil, 200, "OK", "Test")
	})

	t.Run("It's not in production", func(t *testing.T) {
		mocks, subject := let(t)

		mocks.view.EXPECT().
			ErrorTemplate(nil, int(200), "OK", gomock.Any()).Do(
			func(context ctx.Context, code int, title string, data ErrorTemplateData) {
				if len(data.StackTrace) == 0 {
					t.Error("Should have stack trace")
				}
			},
		)

		(&subject).production = false
		subject.ShowError(nil, 200, "OK", "Test")
	})
}
