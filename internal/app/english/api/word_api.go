package api

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/maldan/go-cmhp/cmhp_net"
	"github.com/maldan/go-cmhp/cmhp_slice"
	"os"
	"strings"

	"github.com/maldan/gam-app-english/internal/app/english/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-rapi/rapi_core"
)

type WordApi struct {
}

// GetIndex Get word
func (r WordApi) GetIndex(args core.Word) core.Word {
	out := core.Word{}
	err := cmhp_file.ReadJSON(core.DataDir+"/word/"+args.Name+".json", &out)
	rapi_core.FatalIfError(err)
	return out
}

// GetPlay Play word sound
func (r WordApi) GetPlay(context *rapi_core.Context, args core.Word) string {
	context.IsServeFile = true
	err := os.MkdirAll("/tmp/tts", 0777)
	rapi_core.FatalIfError(err)
	_, err = cmhp_process.Exec("pico2wave", "-w=/tmp/tts/"+args.Name+".wav", args.Name)
	rapi_core.FatalIfError(err)
	return "/tmp/tts/" + args.Name + ".wav"
}

// GetList Get word list
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

// PostIndex Save word
func (r WordApi) PostIndex(args core.Word) {
	args.Category = append(args.Category, strings.ToLower(string(args.Name[0])))
	args.Category = cmhp_slice.Unique(args.Category)

	if args.Translate.Noun == nil {
		args.Translate.Noun = make([]string, 0)
	}
	if args.Translate.Verb == nil {
		args.Translate.Verb = make([]string, 0)
	}
	if args.Translate.Adjective == nil {
		args.Translate.Adjective = make([]string, 0)
	}

	err := cmhp_file.WriteJSON(core.DataDir+"/word/"+args.Name+".json", &args)
	rapi_core.FatalIfError(err)
}

// PatchIndex Save word
func (r WordApi) PatchIndex(args core.Word) {
	args.Category = append(args.Category, strings.ToLower(string(args.Name[0])))
	args.Category = cmhp_slice.Unique(args.Category)
	cmhp_file.WriteJSON(core.DataDir+"/word/"+args.Name+".json", &args)
}

// Delete word
func (r WordApi) DeleteIndex(args core.Word) {
	cmhp_file.Delete(core.DataDir + "/word/" + args.Name + ".json")
}

// GetTranslate Translate a word
func (r WordApi) GetTranslate(args core.Word) core.Word {
	response := cmhp_net.Request(cmhp_net.HttpArgs{
		Method: "GET",
		Url:    "https://context.reverso.net/translation/english-russian/" + args.Name,
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
		},
	})

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))
	rapi_core.FatalIfError(err)

	noun := make([]string, 0)
	verb := make([]string, 0)
	adjective := make([]string, 0)
	adverb := make([]string, 0)

	// Find the review items
	gas := func(i int, s *goquery.Selection) {
		if s.HasClass("adj") {
			adjective = append(adjective, strings.Trim(s.Text(), " \n\r\t"))
		}
		if s.HasClass("adv") {
			adverb = append(adverb, strings.Trim(s.Text(), " \n\r\t"))
		}
		if s.HasClass("n") {
			noun = append(noun, strings.Trim(s.Text(), " \n\r\t"))
		}
		if s.HasClass("v") {
			verb = append(verb, strings.Trim(s.Text(), " \n\r\t"))
		}
	}
	doc.Find("#translations-content > a, #translations-content > div").Each(gas)
	doc.Find("#translations-content > div").Each(gas)

	outWord := core.Word{
		Name: args.Name,
		Translate: core.Translate{
			Noun:      cmhp_slice.Unique(noun),
			Verb:      cmhp_slice.Unique(verb),
			Adjective: cmhp_slice.Unique(adjective),
			Adverb:    cmhp_slice.Unique(adverb),
		},
	}

	return outWord
}
