package Controllers

import (
	"github.com/gin-gonic/gin"
)

type ErrorsController struct {
	TemplateHtml
}

func (e *ErrorsController) E404(c *gin.Context) {
	ErrorsController := ErrorsController{}
	ErrorsController.showHtml(c.Writer, nil, "404")
}
