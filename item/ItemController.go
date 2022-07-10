package item

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func GetItems(ctx *gin.Context) {
	items := []Item{}
	if err := mgm.Coll(&Item{}).SimpleFind(&items, bson.M{"levelId": ctx.Param("levelId")}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func DeleteItem(ctx *gin.Context) {
	i := Find(ctx.Param("itemId"))
	if i == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	if i.UserID != ctx.GetString("user_id") {
		ctx.AbortWithStatusJSON(http.StatusForbidden, ErrForbidden)
		return
	}
	if err := mgm.Coll(i).Delete(i); err != nil {
		log.Println("Error deleting item", i, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
