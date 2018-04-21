package expiration

import (
	"sync"
	"time"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/config"
)

type Expiration interface {
	GetTime() time.Time
	Expired() bool
	Reset()
	Expire()
}

func GetExpiration(context ctx.BackgroundContext) Expiration {
	name := "expiration-fa22848f098081e710887c8a2b930f07"
	if e, ok := context.Ctx(name).(Expiration); ok {
		return e
	}

	daysToExpiry := config.GetConfig(context).DaysToExpiry

	e := &expiration{
		daysToExpiry: daysToExpiry,
		t:            time.Now().AddDate(0, 0, daysToExpiry),
	}

	context.SetCtx(name, e)
	return e
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
