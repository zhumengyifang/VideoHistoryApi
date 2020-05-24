package api

import (
	"gindemo/api/Controllers"
	"github.com/gin-gonic/gin"
)

func router(engine *gin.Engine) {
	Controllers.History(engine)
}
