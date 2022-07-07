package level

import (
	"amirhossein-shakeri/go-prosperity-game/item"
	"log"

	"github.com/kamva/mgm/v3"
)

type Level struct {
	mgm.DefaultModel `bson:",inline"`
	Number           uint        `json:"number" bson:"number"`
	Items            []item.Item `json:"items" bson:"items"`
	Note             string      `json:"note" bson:"note"`
	UserID           string      `json:"userId" bson:"userId"`
}

func New(number uint, items []item.Item, note, userId string) *Level {
	return &Level{Number: number, Items: items, Note: note, UserID: userId}
}

func Create(number uint, items []item.Item, note, userId string) (*Level, error) {
	level := New(number, items, note, userId)
	return level, mgm.Coll(level).Create(level)
}

func Find(id string) *Level {
	level := &Level{}
	if err := mgm.Coll(level).FindByID(id, level); err != nil {
		log.Println("Error finding level by ID", id, err, err.Error())
		return nil
	}
	return level
}
