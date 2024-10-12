package crud

import (
	"context"
	"errors"
	"fmt"
	"rucq/api/data_scheme"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const addr = "localhost:5432"
const user = "rucqUser"
const password = "kljk5lj46ioj4o6"
const db_name = "rucqDB"

func connectDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: db_name,
	})
	return db
}

func InitDB() {
	db := connectDB()
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")
	if err := createSchema(db); err != nil {
		panic(err)
	}

	db.Close()
	fmt.Println("Close DB")
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*data_scheme.User)(nil),
		(*data_scheme.Rooms)(nil),
	}
	fmt.Println("Creating DB Schema")
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			fmt.Println("Creating DB Schema:ERROR")
			return err
		}
	}
	fmt.Println("Creating DB Schema:SUCCESS")
	AddFlag(db)
	return nil
}

func AddUserDB(user *data_scheme.User) error {
	db := connectDB()
	defer db.Close()

	userDB := new(data_scheme.User)
	db.Model(userDB).Where("login = ?", user.Login).Select()
	if user.Login != userDB.Login {
		_, err := db.Model(user).Insert()
		return err
	}

	return errors.New("user already exists")
}

func GetUserDB(user *data_scheme.User) (bool, string) {
	db := connectDB()
	defer db.Close()

	userDB := new(data_scheme.User)
	db.Model(userDB).Where("login = ?", user.Login).Select()
	if user.Login == userDB.Login && user.Password == userDB.Password {
		token := GeneratePassword(32)
		db.Model(userDB).Set("token = ?", token).Where("id = ?id").Update()
		return true, token
	}
	return false, ""
}

func GetMessagesDB(user *data_scheme.MessagesGet) (data_scheme.MessagesSend, error) {
	db := connectDB()
	defer db.Close()

	roomDB := new(data_scheme.Rooms)
	db.Model(roomDB).Where("secret = ?", user.RoomId).Select()
	if roomDB.Secret == "" {
		return data_scheme.MessagesSend{}, errors.New("room doesn't exist")
	}

	userDB := new(data_scheme.User)
	db.Model(userDB).Where("token = ?", user.Secret).Select()
	if userDB.Token == "" {
		return data_scheme.MessagesSend{}, errors.New("user doesn't exist")
	}

	username := findUserInRoom(userDB.Login, roomDB)
	if username == "" {
		return data_scheme.MessagesSend{}, errors.New("user doesn't exist in room")
	}

	messages := data_scheme.MessagesSend{
		Username:    username,
		MessageList: roomDB.Messages,
	}
	return messages, nil
}

func AddMessageDB(message *data_scheme.MesContainerSend) (string, error) {
	db := connectDB()
	defer db.Close()

	roomDB := new(data_scheme.Rooms)
	db.Model(roomDB).Where("secret = ?", message.RoomId).Select()
	if roomDB.Secret == "" {
		return "", errors.New("room doesn't exist")
	}

	userDB := new(data_scheme.User)
	db.Model(userDB).Where("token = ?", message.Secret).Select()
	if userDB.Token == "" {
		return "", errors.New("user doesn't exist")
	}

	username := findUserInRoom(userDB.Login, roomDB)
	if username == "" {
		return "", errors.New("user doesn't exist in room")
	}

	roomDB.Messages = append(roomDB.Messages, data_scheme.MesContainer{
		Username: username,
		Message:  message.Message,
	})
	_, err := db.Model(roomDB).Set("messages = ?messages").Where("id = ?id").Update()
	if err != nil {
		return "", err
	}

	return username, nil
}

func AddRoomDB(user *data_scheme.UsersMap) (string, error) {
	db := connectDB()
	defer db.Close()

	roomDB := new(data_scheme.Rooms)
	roomDB.Users.UsersList = append(roomDB.Users.UsersList, data_scheme.UsersMapDB{
		Login:    user.Login,
		Username: user.Username,
	})
	roomDB.Secret = GeneratePassword(32)
	if _, err := db.Model(roomDB).Insert(); err != nil {
		return "", err
	}

	userDB := new(data_scheme.User)
	db.Model(userDB).Where("login = ?", user.Login).Select()
	userDB.Rooms = append(userDB.Rooms, roomDB.Secret)
	_, err := db.Model(userDB).Set("rooms = ?rooms").Where("id = ?id").Update()
	if err != nil {
		return "", err
	}
	return roomDB.Secret, nil
}

func AddUserToRoomDB(user *data_scheme.UsersMapConnectRoom) (string, error) {
	db := connectDB()
	defer db.Close()

	userDB := new(data_scheme.User)
	db.Model(userDB).Where("login = ?", user.Login).Select()
	if userDB.Login == "" {
		return "", errors.New("user doesn't exist")
	}
	if ex := checkIfExists(userDB.Rooms, user.Secret); ex {
		return "", errors.New("user is already in the room")
	}
	roomDB := new(data_scheme.Rooms)
	db.Model(roomDB).Where("secret = ?", user.Secret).Select()
	if roomDB.Secret == "" {
		return "", errors.New("room doesn't exist")
	}
	if ex := checkIfExistsUserMap(roomDB.Users.UsersList, user.Login, user.Username); ex {
		return "", errors.New("user is already in the room")
	}

	userDB.Rooms = append(userDB.Rooms, user.Secret)
	if _, err := db.Model(userDB).Set("rooms = ?rooms").Where("id = ?id").Update(); err != nil {
		return "", errors.New("internal error")
	}

	roomDB.Users.UsersList = append(roomDB.Users.UsersList, data_scheme.UsersMapDB{
		Login:    user.Login,
		Username: user.Username,
	})
	if _, err := db.Model(roomDB).Set("users = ?users").Where("id = ?id").Update(); err != nil {
		return "", errors.New("internal error")
	}
	return roomDB.Secret, nil
}

func GetRoomsDB(user *data_scheme.User) []string {
	db := connectDB()
	defer db.Close()

	userDB := new(data_scheme.User)
	db.Model(userDB).Where("login = ? AND password = ?", user.Login, user.Password).Select()
	if userDB.Login == "" {
		return []string{""}
	}
	return userDB.Rooms
}

func checkIfExists(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}

func checkIfExistsUserMap(users []data_scheme.UsersMapDB, login string, username string) bool {
	for _, rec := range users {
		if rec.Login == login || rec.Username == username {
			return true
		}
	}
	return false
}

func findUserInRoom(login string, room *data_scheme.Rooms) string {

	for _, rec := range room.Users.UsersList {
		if rec.Login == login {
			return rec.Username
		}
	}
	return ""
}
