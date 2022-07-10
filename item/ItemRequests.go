package item

import "github.com/gin-gonic/gin"

type CreateItemRequest struct {
	Title       string  `json:"title" form:"title" xml:"title" binding:"required"`
	Price       float32 `json:"price" form:"price" xml:"price" binding:"required,gte=0"`
	URL         string  `json:"url" form:"url" xml:"url"`
	Description string  `json:"description" form:"description" xml:"description"`
	LevelID     string  `json:"levelId" form:"levelId" xml:"levelId" binding:"required"`
}

var ErrForbidden = gin.H{"message": "You don't have access to this item, bitch!"}
