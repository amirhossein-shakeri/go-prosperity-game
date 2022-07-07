package level

import "github.com/gin-gonic/gin"

type CreateLevelRequest struct {
	// Number uint   `json:"number" form:"number" xml:"number"`
	Note string `json:"note" form:"note" xml:"note"`
	// Items  []item.CreateItemRequest `json:"items" form:"items" xml:"items"`
}

var ErrForbidden = gin.H{"message": "You don't have access to this level, bitch!"}
