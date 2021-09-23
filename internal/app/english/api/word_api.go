package api

import (
	"os"

	"github.com/maldan/gam-app-english/internal/app/english/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-rapi/rapi_core"
)

type WordApi struct {
}

// Get word
func (r WordApi) GetIndex(args core.Word) core.Word {
	out := core.Word{}
	cmhp_file.ReadJSON(core.DataDir+"/word/"+args.Name+".json", &out)
	return out
}

// Play word sound
func (r WordApi) GetPlay(context *rapi_core.Context, args core.Word) string {
	context.IsServeFile = true
	os.MkdirAll("/tmp/tts", 0777)
	cmhp_process.Exec("pico2wave", "-w=/tmp/tts/"+args.Name+".wav", args.Name)
	return "/tmp/tts/" + args.Name + ".wav"
}

// Get word list
func (r WordApi) GetList() []core.Word {
	out := make([]core.Word, 0)
	l, _ := cmhp_file.List(core.DataDir + "/word")
	for _, f := range l {
		word := core.Word{}
		cmhp_file.ReadJSON(core.DataDir+"/word/"+f.Name(), &word)
		out = append(out, word)
	}
	return out
}

// Save word
func (r WordApi) PostIndex(args core.Word) {
	cmhp_file.WriteJSON(core.DataDir+"/word/"+args.Name+".json", &args)
}

// Save word
func (r WordApi) PatchIndex(args core.Word) {
	cmhp_file.WriteJSON(core.DataDir+"/word/"+args.Name+".json", &args)
}

// Delete word
func (r WordApi) DeleteIndex(args core.Word) {
	cmhp_file.Delete(core.DataDir + "/word/" + args.Name + ".json")
}
