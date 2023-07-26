package backend

import (
	typesPKG "github.com/hony2323/gosearch/backend/types"
)

type Backend struct {
}

func NewBackend() *Backend {
	return &Backend{}
}

// bind all the structs with dummy functions

func (b *Backend) BindDirEntry(typesPKG.DirEntry) {}
func (b *Backend) BindFileInfo(typesPKG.FileInfo) {}
