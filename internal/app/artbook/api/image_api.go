package api

import (
	"os"

	"github.com/maldan/gam-app-artbook/internal/app/artbook/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-cmhp/cmhp_s3"
	"github.com/maldan/go-restserver"
)

type ImageApi struct {
}

// Get by id
func (r ImageApi) GetIndex(args ArgsImage) core.Image {
	// Get file
	var item core.Project
	err := cmhp_file.ReadJSON(core.DataDir+"/work/"+args.ProjectId+".json", &item)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.NotFound, "workId", "Work not found!")
	}

	// Find image
	for _, image := range item.ImageList {
		if image.Id == args.Id {
			return image
		}
	}

	return core.Image{}
}

// Create new
func (r ImageApi) PostIndex(args ArgsImage) {
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
		Path:        "art/" + imageId + ".webp",
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
		Path:        "art/" + imageId + "_thumbnail.webp",
		InputData:   thumb,
		Visibility:  "public-read",
		ContentType: "image/webp",
	})
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "id", err.Error())
	}

	// Open work and add new image
	var work core.Project
	cmhp_file.ReadJSON(core.DataDir+"/work/"+args.ProjectId+".json", &work)
	work.ImageList = append(work.ImageList, core.Image{
		Id:        imageId,
		WorkId:    args.ProjectId,
		Url:       url,
		Thumbnail: thumbUrl,
		Time:      args.Time,
		Created:   args.Created,
	})
	cmhp_file.WriteJSON(core.DataDir+"/work/"+args.ProjectId+".json", &work)
}

// Update
func (r ImageApi) PatchIndex(args core.Image) {
	// Open work and add new image
	var work core.Project
	cmhp_file.ReadJSON(core.DataDir+"/work/"+args.WorkId+".json", &work)

	// Find and update image
	for index, image := range work.ImageList {
		if image.Id == args.Id {
			work.ImageList[index].Time = args.Time
			work.ImageList[index].Created = args.Created
		}
	}

	// Save
	cmhp_file.WriteJSON(core.DataDir+"/work/"+args.WorkId+".json", &work)
}

// Move image down
func (r ImageApi) PatchMoveDown(args core.Image) {
	// Open work
	var work core.Project
	cmhp_file.ReadJSON(core.DataDir+"/work/"+args.WorkId+".json", &work)

	// Find and move image
	for index, image := range work.ImageList {
		if image.Id == args.Id {
			if index == len(work.ImageList)-1 {
				return
			}
			// Swap images
			temp := image
			temp2 := work.ImageList[index+1]
			work.ImageList[index] = temp2
			work.ImageList[index+1] = temp
			break
		}
	}

	// Save
	cmhp_file.WriteJSON(core.DataDir+"/work/"+args.WorkId+".json", &work)
}

// Delete
func (r ImageApi) DeleteIndex(args core.Image) {
	// Open work and add new image
	var work core.Project
	cmhp_file.ReadJSON(core.DataDir+"/work/"+args.WorkId+".json", &work)

	out := make([]core.Image, 0)
	for _, image := range work.ImageList {
		if image.Id != args.Id {
			out = append(out, image)
		}
	}
	work.ImageList = out

	// Save
	cmhp_file.WriteJSON(core.DataDir+"/work/"+args.WorkId+".json", &work)
}
