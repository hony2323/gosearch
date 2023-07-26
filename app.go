package main

import (
	"context"

	backend "github.com/hony2323/gosearch/backend"
	typesPKG "github.com/hony2323/gosearch/backend/types"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ListFolder(dir string) []typesPKG.DirEntry {
	itemsList, err := backend.ListFolder(dir)
	if err != nil {
		return []typesPKG.DirEntry{}
	}

	list := []typesPKG.DirEntry{}
	for _, item := range itemsList {
		list = append(list, typesPKG.MakeJson(item))
	}
	return list
}

// Use all the structs in order to bind them
