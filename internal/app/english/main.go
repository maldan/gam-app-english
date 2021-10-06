package english

import (
	"embed"
	"flag"
	"fmt"

	"github.com/maldan/gam-app-english/internal/app/english/api"
	"github.com/maldan/gam-app-english/internal/app/english/core"
	"github.com/maldan/go-rapi"
	"github.com/maldan/go-rapi/rapi_core"
	"github.com/maldan/go-rapi/rapi_rest"
	"github.com/maldan/go-rapi/rapi_vfs"
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

	/*wordList := make([]core.Word, 0)
	cmhp_file.ReadJSON("word.json", &wordList)
	w := api.WordApi{}
	for _, word := range wordList {
		w.PostIndex(core.Word{
			Name: word.Name,
			Translate: core.Translate{
				Noun:      word.Translate.Noun,
				Verb:      word.Translate.Verb,
				Adjective: word.Translate.Adjective,
			},
		})
	}*/

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
					"main":       api.MainApi{},
					"word":       api.WordApi{},
					"training":   api.TrainingApi{},
					"statistics": api.StatisticsApi{},
				},
			},
		},
		DbPath: core.DataDir,
	})
}
