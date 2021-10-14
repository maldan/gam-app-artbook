package artbook

import (
	"embed"
	"flag"
	"fmt"

	"github.com/maldan/gam-app-artbook/internal/app/artbook/api"
	"github.com/maldan/gam-app-artbook/internal/app/artbook/core"
	"github.com/maldan/go-cmhp/cmhp_s3"
	"github.com/maldan/go-rapi"
	"github.com/maldan/go-rapi/rapi_core"
	"github.com/maldan/go-rapi/rapi_rest"
	"github.com/maldan/go-rapi/rapi_vfs"
)

func Start(frontFs embed.FS) {
	// Server
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16000, "Server Port")
	_ = flag.Int("clientPort", 8080, "Client Port")

	// Data
	var dataDir = flag.String("dataDir", "db", "Data Directory")
	_ = flag.String("appId", "id", "App id")
	flag.Parse()

	// Set
	core.DataDir = *dataDir

	// Init s3
	cmhp_s3.Start(core.DataDir + "/config.json")

	// Init server
	/*restserver.Start(fmt.Sprintf("%s:%d", *host, *port), map[string]interface{}{
		"/": restserver.VirtualFs{Root: "frontend/build/", Fs: frontFs},
		"/api": map[string]interface{}{
			"work":      api.WorkApi{},
			"image":     api.ImageApi{},
			"training":  api.TrainingApi{},
			"reference": api.ReferenceApi{},
		},
		"/system": map[string]interface{}{
			"config": api.ConfigApi{},
		},
	})*/

	// Start server
	rapi.Start(rapi.Config{
		Host: fmt.Sprintf("%s:%d", *host, *port),
		Router: map[string]rapi_core.Handler{
			"/": rapi_vfs.VFSHandler{
				Root: "frontend/build",
				Fs:   frontFs,
			},
			"/api": rapi_rest.ApiHandler{
				Controller: map[string]interface{}{
					"project":    api.ProjectApi{},
					"image":      api.ImageApi{},
					"training":   api.TrainingApi{},
					"statistics": api.StatisticsApi{},
					"reference":  api.ReferenceApi{},
				},
			},
		},
		DbPath: core.DataDir,
	})
}
