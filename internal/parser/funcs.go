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
	"lower":           strings.ToLower,
}

func displayLocation(location models.Location) string {
	var parts []string

	if location.Address != "" {
		parts = append(parts, location.Address)
	}
	if location.PostalCode != "" || location.City != "" {
		cityPart := strings.TrimSpace(location.City + " " + location.PostalCode)
		parts = append(parts, cityPart)
	}
	if location.Region != "" {
		parts = append(parts, location.Region)
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

// localeMap maps short language codes to monday ICU locale identifiers.
var localeMap = map[string]monday.Locale{
	"en":    monday.LocaleEnUS,
	"en_US": monday.LocaleEnUS,
	"en_GB": monday.LocaleEnGB,
	"fr":    monday.LocaleFrFR,
	"fr_FR": monday.LocaleFrFR,
	"fr_CA": monday.LocaleFrCA,
	"de":    monday.LocaleDeDE,
	"de_DE": monday.LocaleDeDE,
	"es":    monday.LocaleEsES,
	"es_ES": monday.LocaleEsES,
	"it":    monday.LocaleItIT,
	"it_IT": monday.LocaleItIT,
	"pt":    monday.LocalePtPT,
	"pt_PT": monday.LocalePtPT,
	"pt_BR": monday.LocalePtBR,
	"nl":    monday.LocaleNlNL,
	"nl_NL": monday.LocaleNlNL,
	"pl":    monday.LocalePlPL,
	"pl_PL": monday.LocalePlPL,
	"ru":    monday.LocaleRuRU,
	"ru_RU": monday.LocaleRuRU,
	"tr":    monday.LocaleTrTR,
	"tr_TR": monday.LocaleTrTR,
	"ja":    monday.LocaleJaJP,
	"ja_JP": monday.LocaleJaJP,
	"zh":    monday.LocaleZhCN,
	"zh_CN": monday.LocaleZhCN,
	"ko":    monday.LocaleKoKR,
	"ko_KR": monday.LocaleKoKR,
}

func resolveLocale(locale string) monday.Locale {
	if l, ok := localeMap[locale]; ok {
		return l
	}
	return monday.Locale(locale)
}

func formatDate(layout string, date string, locale string) string {
	if date == "" {
		return ""
	}

	mondayLocale := resolveLocale(locale)
	for _, format := range dateFormats {
		t, err := time.Parse(format, date)
		if err == nil {
			// Choose output format based on input precision
			var outputLayout string
			switch format {
			case "2006": // Year only
				outputLayout = "2006"
			case "2006-01", "January 2006": // Year + month
				outputLayout = "January 2006"
			default: // Full date or others
				outputLayout = layout
			}
			return monday.Format(t, outputLayout, mondayLocale)
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
