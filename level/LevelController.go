package level

import (
	"amirhossein-shakeri/go-prosperity-game/item"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func GetLevels(ctx *gin.Context) {
	levels := []Level{}
	if err := mgm.Coll(&Level{}).SimpleFind(&levels, bson.M{"userId": ctx.GetString("user_id")}); err != nil { // userId
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, levels)
}

func GetLevel(ctx *gin.Context) {
	level := Find(ctx.Param("levelId"))
	if level.UserID != ctx.GetString("user_id") {
		ctx.AbortWithStatusJSON(http.StatusForbidden, ErrForbidden)
		return
	}
	ctx.JSON(http.StatusOK, level)
}

func PostLevel(ctx *gin.Context) {
	req := &CreateLevelRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	count, err := mgm.Coll(&Level{}).CountDocuments(nil, bson.M{"userId": ctx.GetString("user_id")})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create level", "error": err.Error()})
		log.Panicln("Error creating level", req, err, err.Error())
		return
	}
	number := count + 1

	level, err := Create(uint(number), []item.Item{}, req.Note, ctx.GetString("user_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create level", "error": err.Error()})
		log.Panicln("Error creating level", req, err, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, level)
}

func DeleteLevel(ctx *gin.Context) {
	level := Find(ctx.Param("levelId"))
	if level == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	if level.UserID != ctx.GetString("user_id") {
		ctx.AbortWithStatusJSON(http.StatusForbidden, ErrForbidden)
		return
	}
	if err := mgm.Coll(level).Delete(level); err != nil {
		log.Println("Error deleting level", level.ID, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func PostItem(ctx *gin.Context) {
	req := &item.CreateItemRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// fetch the level
	level := Find(req.LevelID)
	if level == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	// check for cost overflow
	level.LoadItemsIfNotLoaded()
	if req.Price+level.ItemsCost() > level.MaxCost() {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf(
			"%v is too expensive($%.2f) to be in level %v(max $%.2f). Remove some items or add an item up to $%.2f",
			req.Title,
			req.Price,
			level.Number,
			level.MaxCost(),
			level.MaxCost()-level.ItemsCost(),
		)})
		return
	}
	// check for title duplicate in all items of this user
	duplicates := []item.Item{}
	if err := mgm.Coll(&item.Item{}).SimpleFind(&duplicates, bson.M{"userId": ctx.GetString("user_id"), "title": req.Title}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Print("Error finding duplicate items", err.Error())
		return
	}
	// fmt.Println("LEN:", len(duplicates), duplicates)
	// return
	if len(duplicates) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("You've already bought %v in level %v", req.Title, Find(duplicates[0].LevelID).Number)})
		return
	}
	// create new item in the level
	newItem, err := level.CreateNewItem(req.Title, req.Price, req.URL, req.Description)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		log.Println("Failed to create new item in level", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, newItem)
}
