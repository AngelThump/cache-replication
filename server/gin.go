package server

import (
	"strings"

	"github.com/angelthump/cache-replication/client"
	utils "github.com/angelthump/cache-replication/utils"
	"github.com/gin-gonic/gin"
)

func Initalize() {
	if utils.Config.GinReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/hls/:channel/:endUrl", func(c *gin.Context) {
		channel := c.Param("channel")
		endUrl := c.Param("endUrl")

		key := channel + "/" + endUrl

		data, err := client.Rdb.Get(client.Ctx, key).Result()
		if err != nil {
			c.AbortWithStatus(404)
			return
		}

		c.Header("Access-Control-Allow-Origin", "*")

		if strings.HasSuffix(endUrl, ".ts") {
			c.Data(200, "video/mp2t", []byte(data))
		} else if strings.HasSuffix(endUrl, ".m3u8") {
			c.Data(200, "application/x-mpegURL", []byte(data))
		} else {
			c.AbortWithStatus(400)
		}
	})

	router.Run(":" + utils.Config.Port)
}
