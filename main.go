package main

import (
	"flag"
	"log"
)

var bot *MusicBot

func main() {
	libraryPath := flag.String("l", "mp3", "path to music library")
	controlAddr := flag.String("a", "127.0.0.1:7890", "address to spawn control API and UI on")
	flag.Parse()

	mb, err := NewMusicBot(*libraryPath)
	if err != nil {
		log.Fatalf("fatal error creating preceptor instance: %s\n", err.Error())
	}
	bot = mb

	if err = mb.StartPanel(*controlAddr); err != nil {
		log.Fatalf("fatal error running control panel: %s\n", err.Error())
	}
}
