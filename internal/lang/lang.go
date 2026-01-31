package lang

var translations = map[string]map[string]string{
	"en": {
		"Education":    "Education",
		"Experiences":  "Professional Experience",
		"Volunteer":    "Volunteer Work",
		"Languages":    "Languages",
		"Skills":       "Skills",
		"SoftSkills":   "Soft Skills",
		"Publications": "Publications",
		"Projects":     "Projects",
		"Socials":      "Contact",
		"Interests":    "Interests",
		"Profile":      "Profile",
		"Since":        "Since",
		"Certificates": "Certifications",
	},
	"fr": {
		"Education":    "Formation",
		"Experiences":  "Expériences Professionnelles",
		"Volunteer":    "Bénévolat",
		"Languages":    "Langues",
		"Skills":       "Compétences",
		"SoftSkills":   "Compétences Douces",
		"Publications": "Publications",
		"Projects":     "Projets",
		"Socials":      "Contact",
		"Interests":    "Centres d'Intérêt",
		"Profile":      "Profil",
		"Since":        "Depuis",
		"Certificates": "Certifications",
	},
}

func Translate(lang, key string) string {
	if labels, ok := translations[lang]; ok {
		if label, ok := labels[key]; ok {
			return label
		}
	}
	return translations["en"][key]
}
