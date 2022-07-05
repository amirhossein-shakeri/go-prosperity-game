package level

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func GetLevels(ctx *gin.Context) {
	levels := []Level{}
	if err := mgm.Coll(&Level{}).SimpleFind(&levels, bson.M{}); err != nil { // userId
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, levels)
}

// var levels []Level = []Level{
// 	*New(1, nil, ""),
// 	*New(2, nil, ""),
// }
