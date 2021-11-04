package main

import (
	"io/ioutil"

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
