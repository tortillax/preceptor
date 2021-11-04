package main

import (
	"io/ioutil"
	"path"
	"strconv"

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
