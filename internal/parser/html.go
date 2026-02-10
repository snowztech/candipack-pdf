package parser

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"candipack-pdf/internal/models"
)

const templatesPath = "./templates"

type HTMLParser struct{}

func NewHTMLParser() *HTMLParser {
	return &HTMLParser{}
}

type resumeTemplateData struct {
	models.Resume
	Labels map[string]string
}

type coverLetterTemplateData struct {
	models.CoverLetter
	Labels map[string]string
}

func (p *HTMLParser) ParseResume(templateName string, resume models.Resume, labels map[string]string) (string, error) {
	templatePath := filepath.Join(templatesPath, "cv", templateName, "template.gohtml")

	tmpl, err := template.New("template.gohtml").Funcs(funcs).ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	data := resumeTemplateData{
		Resume: resume,
		Labels: labels,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	tempFile, err := os.CreateTemp("", "resume-*.html")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	defer func() {
		if err := tempFile.Close(); err != nil {
			log.Printf("Warning: failed to close temp file: %v", err)
		}
	}()

	if _, err := tempFile.Write(buf.Bytes()); err != nil {
		return "", fmt.Errorf("failed to write to temp file: %w", err)
	}

	return tempFile.Name(), nil
}

func (p *HTMLParser) ParseCoverLetter(templateName string, coverLetter models.CoverLetter, labels map[string]string) (string, error) {
	templatePath := filepath.Join(templatesPath, "cover-letter", templateName, "template.gohtml")

	tmpl, err := template.New("template.gohtml").Funcs(funcs).ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	data := coverLetterTemplateData{
		CoverLetter: coverLetter,
		Labels:      labels,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	tempFile, err := os.CreateTemp("", "coverletter-*.html")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	defer func() {
		if err := tempFile.Close(); err != nil {
			log.Printf("Warning: failed to close temp file: %v", err)
		}
	}()

	if _, err := tempFile.Write(buf.Bytes()); err != nil {
		return "", fmt.Errorf("failed to write to temp file: %w", err)
	}

	return tempFile.Name(), nil
}

func (p *HTMLParser) ParseResumeHTML(templateName string, resume models.Resume, labels map[string]string) (string, error) {
	templatePath := filepath.Join(templatesPath, "cv", templateName, "template.gohtml")

	tmpl, err := template.New("template.gohtml").Funcs(funcs).ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	data := resumeTemplateData{
		Resume: resume,
		Labels: labels,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

func (p *HTMLParser) ParseCoverLetterHTML(templateName string, coverLetter models.CoverLetter, labels map[string]string) (string, error) {
	templatePath := filepath.Join(templatesPath, "cover-letter", templateName, "template.gohtml")

	tmpl, err := template.New("template.gohtml").Funcs(funcs).ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	data := coverLetterTemplateData{
		CoverLetter: coverLetter,
		Labels:      labels,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}
