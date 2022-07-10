package level

import (
	"amirhossein-shakeri/go-prosperity-game/item"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type Level struct {
	mgm.DefaultModel `bson:",inline"`
	Number           uint        `json:"number" bson:"number"`
	Items            []item.Item `json:"items"` // use a method Items()  bson:"items"
	Note             string      `json:"note" bson:"note"`
	UserID           string      `json:"userId" bson:"userId"`
}

func (l *Level) ItemsCost() float32 {
	var sum float32
	for _, item := range l.Items {
		sum += item.Price
	}
	return sum
}

func (l *Level) MaxCost() float32 {
	return float32(l.Number * 1000) // could have different steps by the way ...
}

func (l *Level) LoadItems() []item.Item {
	items := []item.Item{}
	if err := mgm.Coll(&item.Item{}).SimpleFind(&items, bson.M{"levelId": l.ID.Hex()}); err != nil {
		log.Println("Failed getting items", l.ID.Hex(), err.Error())
		return nil
	}
	l.Items = items
	return items
}

func (l *Level) LoadItemsIfNotLoaded() []item.Item {
	if len(l.Items) > 0 {
		return l.Items
	}
	return l.LoadItems()
}

func (l *Level) GenerateNewItemOrder() uint {
	// load items
	// find the greatest order
	// return greatest order + 1
	return uint(0)
}

func (l *Level) CreateNewItem(title string, price float32, url, desc string) (*item.Item, error) {
	newItem, err := item.Create(title, price, url, desc, l.GenerateNewItemOrder(), l.ID.Hex(), l.UserID)
	if err != nil {
		log.Println("Error creating new item in level", l.ID.Hex(), err.Error())
		return nil, err
	}
	l.Items = append(l.Items, *newItem)
	return newItem, nil
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
