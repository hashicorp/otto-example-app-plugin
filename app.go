package main

import (
	"github.com/hashicorp/otto/app"
	"github.com/hashicorp/otto/appfile/detect"
	"github.com/hashicorp/otto/appfile"
)

// AppFactory implements plugin.AppFunc for use to serve the plugin.
func AppFactory() app.App {
	return &App{}
}

// Meta is the metadata for this app type.
var Meta = &app.Meta{
	Tuples: []app.Tuple{
		{"example", "*", "*"},
	},
	Detectors: []*detect.Detector{
		&detect.Detector{
			Type: "example",
			File: []string{"otto.example"},
		},
	},
}

// App implements app.App which is the interface necessary for an app type.
//
// The best example of real implementations is in Otto itself. The "builtin"
// folder contains App implementations that are built-in to Otto. Otto
// uses the same API as plugins, so you can see how the built-in app types
// do things and build your plugin based on that.
type App struct{}

func (a *App) Meta() (*app.Meta, error) {
	return Meta, nil
}

func (a *App) Implicit(ctx *app.Context) (*appfile.File, error) {
	return nil, nil
}

func (a *App) Compile(ctx *app.Context) (*app.CompileResult, error) {
	return nil, nil
}

func (a *App) Build(ctx *app.Context) error {
	return nil
}

func (a *App) Deploy(ctx *app.Context) error {
	return nil
}

func (a *App) Dev(ctx *app.Context) error {
	return nil
}

func (a *App) DevDep(dst, src *app.Context) (*app.DevDep, error) {
	return nil, nil
}
