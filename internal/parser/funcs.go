package parser

import (
	"html/template"
	"strings"
	"time"

	"candipack-pdf/internal/models"

	"github.com/goodsign/monday"
)

var funcs = template.FuncMap{
	"displayLocation": displayLocation,
	"formatDate":      formatDate,
	"trimURLPrefix":   trimURLPrefix,
	"split":           split,
	"safeURL":         safeURL,
}

func displayLocation(location models.Location) string {
	var parts []string

	if location.City != "" {
		parts = append(parts, location.City)
	}
	if location.Region != "" {
		parts = append(parts, location.Region)
	}
	if location.CountryCode != "" {
		parts = append(parts, location.CountryCode)
	}

	if location.CountryCode != "" {
		parts = append(parts, location.CountryCode)
	}

	return strings.Join(parts, ", ")
}

var dateFormats = []string{
	"2006-01-02",
	"2006-01",
	"January 2 2006",
	"January 2006",
	"2006",
}

func formatDate(layout string, date string, locale string) string {
	if date == "" {
		return ""
	}

	for _, format := range dateFormats {
		t, err := time.Parse(format, date)
		if err == nil {
			return monday.Format(t, layout, monday.Locale(locale))
		}
	}

	return date
}

func trimURLPrefix(url string) string {
	prefixes := []string{"https://www.", "http://www.", "https://", "http://", "www."}
	for _, prefix := range prefixes {
		if strings.HasPrefix(url, prefix) {
			return strings.TrimPrefix(url, prefix)
		}
	}
	return url
}

func split(s string, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}

func safeURL(url string) template.URL {
	return template.URL(url)
}
