package handlers

import (
	"candipack-pdf/internal/lang"
	"candipack-pdf/internal/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleCoverLetter() gin.HandlerFunc {
	return func(c *gin.Context) {
		var coverLetterData models.CoverLetter
		if err := c.ShouldBindJSON(&coverLetterData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload: " + err.Error()})
			return
		}

		if coverLetterData.Meta.Template == "" {
			coverLetterData.Meta.Template = "classic"
		}
		if coverLetterData.Meta.Lang == "" {
			coverLetterData.Meta.Lang = "en"
		}

		labels := map[string]string{
			"Subject": lang.Translate(coverLetterData.Meta.Lang, "Subject"),
		}

		htmlFile, err := h.parser.ParseCoverLetter(coverLetterData.Meta.Template, coverLetterData, labels)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse template: " + err.Error()})
			return
		}
		defer func() {
			if err := os.Remove(htmlFile); err != nil {
				log.Printf("Warning: failed to remove temp file: %v", err)
			}
		}()

		pdf, err := h.generator.GeneratePDF(htmlFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF: " + err.Error()})
			return
		}

		c.Data(http.StatusOK, "application/pdf", pdf)
	}
}

func (h *Handler) HandleCoverLetterHTML() gin.HandlerFunc {
	return func(c *gin.Context) {
		var coverLetterData models.CoverLetter
		if err := c.ShouldBindJSON(&coverLetterData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload: " + err.Error()})
			return
		}

		if coverLetterData.Meta.Template == "" {
			coverLetterData.Meta.Template = "classic"
		}
		if coverLetterData.Meta.Lang == "" {
			coverLetterData.Meta.Lang = "en"
		}

		labels := map[string]string{
			"Subject": lang.Translate(coverLetterData.Meta.Lang, "Subject"),
		}

		html, err := h.parser.ParseCoverLetterHTML(coverLetterData.Meta.Template, coverLetterData, labels)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse template: " + err.Error()})
			return
		}

		c.Data(http.StatusOK, "text/html", []byte(html))
	}
}
