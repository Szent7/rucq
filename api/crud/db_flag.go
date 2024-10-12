package crud

import (
	"log"
	"rucq/api/data_scheme"

	"github.com/go-pg/pg/v10"
)

var user1 = data_scheme.User{
	Id:       0,
	Login:    "FSBs1",
	Password: "f;lgkho5",
	Token:    "lD6N1syXDLx3y8ZaBXseWzUrkaKDZya2",
	Rooms: []string{
		"gTEcCjgaQE5ZJuXlJbs0kJMHdcwTA4mZ",
		"DpA0NnpyGs9To0tDfDbY8wKr5Vtn9sGA",
	},
}

var user2 = data_scheme.User{
	Id:       0,
	Login:    "Mayor12mkjg",
	Password: "jghjfgujghj54647bjv57y",
	Token:    "VYk697dLqUuE1fBUWjOfvMBIRlkzjJKZ",
	Rooms: []string{
		"gTEcCjgaQE5ZJuXlJbs0kJMHdcwTA4mZ",
		"DpA0NnpyGs9To0tDfDbY8wKr5Vtn9sGA",
	},
}

var user3 = data_scheme.User{
	Id:       0,
	Login:    "Polkovnyk1971",
	Password: "f;lgkho5",
	Token:    "CgVOtG9Uw6qymuLQPUQ1D3b2Sz6r57oH",
	Rooms: []string{
		"gTEcCjgaQE5ZJuXlJbs0kJMHdcwTA4mZ",
		"DpA0NnpyGs9To0tDfDbY8wKr5Vtn9sGA",
	},
}

var user4 = data_scheme.User{
	Id:       0,
	Login:    "BestPresidentOfRussia2024Ever",
	Password: "24hjccbzdgthtl20",
	Token:    "L6QQNhjHpXGNlY5woNkW9d7tJtbPc0Sp",
	Rooms: []string{
		"gTEcCjgaQE5ZJuXlJbs0kJMHdcwTA4mZ",
		"DpA0NnpyGs9To0tDfDbY8wKr5Vtn9sGA",
	},
}

var user5 = data_scheme.User{
	Id:       0,
	Login:    "jr.Lieutenant",
	Password: "Z7Pqh8v7",
	Token:    "AIZFtSDauMYMcdGfrp6rbeIv0T1fHPdu",
	Rooms: []string{
		"gTEcCjgaQE5ZJuXlJbs0kJMHdcwTA4mZ",
		"DpA0NnpyGs9To0tDfDbY8wKr5Vtn9sGA",
	},
}

var room1 = data_scheme.Rooms{
	Id: 0,
	Users: data_scheme.Users{
		UsersList: []data_scheme.UsersMapDB{
			{
				Login:    "FSBs1",
				Username: "Дружище Смышляков",
			},
			{
				Login:    "Mayor12mkjg",
				Username: "Оперативник Петренко",
			},
			{
				Login:    "Polkovnyk1971",
				Username: "Орел Орлов",
			},
			{
				Login:    "BestPresidentOfRussia2024Ever",
				Username: "Моль",
			},
			{
				Login:    "jr.Lieutenant",
				Username: "Старший Лейтенант",
			},
		},
	},
	Messages: []data_scheme.MesContainer{
		{
			Username: "Дружище Смышляков",
			Message:  "Всем доброе утро! Чего нового?",
		},
		{
			Username: "Оперативник Петренко",
			Message:  "Привет! У меня есть информация о предстоящей сделке.",
		},
		{
			Username: "Орел Орлов",
			Message:  "Что за сделка? Где, когда?",
		},
		{
			Username: "Оперативник Петренко",
			Message:  "Должны встретиться в парке на юго-западе в 14:00.",
		},
		{
			Username: "Моль",
			Message:  "Звучит подозрительно. Кто участники?",
		},
		{
			Username: "Оперативник Петренко",
			Message:  "Информация о клиентах пока что не известна, но один из них может быть небезопасным.",
		},
		{
			Username: "Дружище Смышляков",
			Message:  "Надо направить кого-то на наблюдение. У кого в это время свободно?",
		},
		{
			Username: "Орел Орлов",
			Message:  "Я смогу покрыть эту операцию. Буду на связи.",
		},
		{
			Username: "Моль",
			Message:  "Отлично, но нужно оставаться на дальнем расстоянии.",
		},
		{
			Username: "Орел Орлов",
			Message:  "Если что-то пойдет не так, готовьте наушники!",
		},
		{
			Username: "Моль",
			Message:  "И помните, что у нас есть код на экстренные сообщения: FLAG{G42f12l10}",
		},
		{
			Username: "Дружище Смышляков",
			Message:  "Запомнили, не забудем. Будьте осторожны!",
		},
		{
			Username: "Оперативник Петренко",
			Message:  "Все будет в порядке!",
		},
	},
	Secret: "gTEcCjgaQE5ZJuXlJbs0kJMHdcwTA4mZ",
}

var room2 = data_scheme.Rooms{
	Id: 0,
	Users: data_scheme.Users{
		UsersList: []data_scheme.UsersMapDB{
			{
				Login:    "FSBs1",
				Username: "Комиссар Кузнецов",
			},
			{
				Login:    "Mayor12mkjg",
				Username: "Шерлок Петров",
			},
			{
				Login:    "Polkovnyk1971",
				Username: "Ледяной Мороз",
			},
			{
				Login:    "BestPresidentOfRussia2024Ever",
				Username: "Спикер Думы",
			},
			{
				Login:    "jr.Lieutenant",
				Username: "Старший Лейтенант",
			},
		},
	},
	Messages: []data_scheme.MesContainer{
		{
			Username: "Комиссар Кузнецов",
			Message:  "Доброе утро, команда. Есть новости по делу?",
		},
		{
			Username: "Шерлок Петров",
			Message:  "Привет, Кузнецов. Получил информацию о передвижениях подозреваемых.",
		},
		{
			Username: "Ледяной Мороз",
			Message:  "Мы заметили, что один из них зашел в заброшенное здание на 5-й улице.",
		},
		{
			Username: "Спикер Думы",
			Message:  "Подтверждаю. У нас есть данные о возможной встрече с информантом.",
		},
		{
			Username: "Комиссар Кузнецов",
			Message:  "Есть ли информация о том, кто именно там будет?",
		},
		{
			Username: "Шерлок Петров",
			Message:  "По нашим источникам, это может быть Станислав. Он недавно начал действовать.",
		},
		{
			Username: "Комиссар Кузнецов",
			Message:  "Хорошо. Давайте на всякий случай задействуем резервные силы.",
		},
		{
			Username: "Спикер Думы",
			Message:  "Я подготовлю план выемки. Нужно быть осторожными.",
		},
		{
			Username: "Ледяной Мороз",
			Message:  "Кузнецов, у меня есть флаг на случай, если понадобится дополнительная информация. Он секретен: FLAG{hl;k@*211L:KF}.",
		},
		{
			Username: "Комиссар Кузнецов",
			Message:  "Сообщи об этом только после подтверждения. Безопасность прежде всего.",
		},
		{
			Username: "Спикер Думы",
			Message:  "Давайте все оставшиеся детали обсудим на встрече завтра.",
		},
		{
			Username: "Комиссар Кузнецов",
			Message:  "Договорились. Всем быть на чеку!",
		},
	},
	Secret: "DpA0NnpyGs9To0tDfDbY8wKr5Vtn9sGA",
}

func AddFlag(db *pg.DB) {
	userDB := new([]data_scheme.User)
	db.Model(userDB).Select()

	roomDB := new([]data_scheme.Rooms)
	db.Model(roomDB).Select()
	if len(*roomDB) >= 2 && len(*userDB) >= 5 {
		return
	}

	if _, err := db.Model(&user1).Insert(); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Model(&user2).Insert(); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Model(&user3).Insert(); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Model(&user4).Insert(); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Model(&user5).Insert(); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Model(&room1).Insert(); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Model(&room2).Insert(); err != nil {
		log.Fatal(err)
	}
}
