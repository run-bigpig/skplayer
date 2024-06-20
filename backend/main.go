package main

import (
	"embed"
	"github.com/run-bigpig/skplayer/internal/bootstrap"
)

//go:embed web
var web embed.FS

func main() {
	bootstrap.Boot(web)
}
