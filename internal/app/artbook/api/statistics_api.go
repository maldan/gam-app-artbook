package api

import (
	"strings"
	"time"

	"github.com/maldan/gam-app-artbook/internal/app/artbook/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_time"
)

type StatisticsApi struct {
}

// Get year calory stat
func (r StatisticsApi) GetYearMap(args ArgsDate) map[string]float32 {
	w := ArtApi{}

	// Get work list
	files, _ := cmhp_file.List(core.DataDir + "/work")
	workList := make([]core.Art, 0)
	for _, file := range files {
		workList = append(workList, w.GetIndex(ArgsId{Id: strings.Replace(file.Name(), ".json", "", 1)}))
	}

	// Get training list
	t := TrainingApi{}
	files, _ = cmhp_file.List(core.DataDir + "/training")
	trainingList := make([]core.Training, 0)
	for _, file := range files {
		trainingList = append(trainingList, t.GetIndex(ArgsId{Id: strings.Replace(file.Name(), ".json", "", 1)}))
	}

	// Fill empty
	out := map[string]float32{}
	t1 := time.Date(args.Date.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 366; i++ {
		t2 := t1.AddDate(0, 0, i)
		out[cmhp_time.Format(t2, "YYYY-MM-DD")] = 0
	}

	// Fill real
	for _, work := range workList {
		for _, image := range work.ImageList {
			if image.Created.Year() == args.Date.Year() {
				out[cmhp_time.Format(image.Created, "YYYY-MM-DD")] += float32(image.Time) / 3600
			}
		}
	}
	for _, training := range trainingList {
		if training.Created.Year() == args.Date.Year() {
			out[cmhp_time.Format(training.Created, "YYYY-MM-DD")] += float32(training.Time) / 3600
		}
	}

	return out
}
