package api

import (
	"strings"

	"github.com/maldan/gam-app-artbook/internal/app/artbook/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
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
	/*sort.SliceStable(out, func(i, j int) bool {
		return out[i].Priority > out[j].Priority
	})*/
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
	cmhp_file.WriteJSON(core.DataDir+"/work/"+args.Id+".json", &args)
}

// Delete
func (r WorkApi) DeleteIndex(args ArgsId) {
	cmhp_file.Delete(core.DataDir + "/work/" + args.Id + ".json")
}
