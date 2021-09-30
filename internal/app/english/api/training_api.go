package api

import (
	"github.com/maldan/gam-app-english/internal/app/english/core"
	"github.com/maldan/go-cmhp/cmhp_file"
)

type TrainingApi struct {
}

func (r TrainingApi) GetWord(args ArgsCategory) core.Word {
	wordList, _ := cmhp_file.List(core.DataDir + "/word")
	for _, wordFile := range wordList {
		word := core.Word{}
		cmhp_file.ReadJSON(core.DataDir+"/word/"+wordFile.Name(), &word)

		for _, cat := range word.Category {
			if cat == args.Category && word.Power <= 0 {
				return word
			}
		}
	}

	word := core.Word{}
	return word
}

func (r TrainingApi) PostKnowWord(args core.Word) {
	word := core.Word{}
	cmhp_file.ReadJSON(core.DataDir+"/word/"+args.Name+".json", &word)
	word.Power += 1
	cmhp_file.WriteJSON(core.DataDir+"/word/"+args.Name+".json", &word)
}

func (r TrainingApi) PostDontKnowWord(args core.Word) {
	word := core.Word{}
	cmhp_file.ReadJSON(core.DataDir+"/word/"+args.Name+".json", &word)
	word.Power -= 1
	cmhp_file.WriteJSON(core.DataDir+"/word/"+args.Name+".json", &word)
}
