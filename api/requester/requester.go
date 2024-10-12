package requester

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rucq/api/crud"
	"rucq/api/data_scheme"
	"rucq/webserver"

	"github.com/gin-gonic/gin"
)

// * Sign in
func Authenticate(c *gin.Context) {
	var User data_scheme.User
	if err := c.BindJSON(&User); err != nil {
		fmt.Println(err)
		return
	}
	if authenticated, token := crud.GetUserDB(&User); authenticated {
		c.IndentedJSON(http.StatusAccepted, data_scheme.User{
			Token: token,
		})
		return
	}
	c.IndentedJSON(http.StatusForbidden, data_scheme.User{
		Token: "0",
	})
}

// * Sign up
func AddUser(c *gin.Context) {
	message := struct {
		Message string `json:"message"`
	}{
		Message: "user created",
	}

	var User data_scheme.User
	if err := c.BindJSON(&User); err != nil {
		fmt.Println(err)
		message.Message = "error"
		c.IndentedJSON(http.StatusInternalServerError, message)
		return
	}
	if err := crud.AddUserDB(&User); err != nil {
		fmt.Println(err)
		if err.Error() == "user already exists" {
			message.Message = "user already exists"
			c.IndentedJSON(http.StatusConflict, message)
			return
		}
		message.Message = "error"
		c.IndentedJSON(http.StatusInternalServerError, message)
		return
	}

	c.IndentedJSON(http.StatusAccepted, message)
}

// * Get messages
func GetMessages(c *gin.Context) {
	var User data_scheme.MessagesGet
	if err := c.BindJSON(&User); err != nil {
		return
	}

	mesSend, err := crud.GetMessagesDB(&User)
	if err != nil {
		fmt.Println(err)
		mesSend.Error = err.Error()
		c.IndentedJSON(http.StatusInternalServerError, mesSend)
		return
	}
	mesSend.Error = ""
	c.IndentedJSON(http.StatusOK, mesSend)
}

// * Send message from user
func SendMessage(c *gin.Context) {
	message := struct {
		Message string `json:"message"`
	}{
		Message: "",
	}
	var Mes data_scheme.MesContainerSend
	if err := c.BindJSON(&Mes); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, message)
		return
	}
	username, err := crud.AddMessageDB(&Mes)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, message)
		return
	}
	message.Message = username
	c.IndentedJSON(http.StatusCreated, message)
	for client := range webserver.HubMain.Clients {
		if client.RoomId == Mes.RoomId && client.Username != username {
			originMessage := data_scheme.MessagesSend{
				Username: username,
				MessageList: []data_scheme.MesContainer{
					{
						Username: username,
						Message:  Mes.Message,
					},
				},
				Error: "",
			}
			rawMessage, _ := json.Marshal(originMessage)
			client.Send <- rawMessage
		}
	}
}

// * Add user to room
func ConnectRoom(c *gin.Context) {
	message := struct {
		Secret string `json:"secret"`
		Error  string `json:"error"`
	}{
		Secret: "",
		Error:  "",
	}

	var User data_scheme.UsersMapConnectRoom
	if err := c.BindJSON(&User); err != nil {
		fmt.Println(err)
		return
	}
	secret, err := crud.AddUserToRoomDB(&User)
	if err != nil {
		fmt.Println(err)
		message.Secret = ""
		message.Error = err.Error()
		c.IndentedJSON(http.StatusInternalServerError, message)
		return
	}
	message.Secret = secret
	c.IndentedJSON(http.StatusAccepted, message)
}

// * Create room by user
func CreateRoom(c *gin.Context) {
	message := struct {
		Secret string `json:"secret"`
	}{
		Secret: "",
	}

	var User data_scheme.UsersMap
	if err := c.BindJSON(&User); err != nil {
		fmt.Println(err)
		return
	}
	secret, err := crud.AddRoomDB(&User)
	if err != nil {
		fmt.Println(err)
		message.Secret = ""
		c.IndentedJSON(http.StatusInternalServerError, message)
		return
	}
	message.Secret = secret
	c.IndentedJSON(http.StatusAccepted, message)
}

// * Get rooms id
func GetRooms(c *gin.Context) {
	message := struct {
		Rooms []string `json:"rooms"`
	}{
		Rooms: []string{""},
	}

	var User data_scheme.User
	if err := c.BindJSON(&User); err != nil {
		fmt.Println(err)
		return
	}
	rooms := crud.GetRoomsDB(&User)
	if len(rooms) == 0 || rooms[0] == "" {
		c.IndentedJSON(http.StatusNoContent, message)
		return
	}
	message.Rooms = rooms
	c.IndentedJSON(http.StatusOK, message)
}
