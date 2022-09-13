package api

import (
	"sort"
	"strings"
	"time"

	"github.com/maldan/gam-app-artbook/internal/app/artbook/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-restserver"
)

type ProjectApi struct {
}

// Get by id
func (r ProjectApi) GetIndex(args ArgsId) core.Project {
	// Get file
	var item core.Project
	err := cmhp_file.ReadJSON(core.DataDir+"/work/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Work not found!")
	}

	for i, _ := range item.ImageList {
		item.ImageList[i].Url = "//" + core.Hostname + "/db" + item.ImageList[i].Url
		item.ImageList[i].Thumbnail = "//" + core.Hostname + "/db" + item.ImageList[i].Thumbnail
	}

	return item
}

// Get list
func (r ProjectApi) GetList() []core.Project {
	files, _ := cmhp_file.List(core.DataDir + "/work")
	out := make([]core.Project, 0)
	for _, file := range files {
		out = append(out, r.GetIndex(ArgsId{Id: strings.Replace(file.Name, ".json", "", 1)}))
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

// Create new
func (r ProjectApi) PostIndex(args core.Project) {
	args.Id = cmhp_crypto.UID(10)
	args.ImageList = make([]core.Image, 0)
	cmhp_file.Write(core.DataDir+"/work/"+args.Id+".json", &args)
}

// Update
func (r ProjectApi) PatchIndex(args core.Project) {
	// Get file
	var item core.Project
	err := cmhp_file.ReadJSON(core.DataDir+"/work/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Work not found!")
	}

	// Set values
	item.Title = args.Title
	item.Tags = args.Tags

	// Write to file
	cmhp_file.Write(core.DataDir+"/work/"+args.Id+".json", &item)
}

// Delete
func (r ProjectApi) DeleteIndex(args ArgsId) {
	cmhp_file.Delete(core.DataDir + "/work/" + args.Id + ".json")
}
