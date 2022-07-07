package level

import "amirhossein-shakeri/go-prosperity-game/item"

type CreateLevelRequest struct {
	Number string                   `json:"number" form:"number" xml:"number"`
	Note   string                   `json:"note" form:"note" xml:"note"`
	Items  []item.CreateItemRequest `json:"items" form:"items" xml:"items"`
}
