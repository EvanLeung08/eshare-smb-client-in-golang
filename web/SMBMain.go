package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	client, err := NewSMBClient("192.168.50.69", "user", "123456", "LANdrive")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		localPath := c.PostForm("localPath")
		remotePath := c.PostForm("remotePath")
		fmt.Printf("localPath: %s, remotePath: %s\n", localPath, remotePath)
		err := client.Upload(localPath, remotePath)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Upload successful"})
	})

	router.GET("/download", func(c *gin.Context) {
		basePath := c.Query("basePath")
		remotePattern := c.Query("remotePattern")
		localPath := c.Query("localPath")

		err := client.Download(basePath, remotePattern, localPath)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Download successful"})
	})

	router.PUT("/rename", func(c *gin.Context) {
		basePath := c.Query("basePath")
		oldPattern := c.Query("oldPattern")
		suffix := c.Query("suffix")

		err := client.Rename(basePath, oldPattern, suffix)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Rename successful"})
	})

	router.GET("/search", func(c *gin.Context) {
		basePath := c.Query("basePath")
		pattern := c.Query("pattern")

		results, err := client.Search(basePath, pattern)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"files": results})
	})

	router.DELETE("/delete", func(c *gin.Context) {
		basePath := c.Query("basePath")
		remotePattern := c.Query("remotePattern")

		err := client.Delete(basePath, remotePattern)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Delete successful"})
	})

	router.Run(":8080")
}
