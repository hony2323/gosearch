package server

import (
	"io/fs"

	backendPKG "github.com/hony2323/gosearch/backend"
	pb "github.com/hony2323/gosearch/backend/types/protos"
)

type Server struct{}

func (s *Server) GetResults(req *pb.ItemsRequest, stream pb.ItemsResponse) error {
	dirs, err := backendPKG.ListFolder("C:/Program Files")
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		protoDir = pb.DirEntry{
			Info: dir.Info(),
		}
	}

}

func ConvertInfo(fileInfo fs.FileInfo) pb.FileInfo {
	return pb.FileInfo{
		Name:    fileInfo.Name(),
		Size:    fileInfo.Size(),
		ModTime: fileInfo.ModTime(),
		IsDir:   fileInfo.IsDir(),
	}
}
