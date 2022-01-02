package api

import (
	"time"

	"github.com/maldan/gam-app-english/internal/app/english/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_time"
)

type StatisticsApi struct {
}

func (r StatisticsApi) GetToday() core.Statistics {
	stat := core.Statistics{}
	cmhp_file.ReadJSON(core.DataDir+"/stat/"+cmhp_time.Format(time.Now(), "YYYY-MM-DD")+".json", &stat)
	return stat
}

func (r StatisticsApi) GetTotalByDate(args ArgsDate) int {
	stat := core.Statistics{}
	cmhp_file.ReadJSON(core.DataDir+"/stat/"+cmhp_time.Format(args.Date, "YYYY-MM-DD")+".json", &stat)
	return len(stat.Correct) + len(stat.Wrong)
}

func (r StatisticsApi) PostCorrect(args core.Word) {
	stat := core.Statistics{}
	cmhp_file.ReadJSON(core.DataDir+"/stat/"+cmhp_time.Format(time.Now(), "YYYY-MM-DD")+".json", &stat)
	stat.Created = time.Now()
	stat.Correct = append(stat.Correct, args.Name)
	cmhp_file.WriteJSON(core.DataDir+"/stat/"+cmhp_time.Format(time.Now(), "YYYY-MM-DD")+".json", &stat)
}

func (r StatisticsApi) PostWrong(args core.Word) {
	stat := core.Statistics{}
	cmhp_file.ReadJSON(core.DataDir+"/stat/"+cmhp_time.Format(time.Now(), "YYYY-MM-DD")+".json", &stat)
	stat.Created = time.Now()
	stat.Wrong = append(stat.Wrong, args.Name)
	cmhp_file.WriteJSON(core.DataDir+"/stat/"+cmhp_time.Format(time.Now(), "YYYY-MM-DD")+".json", &stat)
}

// Get year calory stat
func (f StatisticsApi) GetYearMap(args ArgsDate) map[string]interface{} {
	out := map[string]interface{}{}
	t1 := time.Date(args.Date.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; i < 366; i++ {
		t2 := t1.AddDate(0, 0, i)
		out[cmhp_time.Format(t2, "YYYY-MM-DD")] = f.GetTotalByDate(ArgsDate{Date: t2})
	}

	return out
}
