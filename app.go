package main

import (
	"context"
	"fmt"
	"io/fs"
	"time"

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
	Info  FileInfoJson `json:"FileInfoJson"`
	Name  string       `json:"name"`
	Type  string       `json:"type"`
	IsDir bool         `json:"isDir"`
}

type FileInfoJson struct {
	Name    string    `json:"name"`     // base name of the file
	Size    int64     `json:"size"`     // length in bytes for regular files; system-dependent for others
	ModTime time.Time `json:"modeTime"` // modification time
	IsDir   bool      `json:"isDir"`    // abbreviation for Mode().IsDir()
	Sys     any       `json:"Sys"`      // underlying data source (can return nil)
}

func makeJson(a fs.DirEntry) DirEntryJson {
	info, _ := a.Info()
	return DirEntryJson{
		Name:  a.Name(),
		Type:  a.Type().String(),
		IsDir: a.IsDir(),
		Info: FileInfoJson{
			Name:    info.Name(),
			Size:    info.Size(),
			ModTime: info.ModTime(),
			IsDir:   info.IsDir(),
			Sys:     info.Sys(),
		},
	}
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
