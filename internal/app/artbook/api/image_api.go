package api

import (
	"os"

	"github.com/maldan/gam-app-artbook/internal/app/artbook/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
)

type ImageApi struct {
}

// Create new
func (r ImageApi) PostIndex(args ArgsImage) {
	// Create temp file
	tempFile := os.TempDir() + "/" + cmhp_crypto.UID(10)
	cmhp_file.WriteBin(tempFile, args.Files[0])
	defer cmhp_file.Delete(tempFile)

	// Convert to jpeg & remove later
	cmhp_process.Exec("magick", tempFile, tempFile+".webp")
	defer cmhp_file.Delete(tempFile + ".webp")

	// Open work and add new image
	var work core.Work
	cmhp_file.ReadJSON(core.DataDir+"/work/"+args.WorkId+".json", &work)
	work.ImageList = append(work.ImageList, core.Image{
		Id: cmhp_crypto.UID(10),
	})
	cmhp_file.WriteJSON(core.DataDir+"/work/"+args.WorkId+".json", &work)
}

// Update
func (r ImageApi) PatchIndex(args core.Image) {
	//cmhp_file.WriteJSON(core.DataDir+"/work/"+args.Id+".json", &args)
}

// Delete
func (r ImageApi) DeleteIndex(args core.Image) {
	//cmhp_file.Delete(core.DataDir + "/work/" + args.Id + ".json")
}
