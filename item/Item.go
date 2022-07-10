package item

import (
	"log"

	"github.com/kamva/mgm/v3"
)

type Item struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string  `json:"title" bson:"title"`
	Price            float32 `json:"price" bson:"price"`
	URL              string  `json:"url" bson:"url"`
	Description      string  `json:"description" bson:"description"`
	Order            uint    `json:"order" bson:"order"`
	LevelID          string  `json:"levelId" bson:"levelId"`
	UserID           string  `json:"userId" bson:"userId"`
}

func New(title string, price float32, url, desc string, order uint, levelId, userId string) *Item {
	return &Item{Title: title, Price: price, URL: url, Description: desc, Order: order, LevelID: levelId, UserID: userId}
}

func Create(title string, price float32, url, desc string, order uint, levelId, userId string) (*Item, error) {
	item := New(title, price, url, desc, order, levelId, userId)
	return item, mgm.Coll(item).Create(item)
}

func Find(id string) *Item {
	item := &Item{}
	if err := mgm.Coll(item).FindByID(id, item); err != nil {
		log.Println("Error finding item by ID", id, err.Error())
		return nil
	}
	return item
}
