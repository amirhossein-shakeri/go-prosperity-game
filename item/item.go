package item

import "github.com/kamva/mgm/v3"

type Item struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string  `json:"title" bson:"title"`
	Price            float32 `json:"price" bson:"price"`
	URL              string  `json:"url" bson:"url"`
	Description      string  `json:"description" bson:"description"`
}

func New(title string, price float32, url string, desc string) *Item {
	return &Item{Title: title, Price: price, URL: url, Description: desc}
}
