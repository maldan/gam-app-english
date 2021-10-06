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

func (r TrainingApi) GetCategoryList() []core.Category {
	// Prepare
	list := make([]core.Category, 0)
	catMap := make(map[string]int)

	// Get word list
	wordList, _ := cmhp_file.List(core.DataDir + "/word")
	for _, wordFile := range wordList {
		// Get word
		word := core.Word{}
		cmhp_file.ReadJSON(core.DataDir+"/word/"+wordFile.Name(), &word)

		// Fill category
		for _, category := range word.Category {
			catMap[category] += 1
		}
	}

	// Fill result
	for category, amount := range catMap {
		list = append(list, core.Category{Name: category, Amount: amount})
	}

	return list
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

// Reset word power
func (r TrainingApi) PostReset(args ArgsCategory) {
	// Get words
	wordList, _ := cmhp_file.List(core.DataDir + "/word")
	for _, wordFile := range wordList {
		// Read word
		word := core.Word{}
		cmhp_file.ReadJSON(core.DataDir+"/word/"+wordFile.Name(), &word)

		// If word has category
		for _, cat := range word.Category {
			if cat == args.Category {
				// Reset power and save
				word.Power = 0
				cmhp_file.WriteJSON(core.DataDir+"/word/"+word.Name+".json", &word)
			}
		}
	}
}
