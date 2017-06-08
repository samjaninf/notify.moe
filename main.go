package main

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/pages/airing"
	"github.com/animenotifier/notify.moe/pages/anime"
	"github.com/animenotifier/notify.moe/pages/awards"
	"github.com/animenotifier/notify.moe/pages/dashboard"
	"github.com/animenotifier/notify.moe/pages/forum"
	"github.com/animenotifier/notify.moe/pages/forums"
	"github.com/animenotifier/notify.moe/pages/genre"
	"github.com/animenotifier/notify.moe/pages/genres"
	"github.com/animenotifier/notify.moe/pages/posts"
	"github.com/animenotifier/notify.moe/pages/profile"
	"github.com/animenotifier/notify.moe/pages/search"
	"github.com/animenotifier/notify.moe/pages/threads"
	"github.com/animenotifier/notify.moe/pages/users"
)

var app = aero.New()

// APIKeys ...
type APIKeys struct {
	Google struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"google"`
}

func main() {
	// CSS
	app.SetStyle(components.CSS())

	// HTTPS
	app.Security.Load("security/fullchain.pem", "security/privkey.pem")

	// Session store
	app.Sessions.Store = arn.NewAerospikeStore("Session")

	// Session middleware
	app.Use(func(ctx *aero.Context, next func()) {
		// Handle the request first
		next()

		// Update session if it has been modified
		if ctx.HasSession() && ctx.Session().Modified() {
			app.Sessions.Store.Set(ctx.Session().ID(), ctx.Session())
		}
	})

	// Layout
	app.Layout = func(ctx *aero.Context, content string) string {
		return components.Layout(content)
	}

	// Ajax routes
	app.Ajax("/", dashboard.Get)
	app.Ajax("/anime", search.Get)
	app.Ajax("/anime/:id", anime.Get)
	app.Ajax("/genres", genres.Get)
	app.Ajax("/genres/:name", genre.Get)
	app.Ajax("/forum", forums.Get)
	app.Ajax("/forum/:tag", forum.Get)
	app.Ajax("/threads/:id", threads.Get)
	app.Ajax("/posts/:id", posts.Get)
	app.Ajax("/user/:nick", profile.Get)
	app.Ajax("/user/:nick/threads", threads.GetByUser)
	app.Ajax("/airing", airing.Get)
	app.Ajax("/users", users.Get)
	app.Ajax("/awards", awards.Get)

	EnableGoogleLogin(app)

	app.Get("/images/cover/:file", func(ctx *aero.Context) string {
		format := ".jpg"

		if ctx.CanUseWebP() {
			format = ".webp"
		}

		return ctx.File("images/cover/" + ctx.Get("file") + format)
	})

	// Favicon
	app.Get("/favicon.ico", func(ctx *aero.Context) string {
		return ctx.File("images/icons/favicon.ico")
	})

	// Web manifest
	app.Get("/manifest.json", func(ctx *aero.Context) string {
		return ctx.JSON(app.Config.Manifest)
	})

	// Scripts
	app.Get("/scripts.js", func(ctx *aero.Context) string {
		return ctx.File("temp/scripts.js")
	})

	// For benchmarks
	app.Get("/hello", func(ctx *aero.Context) string {
		return ctx.Text("Hello World")
	})

	// Let's go
	app.Run()
}