package main

import (
	"io/ioutil"
	"math/rand"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type MusicBot struct {
	lp     string
	router *gin.Engine

	playlist    []string
	currentSong string
}

func NewMusicBot(libraryPath string) (*MusicBot, error) {
	r := gin.Default()
	r.GET("/info/playlists", handleInfoPlaylists)
	r.GET("/info/playlist", handleInfoPlaylist)
	r.GET("/info/status", handleInfoStatus)
	//r.GET("/action/play", handleActionPlay)
	//r.GET("/action/stop", handleActionStop)
	//r.GET("/action/next", handleActionNext)
	r.GET("/action/setPlaylist/:id", handleActionSetPlaylist)

	mb := &MusicBot{
		lp:     libraryPath,
		router: r,

		playlist:    []string{},
		currentSong: "idle",
	}

	return mb, nil
}

func (mb *MusicBot) SetPlaylist(playlistPath string) error {
	pl := []string{}

	fsi, err := ioutil.ReadDir(playlistPath)
	if err != nil {
		return err
	}

	for _, f := range fsi {
		if !f.IsDir() && strings.HasSuffix(strings.ToLower(f.Name()), ".mp3") {
			pl = append(pl, path.Join(playlistPath, f.Name()))
		}
	}

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(pl), func(i, j int) {
		pl[i], pl[j] = pl[j], pl[i]
	})

	mb.playlist = pl
	return nil
}

func (mb *MusicBot) StartPanel(address string) error {
	return mb.router.Run(address)
}
