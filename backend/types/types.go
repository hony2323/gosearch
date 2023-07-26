package types

import "time"

type DirEntry struct {
	Info  FileInfo `json:"Info"`
	Name  string   `json:"name"`
	Type  string   `json:"type"`
	IsDir bool     `json:"isDir"`
}

type FileInfo struct {
	Name    string    `json:"name"`     // base name of the file
	Size    int64     `json:"size"`     // length in bytes for regular files; system-dependent for others
	ModTime time.Time `json:"modeTime"` // modification time
	IsDir   bool      `json:"isDir"`    // abbreviation for Mode().IsDir()
	Sys     any       `json:"Sys"`      // underlying data source (can return nil)
}
