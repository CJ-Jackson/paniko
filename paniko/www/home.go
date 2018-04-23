package www

import (
	"time"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/expiration"
)

type HomeController struct {
	homeView   HomeView
	expiration expiration.Expiration
}

func NewHomeController(context ctx.BackgroundContext) HomeController {

	return HomeController{
		homeView:   NewHomeView(context),
		expiration: expiration.GetExpiration(context),
	}
}

func (c HomeController) Index(context ctx.Context) {
	c.homeView.Index(context, HomeViewIndexData{
		Expiry: c.expiration.GetTime().Format(time.RFC1123),
	})
}

func (c HomeController) IAmAlive(context ctx.Context) {
	c.expiration.Reset()
	c.homeView.Json(context, JsonData{
		Alert: "",
		When:  c.expiration.GetTime().Format(time.RFC1123),
	})
}

func (c HomeController) InDanger(context ctx.Context) {
	c.expiration.Expire()
	c.homeView.Json(context, JsonData{
		Alert: "",
		When:  c.expiration.GetTime().Format(time.RFC1123),
	})
}
