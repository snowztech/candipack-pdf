package handlers

import (
	"candipack-pdf/internal/generator"
	"candipack-pdf/internal/lang"
	"candipack-pdf/internal/models"
	"candipack-pdf/internal/parser"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	parser    *parser.HTMLParser
	generator *generator.Generator
}

func New() *Handler {
	return &Handler{
		parser:    parser.NewHTMLParser(),
		generator: generator.NewGenerator(),
	}
}

func (h *Handler) HandleResume() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resumeData models.Resume
		if err := c.ShouldBindJSON(&resumeData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload: " + err.Error()})
			return
		}

		if resumeData.Meta.Template == "" {
			resumeData.Meta.Template = "professional"
		}
		if resumeData.Meta.Lang == "" {
			resumeData.Meta.Lang = "en"
		}

		labels := map[string]string{
			"Education":    lang.Translate(resumeData.Meta.Lang, "Education"),
			"Experiences":  lang.Translate(resumeData.Meta.Lang, "Experiences"),
			"Volunteer":    lang.Translate(resumeData.Meta.Lang, "Volunteer"),
			"Publications": lang.Translate(resumeData.Meta.Lang, "Publications"),
			"Projects":     lang.Translate(resumeData.Meta.Lang, "Projects"),
			"Skills":       lang.Translate(resumeData.Meta.Lang, "Skills"),
			"SoftSkills":   lang.Translate(resumeData.Meta.Lang, "SoftSkills"),
			"Languages":    lang.Translate(resumeData.Meta.Lang, "Languages"),
			"Interests":    lang.Translate(resumeData.Meta.Lang, "Interests"),
			"Profile":      lang.Translate(resumeData.Meta.Lang, "Profile"),
			"Since":        lang.Translate(resumeData.Meta.Lang, "Since"),
			"Certificates": lang.Translate(resumeData.Meta.Lang, "Certificates"),
			"Socials":      lang.Translate(resumeData.Meta.Lang, "Socials"),
		}

		htmlFile, err := h.parser.ParseResume(resumeData.Meta.Template, resumeData, labels)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse template: " + err.Error()})
			return
		}
		defer os.Remove(htmlFile)

		pdf, err := h.generator.GeneratePDF(htmlFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF: " + err.Error()})
			return
		}

		c.Data(http.StatusOK, "application/pdf", pdf)
	}
}

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

		htmlFile, err := h.parser.ParseCoverLetter(coverLetterData.Meta.Template, coverLetterData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse template: " + err.Error()})
			return
		}
		defer os.Remove(htmlFile)

		pdf, err := h.generator.GeneratePDF(htmlFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF: " + err.Error()})
			return
		}

		c.Data(http.StatusOK, "application/pdf", pdf)
	}
}

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
