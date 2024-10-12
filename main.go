// If you disassembled this, then know that all the code was written in a day :_(
package main

import (
	"rucq/api/crud"
	"rucq/api/requester"
	"rucq/webserver"

	"github.com/gin-gonic/gin"
)

const webIP = "0.0.0.0:10015"

// const websocketIP = "0.0.0.0:10016"
const apiIP = "0.0.0.0:10017"

func main() {
	go webserver.StartWebserver(webIP)
	//go requester.StartWebSocketServer(websocketIP)

	crud.InitDB()
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.POST("/authUser", requester.Authenticate)
	r.POST("/addUser", requester.AddUser)
	r.POST("/getMessages", requester.GetMessages)
	r.POST("/rooms", requester.GetRooms)
	r.POST("/sendMessage", requester.SendMessage)
	r.POST("/connectRoom", requester.ConnectRoom)
	r.POST("/createRoom", requester.CreateRoom)

	r.Run(apiIP)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
