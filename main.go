package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-english/internal/app/english"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
