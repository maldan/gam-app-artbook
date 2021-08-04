package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-artbook/internal/app/artbook"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
