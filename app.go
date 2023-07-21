package main

import (
	"context"
	"fmt"
	"io/fs"

	backend "github.com/hony2323/gosearch/backend"
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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type DirEntryJson struct {
	Name  string `json:"name"`
	Info  string `json:"info"`
	IsDir bool   `json:"isDir"`
}

func makeJson(a fs.DirEntry) DirEntryJson {
	info, _ := a.Info()
	return DirEntryJson{Name: a.Name(), Info: info.Mode().String(), IsDir: a.IsDir()}
}

func (a *App) ListFolder(dir string) []DirEntryJson {
	itemsList, err := backend.ListFolder(dir)
	if err != nil {
		return []DirEntryJson{}
	}

	list := []DirEntryJson{}
	for _, item := range itemsList {
		list = append(list, makeJson(item))
	}
	return list
}
