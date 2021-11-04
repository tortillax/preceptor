package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"path"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/jonas747/dca"
)

type MusicBot struct {
	lp     string
	router *gin.Engine

	session *discordgo.Session
	voice   *discordgo.VoiceConnection

	playlist    []string
	currentSong string
}

func NewMusicBot(libraryPath string) (*MusicBot, error) {
	r := gin.Default()
	r.GET("/info/playlists", handleInfoPlaylists)
	r.GET("/info/playlist", handleInfoPlaylist)
	r.GET("/info/status", handleInfoStatus)
	r.GET("/info/servers", handleInfoServers)
	r.GET("/info/channels/:guild", handleInfoChannels)
	r.GET("/action/connect/:server/:channel", handleActionConnect)
	r.GET("/action/disconnect", handleActionDisconnect)
	r.GET("/action/next", handleActionNext)
	r.GET("/action/setPlaylist/:id", handleActionSetPlaylist)

	mb := &MusicBot{
		lp:     libraryPath,
		router: r,

		session: nil,
		voice:   nil,

		playlist:    []string{},
		currentSong: "idle",
	}

	return mb, nil
}

func (mb *MusicBot) Connect(token string) error {
	ds, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}

	ds.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	if err = ds.Open(); err != nil {
		return err
	}

	bot.session = ds
	return nil
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

func (mb *MusicBot) Next() error {
	if len(mb.playlist) == 0 {
		return fmt.Errorf("playlist is empty")
	}

	mb.currentSong = mb.playlist[0]
	mb.playlist = mb.playlist[1:]

	go mb.play(mb.currentSong)
	return nil
}

func (mb *MusicBot) StartPanel(address string) error {
	return mb.router.Run(address)
}

func (mb *MusicBot) play(path string) {
	enc, err := dca.EncodeFile(path, dca.StdEncodeOptions)
	if err != nil {
		// Handle the error
		return
	}
	defer enc.Cleanup()

	mb.voice.Speaking(true)
	done := make(chan error)
	dca.NewStream(enc, mb.voice, done)
	mb.voice.Speaking(false)
	err = <-done
	if err != nil && err != io.EOF {
		// Handle the error
		return
	}

	mb.Next()
}
