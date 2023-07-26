package main

import (
	"context"
	"io/fs"

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

func makeJson(a fs.DirEntry) typesPKG.DirEntry {
	info, _ := a.Info()
	return typesPKG.DirEntry{
		Name:  a.Name(),
		Type:  a.Type().String(),
		IsDir: a.IsDir(),
		Info: typesPKG.FileInfo{
			Name:    info.Name(),
			Size:    info.Size(),
			ModTime: info.ModTime(),
			IsDir:   info.IsDir(),
			Sys:     info.Sys(),
		},
	}
}

func (a *App) ListFolder(dir string) []typesPKG.DirEntry {
	itemsList, err := backend.ListFolder(dir)
	if err != nil {
		return []typesPKG.DirEntry{}
	}

	list := []typesPKG.DirEntry{}
	for _, item := range itemsList {
		list = append(list, makeJson(item))
	}
	return list
}

// Use all the structs in order to bind them
