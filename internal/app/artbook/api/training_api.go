package api

import (
	"os"
	"sort"
	"strings"

	"github.com/maldan/gam-app-artbook/internal/app/artbook/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-cmhp/cmhp_s3"
	"github.com/maldan/go-restserver"
)

type TrainingApi struct {
}

// Get by id
func (r TrainingApi) GetIndex(args ArgsId) core.Training {
	// Get file
	var item core.Training
	err := cmhp_file.ReadJSON(core.DataDir+"/training/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(404, restserver.ErrorType.NotFound, "id", "Training not found!")
	}
	return item
}

// Get list
func (r TrainingApi) GetList() []core.Training {
	files, _ := cmhp_file.List(core.DataDir + "/training")
	out := make([]core.Training, 0)
	for _, file := range files {
		out = append(out, r.GetIndex(ArgsId{Id: strings.Replace(file.Name(), ".json", "", 1)}))
	}
	sort.SliceStable(out, func(i, j int) bool {
		return out[i].Created.Unix() > out[j].Created.Unix()
	})
	return out
}

// Create new
func (r TrainingApi) PostIndex(args ArgsTraining) {
	// Create temp file
	tempFile := os.TempDir() + "/" + cmhp_crypto.UID(10)
	cmhp_file.WriteBin(tempFile, args.Files[0])
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
		Path:        "art_training/" + imageId + ".webp",
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
		Path:        "art_training/" + imageId + "_thumbnail.webp",
		InputData:   thumb,
		Visibility:  "public-read",
		ContentType: "image/webp",
	})
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "id", err.Error())
	}

	// Create new training
	training := core.Training{
		Id:        imageId,
		Title:     args.Title,
		Tags:      strings.Split(args.Tags, ","),
		Time:      args.Time,
		Url:       url,
		Thumbnail: thumbUrl,
		Created:   args.Created,
	}

	// Save to file
	cmhp_file.WriteJSON(core.DataDir+"/training/"+training.Id+".json", &training)
}

// Update
func (r TrainingApi) PatchIndex(args core.Training) {
	// Get file
	var item core.Training
	err := cmhp_file.ReadJSON(core.DataDir+"/training/"+args.Id+".json", &item)
	if err != nil {
		restserver.Fatal(404, restserver.ErrorType.NotFound, "id", "Training not found!")
	}

	// Set values
	item.Title = args.Title
	item.Tags = args.Tags
	item.Time = args.Time
	item.Created = args.Created

	// Write to file
	cmhp_file.WriteJSON(core.DataDir+"/training/"+args.Id+".json", &item)
}

// Delete
func (r TrainingApi) DeleteIndex(args ArgsId) {
	cmhp_file.Delete(core.DataDir + "/training/" + args.Id + ".json")
}
