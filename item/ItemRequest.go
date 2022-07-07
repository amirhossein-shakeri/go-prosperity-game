package item

type CreateItemRequest struct {
	Title       string  `json:"title" form:"title" xml:"title" binding:"required"`
	Price       float32 `json:"price" form:"price" xml:"price" binding:"required,gte=0"`
	URL         string  `json:"url" form:"url" xml:"url"`
	Description string  `json:"description" form:"description" xml:"description"`
}
