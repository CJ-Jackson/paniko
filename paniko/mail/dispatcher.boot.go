package mail

import "time"

func bootDispatcher(dispatcher Dispatcher) {
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			dispatcher.CheckExpirationAndDispatch()
		}
	}()
}
