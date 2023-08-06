package main

import (
	"embed"
	"net"

	backendPKG "github.com/hony2323/gosearch/backend"
	"github.com/hony2323/gosearch/backend/server"
	pb "github.com/hony2323/gosearch/backend/types/protos"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"google.golang.org/grpc"
)

//go:embed all:frontend/dist
var assets embed.FS
var serverPort = 50051

func main() {

	log.Infof("Atempting to listen to port %v", serverPort)
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Infof("Server is alread listening on port %v", serverPort)
	}
	go func() {
		s := grpc.NewServer()
		pb.RegisterItemsResultsServer(s, &server.Server{})
		log.Infof("Server is running on :%v", serverPort)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	// Create an instance of the app structure
	app := NewApp()
	backend := backendPKG.NewBackend()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "gosearch",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			backend,
		},
	})

	log.Info("after wails")
	if err != nil {
		println("Error:", err.Error())
	}

}
