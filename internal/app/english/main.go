package english

import (
	"embed"
	"flag"
	"fmt"

	"github.com/maldan/gam-app-english/internal/app/english/api"
	"github.com/maldan/gam-app-english/internal/app/english/core"
	"github.com/maldan/go-restserver"
)

func Start(frontFs embed.FS) {
	// Server
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16000, "Server Port")

	// Data
	var dataDir = flag.String("dataDir", "db", "Data Directory")
	_ = flag.String("appId", "id", "App id")
	flag.Parse()

	// Set
	core.DataDir = *dataDir

	// Init server
	restserver.Start(fmt.Sprintf("%s:%d", *host, *port), map[string]interface{}{
		"/": restserver.VirtualFs{Root: "frontend/build/", Fs: frontFs},
		"/api": map[string]interface{}{
			"main": api.MainApi{},
		},
	})
}
