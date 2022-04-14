/*
Server
*/
package main

import (
	dataSource "github.com/WhitePaper233/PixGreatServer/src"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	if MetadataLoadToMem {
		MetadataMap, err := dataSource.LoadMetadata(dataSource.GetIndexList())
		if err != nil {
			panic(err)
		}
		router.GET("/", func(ctx *gin.Context) {
			data := dataSource.GetRandomMetadataFromMem(MetadataMap)
			ctx.JSON(200, data)
		})
		router.Run()
	} else {
		router.GET("/", func(ctx *gin.Context) {
			data, err := dataSource.GetRandomMetadata()
			if err != nil {
				panic(err)
			}
			ctx.JSON(200, data)
		})
		router.Run()
	}
}
