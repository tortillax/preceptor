package main

import (
	"embed"
	"flag"
	"log"
)

//go:embed www/*
var fs embed.FS

var bot *MusicBot

func main() {
	libraryPath := flag.String("l", "mp3", "path to music library")
	controlAddr := flag.String("a", "127.0.0.1:7890", "address to spawn control API and UI on")
	discordToken := flag.String("t", "", "discord API token")
	flag.Parse()

	if *discordToken == "" {
		log.Fatalf("run with -t argument to provide discord API token\n")
	}

	mb, err := NewMusicBot(*libraryPath)
	if err != nil {
		log.Fatalf("fatal error creating preceptor instance: %s\n", err.Error())
	}
	bot = mb

	if err = bot.Connect(*discordToken); err != nil {
		log.Fatalf("fatal error connecting to discord API: %s\n", err.Error())
	}

	if err = mb.StartPanel(*controlAddr); err != nil {
		log.Fatalf("fatal error running control panel: %s\n", err.Error())
	}
}
