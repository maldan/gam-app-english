package api

import (
	"github.com/maldan/gam-app-english/internal/app/english/core"
	"github.com/maldan/go-cmhp/cmhp_file"
)

type WordApi struct {
}

// Get word
func (r WordApi) GetIndex(args core.Word) core.Word {
	out := core.Word{}
	cmhp_file.ReadJSON(core.DataDir+"/word/"+args.Name+".json", &out)
	return out
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
