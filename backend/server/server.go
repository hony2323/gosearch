package server

import (
	"io/fs"

	backendPKG "github.com/hony2323/gosearch/backend"
	pb "github.com/hony2323/gosearch/backend/types/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnimplementedItemsResultsServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetResults(req *pb.ItemsRequest, stream pb.ItemsResults_GetResultsServer) error {
	dirs, err := backendPKG.ListFolder("C:/Program Files")
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		protoDir := ConvertDirEnrty(dir)
		response := pb.ItemsResponse{
			Entry: &protoDir,
		}
		err = stream.Send(&response)
		if err != nil {
			return err
		}
	}

	return nil
}

func ConvertInfo(fileInfo fs.FileInfo) *pb.FileInfo {
	modTime := timestamppb.New(fileInfo.ModTime())
	return &pb.FileInfo{
		Name:    fileInfo.Name(),
		Size:    fileInfo.Size(),
		ModTime: modTime,
		IsDir:   fileInfo.IsDir(),
	}
}

func ConvertDirEnrty(dirEntry fs.DirEntry) pb.DirEntry {
	info, _ := dirEntry.Info()
	return pb.DirEntry{
		Info:  ConvertInfo(info),
		Name:  dirEntry.Name(),
		Type:  dirEntry.Type().String(),
		IsDir: dirEntry.IsDir(),
	}
}
