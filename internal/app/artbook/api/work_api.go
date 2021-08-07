package api

import (
	"sort"
	"strings"
	"time"

	"github.com/maldan/gam-app-artbook/internal/app/artbook/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_time"
	"github.com/maldan/go-restserver"
)

type WorkApi struct {
}

// Get by id
func (r WorkApi) GetIndex(args ArgsId) core.Work {
	// Get file
	var item core.Work
	err := cmhp_file.ReadJSON(core.DataDir+"/work/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Work not found!")
	}
	return item
}

// Get list
func (r WorkApi) GetList() []core.Work {
	files, _ := cmhp_file.List(core.DataDir + "/work")
	out := make([]core.Work, 0)
	for _, file := range files {
		out = append(out, r.GetIndex(ArgsId{Id: strings.Replace(file.Name(), ".json", "", 1)}))
	}
	sort.SliceStable(out, func(i, j int) bool {
		// First max date
		maxDate1 := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
		for ii, e := range out[i].ImageList {
			if ii == 0 || e.Created.Unix() > maxDate1.Unix() {
				maxDate1 = e.Created
			}
		}

		// Second max date
		maxDate2 := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
		for ii, e := range out[j].ImageList {
			if ii == 0 || e.Created.Unix() > maxDate2.Unix() {
				maxDate2 = e.Created
			}
		}

		return maxDate1.Unix() > maxDate2.Unix()
	})
	return out
}

// Get year calory stat
func (r WorkApi) GetYearMap(args ArgsDate) map[string]int {
	// Get work list
	files, _ := cmhp_file.List(core.DataDir + "/work")
	workList := make([]core.Work, 0)
	for _, file := range files {
		workList = append(workList, r.GetIndex(ArgsId{Id: strings.Replace(file.Name(), ".json", "", 1)}))
	}

	// Get training list
	t := TrainingApi{}
	files, _ = cmhp_file.List(core.DataDir + "/training")
	trainingList := make([]core.Training, 0)
	for _, file := range files {
		trainingList = append(trainingList, t.GetIndex(ArgsId{Id: strings.Replace(file.Name(), ".json", "", 1)}))
	}

	// Fill empty
	out := map[string]int{}
	t1 := time.Date(args.Date.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 366; i++ {
		t2 := t1.AddDate(0, 0, i)
		out[cmhp_time.Format(t2, "YYYY-MM-DD")] = 0
	}

	// Fill real
	for _, work := range workList {
		for _, image := range work.ImageList {
			if image.Created.Year() == args.Date.Year() {
				out[cmhp_time.Format(image.Created, "YYYY-MM-DD")] += image.Time
			}
		}
	}
	for _, training := range trainingList {
		if training.Created.Year() == args.Date.Year() {
			out[cmhp_time.Format(training.Created, "YYYY-MM-DD")] += training.Time
		}
	}

	return out
}

// Create new
func (r WorkApi) PostIndex(args core.Work) {
	args.Id = cmhp_crypto.UID(10)
	args.ImageList = make([]core.Image, 0)
	cmhp_file.WriteJSON(core.DataDir+"/work/"+args.Id+".json", &args)
}

// Update
func (r WorkApi) PatchIndex(args core.Work) {
	// Get file
	var item core.Work
	err := cmhp_file.ReadJSON(core.DataDir+"/work/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Work not found!")
	}

	// Set values
	item.Title = args.Title
	item.Tags = args.Tags

	// Write to file
	cmhp_file.WriteJSON(core.DataDir+"/work/"+args.Id+".json", &item)
}

// Delete
func (r WorkApi) DeleteIndex(args ArgsId) {
	cmhp_file.Delete(core.DataDir + "/work/" + args.Id + ".json")
}
