package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleTemplates() gin.HandlerFunc {
	return func(c *gin.Context) {
		templates := map[string]interface{}{
			"cv": []map[string]string{
				{"id": "professional", "name": "Professional", "description": "2-column layout with sidebar"},
				{"id": "simple", "name": "Simple", "description": "1-column minimalist layout"},
			},
			"coverLetter": []map[string]string{
				{"id": "classic", "name": "Classic", "description": "Traditional letter format"},
			},
		}
		c.JSON(http.StatusOK, templates)
	}
}