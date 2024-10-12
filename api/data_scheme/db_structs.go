package data_scheme

type User struct {
	Id       int64
	Login    string   `pg:"type:varchar(64)" json:"login"`
	Password string   `pg:"type:varchar(128)" json:"password"`
	Token    string   `pg:"type:varchar(32)" json:"token"`
	Rooms    []string `pg:",array" json:"rooms"`
}

type Rooms struct {
	Id       int64
	Users    Users          `pg:"type:json"`
	Messages []MesContainer `pg:"type:json"`
	Secret   string         `pg:"type:varchar(32)"`
}

type Messages struct {
	MessageList []MesContainer `json:"messagelist"`
}

type MessagesSend struct {
	Username    string         `json:"username"`
	MessageList []MesContainer `json:"messagelist"`
	Error       string         `json:"error"`
}

type MessagesGet struct {
	Secret string `json:"secret"`
	RoomId string `json:"roomid"`
}

type MesContainer struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type MesContainerSend struct {
	Message string `json:"message"`
	Secret  string `json:"secret"`
	RoomId  string `json:"roomid"`
}

type Users struct {
	UsersList []UsersMapDB `json:"userslist"`
}

type UsersMapDB struct {
	Login    string `json:"login"`
	Username string `json:"username"`
}

type UsersMap struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UsersMapConnectRoom struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Username string `json:"username"`
	Secret   string `json:"secret"`
}
