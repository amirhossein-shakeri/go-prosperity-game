package level

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLevels(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, levels)
}

var levels []Level = []Level{
	*New(1, nil, ""),
	*New(2, nil, ""),
}
