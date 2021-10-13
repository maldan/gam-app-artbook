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

type ArtApi struct {
}

// Get by id
func (r ArtApi) GetIndex(args ArgsId) core.Art {
	// Get file
	var item core.Art
	err := cmhp_file.ReadJSON(core.DataDir+"/work/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "id", "Work not found!")
	}
	return item
}

// Get list
func (r ArtApi) GetList() []core.Art {
	files, _ := cmhp_file.List(core.DataDir + "/work")
	out := make([]core.Art, 0)
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

// Create new
func (r ArtApi) PostIndex(args core.Art) {
	args.Id = cmhp_crypto.UID(10)
	args.ImageList = make([]core.Image, 0)
	cmhp_file.WriteJSON(core.DataDir+"/work/"+args.Id+".json", &args)
}

// Update
func (r ArtApi) PatchIndex(args core.Art) {
	// Get file
	var item core.Art
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
func (r ArtApi) DeleteIndex(args ArgsId) {
	cmhp_file.Delete(core.DataDir + "/work/" + args.Id + ".json")
}
