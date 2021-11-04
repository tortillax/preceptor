package main

import (
	"io/ioutil"
	"path"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func handleInfoPlaylists(c *gin.Context) {
	lists := []string{}

	fsi, err := ioutil.ReadDir(bot.lp)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   err.Error(),
			"context": "reading library",
		})
		return
	}

	for _, r := range fsi {
		if r.IsDir() {
			lists = append(lists, r.Name())
		}
	}

	c.JSON(200, lists)
}

func handleInfoPlaylist(c *gin.Context) {
	if len(bot.playlist) == 0 {
		c.JSON(400, gin.H{
			"error":   "playlist is not set",
			"context": "reading playlist",
		})
		return
	}

	c.JSON(200, bot.playlist)
}

func handleInfoStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": bot.currentSong,
	})
}

type gs struct {
	Name string
	Id   string
}

func handleInfoServers(c *gin.Context) {
	ga := make([]*gs, 0)
	for _, g := range bot.session.State.Guilds {
		tg := &gs{
			Name: g.Name,
			Id:   g.ID,
		}

		ga = append(ga, tg)
	}

	c.JSON(200, ga)
}

func handleInfoChannels(c *gin.Context) {
	guildString := c.Param("guild")

	channels, err := bot.session.GuildChannels(guildString)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   err.Error(),
			"context": "reading channels from server",
		})
		return
	}

	cha := make([]*gs, 0)
	for _, ch := range channels {
		chgs := &gs{}
		switch ch.Type {
		case discordgo.ChannelTypeGuildVoice:
			chgs.Name = ch.Name
			chgs.Id = ch.ID

			cha = append(cha, chgs)
		}
	}
	c.JSON(200, cha)
}

func handleActionConnect(c *gin.Context) {
	serverID := c.Param("server")
	channelID := c.Param("channel")

	vc, err := bot.session.ChannelVoiceJoin(serverID, channelID, false, true)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   err.Error(),
			"context": "joining voice channel",
		})
		return
	}

	bot.voice = vc
	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func handleActionDisconnect(c *gin.Context) {
	if bot.voice == nil {
		c.JSON(400, gin.H{
			"error":   "not connected",
			"context": "joining voice channel",
		})
		return
	}

	if err := bot.voice.Disconnect(); err != nil {
		c.JSON(400, gin.H{
			"error":   err.Error(),
			"context": "joining voice channel",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func handleActionNext(c *gin.Context) {
	if err := bot.Next(); err != nil {
		c.JSON(400, gin.H{
			"error":   err.Error(),
			"context": "playing next song",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})
}

func handleActionSetPlaylist(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   err.Error(),
			"context": "converting ID",
		})
		return
	}

	fsi, err := ioutil.ReadDir(bot.lp)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   err.Error(),
			"context": "reading library",
		})
		return
	}

	counter := 0
	for _, r := range fsi {
		if r.IsDir() {
			if counter == idInt {
				err := bot.SetPlaylist(path.Join(bot.lp, r.Name()))
				if err != nil {
					c.JSON(400, gin.H{
						"error":   err.Error(),
						"context": "setting playlist",
					})
					return
				}

				c.JSON(200, gin.H{
					"message": "playlist " + r.Name() + " set",
				})
				return
			}
			counter++
		}
	}

	c.JSON(400, gin.H{
		"error":   "playlist  " + id + " not found",
		"context": "selecting playlist",
	})
}
