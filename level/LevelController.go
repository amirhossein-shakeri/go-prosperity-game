package level

import (
	"amirhossein-shakeri/go-prosperity-game/item"
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
