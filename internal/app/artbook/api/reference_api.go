package api

import (
	"os"
	"strings"

	"github.com/maldan/gam-app-artbook/internal/app/artbook/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-cmhp/cmhp_s3"
	"github.com/maldan/go-restserver"
)

type ReferenceApi struct {
}

// Get by id
func (r ReferenceApi) GetIndex(args ArgsId) core.Reference {
	// Get file
	var item core.Reference
	err := cmhp_file.ReadJSON(core.DataDir+"/reference/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(404, restserver.ErrorType.NotFound, "id", "Reference not found!")
	}
	return item
}

// Get list
func (r ReferenceApi) GetList() []core.Reference {
	files, _ := cmhp_file.List(core.DataDir + "/reference")
	out := make([]core.Reference, 0)
	for _, file := range files {
		out = append(out, r.GetIndex(ArgsId{Id: strings.Replace(file.Name(), ".json", "", 1)}))
	}
	return out
}

// Create new
func (r ReferenceApi) PostIndex(args ArgsReference) {
	// Create temp file
	tempFile := os.TempDir() + "/" + cmhp_crypto.UID(10)
	cmhp_file.WriteBin(tempFile, args.Image.Data)
	defer cmhp_file.Delete(tempFile)

	// Convert  & remove later
	cmhp_process.Exec("magick", tempFile, "-quality", "90", "-define", "webp:lossless=false", tempFile+".webp")
	defer cmhp_file.Delete(tempFile + ".webp")
	imageId, err := cmhp_file.HashSha1(tempFile + ".webp")
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "id", err.Error())
	}

	// Thumbnail & remove later
	cmhp_process.Exec("magick", tempFile, "-quality", "90", "-define", "webp:lossless=false", "-thumbnail", "256x256^", "-gravity", "center", "-extent", "256x256", tempFile+"_thumbnail.webp")
	defer cmhp_file.Delete(tempFile + "_thumbnail.webp")

	// Upload image to s3
	image, err := cmhp_file.ReadBin(tempFile + ".webp")
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "id", err.Error())
	}
	url, err := cmhp_s3.Write(cmhp_s3.WriteArgs{
		Path:        "art_reference/" + imageId + ".webp",
		InputData:   image,
		Visibility:  "public-read",
		ContentType: "image/webp",
	})
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "id", err.Error())
	}

	// Upload thumbnail to s3
	thumb, err := cmhp_file.ReadBin(tempFile + "_thumbnail.webp")
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "id", err.Error())
	}
	thumbUrl, err := cmhp_s3.Write(cmhp_s3.WriteArgs{
		Path:        "art_reference/" + imageId + "_thumbnail.webp",
		InputData:   thumb,
		Visibility:  "public-read",
		ContentType: "image/webp",
	})
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "id", err.Error())
	}

	// Create new reference
	tags := make([]string, 0)
	if args.Tags != "" {
		tags = strings.Split(args.Tags, ",")
	}
	reference := core.Reference{
		Id:        imageId,
		Category:  args.Category,
		Tags:      tags,
		Url:       url,
		Thumbnail: thumbUrl,
	}

	// Save to file
	cmhp_file.WriteJSON(core.DataDir+"/reference/"+reference.Id+".json", &reference)
}

// Update
func (r ReferenceApi) PatchIndex(args core.Reference) {
	// Get file
	var item core.Reference
	err := cmhp_file.ReadJSON(core.DataDir+"/reference/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(404, restserver.ErrorType.NotFound, "id", "Reference not found!")
	}

	// Set values
	item.Category = args.Category
	item.Tags = args.Tags

	// Write to file
	cmhp_file.WriteJSON(core.DataDir+"/reference/"+args.Id+".json", &item)
}

// Delete
func (r ReferenceApi) DeleteIndex(args ArgsId) {
	cmhp_file.Delete(core.DataDir + "/reference/" + args.Id + ".json")
}
