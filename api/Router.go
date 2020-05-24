package api

import (
	"gindemo/api/Controllers"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	Controllers.History(engine)
}
