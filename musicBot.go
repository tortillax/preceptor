package main

import "github.com/gin-gonic/gin"

type MusicBot struct {
	lp     string
	router *gin.Engine
}

func NewMusicBot(libraryPath string) (*MusicBot, error) {
	r := gin.Default()
	r.GET("/info/playlists", handleInfoPlaylists)

	mb := &MusicBot{
		lp:     libraryPath,
		router: r,
	}

	return mb, nil
}

func (mb *MusicBot) StartPanel(address string) error {
	return mb.router.Run(address)
}
