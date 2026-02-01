package handlers

import (
	"candipack-pdf/internal/lang"
	"candipack-pdf/internal/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

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
			"Education":      lang.Translate(resumeData.Meta.Lang, "Education"),
			"Experiences":    lang.Translate(resumeData.Meta.Lang, "Experiences"),
			"Volunteer":      lang.Translate(resumeData.Meta.Lang, "Volunteer"),
			"Publications":   lang.Translate(resumeData.Meta.Lang, "Publications"),
			"Projects":       lang.Translate(resumeData.Meta.Lang, "Projects"),
			"Skills":         lang.Translate(resumeData.Meta.Lang, "Skills"),
			"SoftSkills":     lang.Translate(resumeData.Meta.Lang, "SoftSkills"),
			"Languages":      lang.Translate(resumeData.Meta.Lang, "Languages"),
			"Interests":      lang.Translate(resumeData.Meta.Lang, "Interests"),
			"Profile":        lang.Translate(resumeData.Meta.Lang, "Profile"),
			"Since":          lang.Translate(resumeData.Meta.Lang, "Since"),
			"Certificates":   lang.Translate(resumeData.Meta.Lang, "Certificates"),
			"Socials":        lang.Translate(resumeData.Meta.Lang, "Socials"),
			"AdditionalInfo": lang.Translate(resumeData.Meta.Lang, "AdditionalInfo"),
			"References":     lang.Translate(resumeData.Meta.Lang, "References"),
		}

		htmlFile, err := h.parser.ParseResume(resumeData.Meta.Template, resumeData, labels)
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

func (h *Handler) HandleResumeHTML() gin.HandlerFunc {
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
			"Education":      lang.Translate(resumeData.Meta.Lang, "Education"),
			"Experiences":    lang.Translate(resumeData.Meta.Lang, "Experiences"),
			"Volunteer":      lang.Translate(resumeData.Meta.Lang, "Volunteer"),
			"Publications":   lang.Translate(resumeData.Meta.Lang, "Publications"),
			"Projects":       lang.Translate(resumeData.Meta.Lang, "Projects"),
			"Skills":         lang.Translate(resumeData.Meta.Lang, "Skills"),
			"SoftSkills":     lang.Translate(resumeData.Meta.Lang, "SoftSkills"),
			"Languages":      lang.Translate(resumeData.Meta.Lang, "Languages"),
			"Interests":      lang.Translate(resumeData.Meta.Lang, "Interests"),
			"Profile":        lang.Translate(resumeData.Meta.Lang, "Profile"),
			"Since":          lang.Translate(resumeData.Meta.Lang, "Since"),
			"Certificates":   lang.Translate(resumeData.Meta.Lang, "Certificates"),
			"Socials":        lang.Translate(resumeData.Meta.Lang, "Socials"),
			"AdditionalInfo": lang.Translate(resumeData.Meta.Lang, "AdditionalInfo"),
			"References":     lang.Translate(resumeData.Meta.Lang, "References"),
		}

		html, err := h.parser.ParseResumeHTML(resumeData.Meta.Template, resumeData, labels)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse template: " + err.Error()})
			return
		}

		c.Data(http.StatusOK, "text/html", []byte(html))
	}
}
