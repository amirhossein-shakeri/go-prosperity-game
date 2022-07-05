package level

import (
	"amirhossein-shakeri/go-prosperity-game/item"

	"github.com/kamva/mgm/v3"
)

type Level struct {
	mgm.DefaultModel `bson:",inline"`
	Number           uint        `json:"number" bson:"number"`
	Items            []item.Item `json:"items" bson:"items"`
	Note             string      `json:"note" bson:"note"`
	UserID           string      `json:"userId" bson:"userId"`
}

func New(number uint, items []item.Item, note string) *Level {
	return &Level{Number: number, Items: items, Note: note}
}
