//go:generate mockgen -write_package_comment=false -package=expiration -source=expiration.go -destination=expiration.mock.go
//go:generate debugflag expiration.mock.go

package expiration

import (
	"sync"
	"time"

	"github.com/CJ-Jackson/paniko/paniko/common"
	"github.com/cjtoolkit/ctx"
)

type Expiration interface {
	GetTime() time.Time
	Expired() bool
	Reset()
	Expire()
}

func GetExpiration(context ctx.BackgroundContext) Expiration {
	return context.Persist("expiration-fa22848f098081e710887c8a2b930f07", func() (interface{}, error) {
		daysToExpiry := common.GetConfig(context).DaysToExpiry

		e := &expiration{
			daysToExpiry: daysToExpiry,
			t:            time.Now().AddDate(0, 0, daysToExpiry),
		}

		return e, nil
	}).(Expiration)
}

type expiration struct {
	sync.RWMutex
	daysToExpiry int
	t            time.Time
}

func (e *expiration) GetTime() time.Time {
	e.RLock()
	defer e.RUnlock()

	return e.t
}

func (e *expiration) Expired() bool {
	e.RLock()
	defer e.RUnlock()

	return time.Now().After(e.t)
}

func (e *expiration) Reset() {
	e.Lock()
	defer e.Unlock()

	e.t = time.Now().AddDate(0, 0, e.daysToExpiry)
}

func (e *expiration) Expire() {
	e.Lock()
	defer e.Unlock()

	e.t = time.Now()
}
