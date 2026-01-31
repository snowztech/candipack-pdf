# candipack-pdf

Service de génération de PDF pour CV et lettres de motivation.

## Features

- Génération de CV en PDF à partir de données JSON
- Génération de lettres de motivation en PDF
- Templates HTML/CSS personnalisables
- API REST simple
- Support multilingue (fr, en)

## Quick Start

```bash
# Clone
git clone https://github.com/lucasnevespereira/candipack-pdf.git
cd candipack-pdf

# Run
go run ./cmd/server

# Test
curl -X POST http://localhost:9000/resume \
  -H "Content-Type: application/json" \
  -d @examples/resume.json \
  -o cv.pdf
```

## API

| Méthode | Endpoint        | Description                            |
| ------- | --------------- | -------------------------------------- |
| `POST`  | `/resume`       | Génère un CV en PDF                    |
| `POST`  | `/cover-letter` | Génère une lettre de motivation en PDF |
| `GET`   | `/templates`    | Liste les templates disponibles        |
| `GET`   | `/health`       | Health check                           |

### Headers

```
Content-Type: application/json
X-API-Key: your-api-key (optionnel, si configuré)
```

## Templates disponibles

### CV

- `professional` - 2 colonnes avec sidebar (défaut)
- `simple` - 1 colonne, minimaliste

### Cover Letter

- `classic` - Format lettre traditionnel

---

# MVP Specification

## Structure du projet

```
candipack-pdf/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── internal/
│   ├── handlers/
│   │   ├── resume.go                # POST /resume
│   │   ├── coverletter.go       # POST /cover-letter
│   │   └── templates.go         # GET /templates
│   ├── generator/
│   │   └── pdf.go               # HTML to PDF (chromedp)
│   ├── middleware/
│   │   └── auth.go              # API Key auth
│   └── models/
│       ├── cv.go                # CV JSON structure
│       └── coverletter.go       # Cover letter JSON structure
├── templates/
│   ├── cv/
│   │   ├── professional/
│   │   │   └── template.html
│   │   └── simple/
│   │       └── template.html
│   └── cover-letter/
│       └── classic/
│           └── template.html
├── examples/
│   ├── cv.json
│   └── cover-letter.json
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

---

## JSON Schemas

### CV (`POST /resume`)

```json
{
  "template": "professional",
  "locale": "fr",
  "data": {
    "fullName": "Frederica Pereira",
    "jobTitle": "Assistante Administrative",
    "photo": "https://example.com/photo.jpg",
    "summary": "Mon expérience de 3 ans dans le conseil et l'accompagnement m'a permis de développer de solides compétences en accueil, écoute active et gestion de dossiers.",
    "contact": {
      "phone": "+41 79 823 31 18",
      "email": "frederica@example.com",
      "linkedin": "linkedin.com/in/frederica"
    },
    "address": {
      "street": "Route Cantonale 16",
      "city": "3975 Randogne"
    },
    "skills": [
      "Accueil & Service client",
      "Gestion administrative",
      "Suite Microsoft Office",
      "Rédaction de courriers"
    ],
    "languages": [
      { "name": "Français", "level": "Bilingue" },
      { "name": "Portugais", "level": "Langue maternelle" },
      { "name": "Anglais", "level": "Intermédiaire (B1)" }
    ],
    "experiences": [
      {
        "title": "Réceptionniste",
        "company": "Hôtel & Spa Montana",
        "location": "Crans-Montana",
        "period": "2022 - 2025",
        "missions": [
          "Accueil des clients et gestion des réservations",
          "Coordination avec les équipes internes",
          "Gestion des réclamations et demandes spéciales"
        ]
      },
      {
        "title": "Assistante Administrative",
        "company": "Fiduciaire ABC",
        "location": "Sierre",
        "period": "2020 - 2022",
        "missions": [
          "Gestion du courrier et des appels",
          "Préparation de documents comptables",
          "Classement et archivage"
        ]
      }
    ],
    "education": [
      {
        "degree": "CFC Employée de commerce",
        "school": "ECCG Sierre",
        "period": "2017 - 2020",
        "status": "Obtenu"
      }
    ],
    "additionalInfo": {
      "availability": "immédiate",
      "activityRate": "80-100%",
      "workPermit": "Permis B",
      "references": "Disponibles sur demande"
    }
  }
}
```

### Cover Letter (`POST /cover-letter`)

```json
{
  "template": "classic",
  "locale": "fr",
  "data": {
    "senderName": "Frederica Pereira",
    "senderAddress": "Route Cantonale 16, 3975 Randogne",
    "senderPhone": "+41 79 823 31 18",
    "senderEmail": "frederica@example.com",
    "recipientCompany": "AZ Conseils Sàrl",
    "recipientAddress": "Rue de l'Industrie 10, 1950 Sion",
    "date": "Randogne, le 31 janvier 2026",
    "subject": "Candidature au poste de Collaboratrice Administrative",
    "salutation": "Madame, Monsieur,",
    "paragraphs": [
      "Je vous fais parvenir ma candidature pour le poste de Collaboratrice Administrative au sein de votre entreprise AZ Conseils Sàrl.",
      "Fort de mon expérience de trois ans dans l'accueil et l'administration, j'ai développé des compétences solides en gestion de dossiers, coordination d'équipes et relation client. Mon passage à l'Hôtel & Spa Montana m'a permis de perfectionner mon sens de l'organisation et ma capacité à gérer plusieurs tâches simultanément.",
      "Votre entreprise, reconnue pour son dynamisme dans le secteur de la construction et de la promotion immobilière, correspond parfaitement à mes aspirations professionnelles. Je suis convaincue que mon profil polyvalent et mon engagement seraient des atouts pour votre équipe.",
      "Je serais ravie de vous rencontrer afin de vous exposer plus en détail mes motivations et mon parcours."
    ],
    "closing": "Je vous prie de recevoir, Madame, Monsieur, mes salutations distinguées.",
    "signature": "Frederica Pereira"
  }
}
```

---

## Templates HTML

### CV Professional (2 colonnes)

```html
<!DOCTYPE html>
<html lang="{{.Locale}}">
  <head>
    <meta charset="UTF-8" />
    <link
      href="https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;500;600;700&display=swap"
      rel="stylesheet"
    />
    <style>
      @page {
        size: A4;
        margin: 0;
      }
      * {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
      }
      html,
      body {
        width: 210mm;
        height: 297mm;
        font-family: "Open Sans", sans-serif;
        font-size: 11pt;
        line-height: 1.4;
        color: #333;
      }
      .container {
        display: flex;
        height: 100%;
      }

      /* Sidebar */
      .sidebar {
        width: 38%;
        background: #3d5167;
        color: white;
        padding: 30px 24px;
      }
      .photo {
        width: 120px;
        height: 120px;
        border-radius: 50%;
        object-fit: cover;
        border: 3px solid rgba(255, 255, 255, 0.2);
        margin: 0 auto 24px;
        display: block;
      }
      .sidebar h2 {
        font-size: 14pt;
        font-weight: 600;
        margin-bottom: 12px;
        padding-bottom: 6px;
        border-bottom: 1px solid rgba(255, 255, 255, 0.2);
      }
      .sidebar section {
        margin-bottom: 24px;
      }
      .sidebar p,
      .sidebar li {
        font-size: 10pt;
        line-height: 1.6;
      }
      .sidebar ul {
        list-style: none;
      }
      .sidebar li {
        margin-bottom: 4px;
      }
      .sidebar li::before {
        content: "• ";
        opacity: 0.7;
      }
      .lang-item {
        display: flex;
        justify-content: space-between;
        margin-bottom: 4px;
      }
      .lang-level {
        opacity: 0.8;
        font-size: 9pt;
      }

      /* Main */
      .main {
        width: 62%;
        padding: 30px 32px;
      }
      .header {
        text-align: center;
        margin-bottom: 20px;
      }
      .name {
        font-size: 26pt;
        font-weight: 400;
        color: #2d3e50;
        margin-bottom: 4px;
      }
      .job-title {
        font-size: 14pt;
        font-weight: 600;
        color: #4a5568;
        margin-bottom: 12px;
      }
      .summary {
        font-size: 10pt;
        color: #4a5568;
        text-align: justify;
        line-height: 1.6;
      }
      .main h2 {
        font-size: 13pt;
        font-weight: 600;
        color: #2d3e50;
        padding-bottom: 6px;
        border-bottom: 2px solid #2d3e50;
        margin-bottom: 14px;
      }
      .main section {
        margin-bottom: 20px;
      }

      /* Experience */
      .experience {
        margin-bottom: 14px;
        position: relative;
        padding-left: 16px;
      }
      .experience::before {
        content: "";
        position: absolute;
        left: 0;
        top: 6px;
        width: 8px;
        height: 8px;
        border: 2px solid #6b7280;
        border-radius: 50%;
        background: white;
      }
      .exp-title {
        font-size: 12pt;
        font-weight: 600;
        color: #2d3e50;
      }
      .exp-company {
        font-size: 10pt;
        color: #e07c4c;
        font-weight: 500;
        margin-bottom: 4px;
      }
      .exp-missions {
        padding-left: 16px;
        font-size: 9pt;
        color: #4a5568;
      }
      .exp-missions li {
        margin-bottom: 2px;
      }

      /* Education */
      .edu-item {
        font-size: 10pt;
        color: #4a5568;
        margin-bottom: 4px;
      }
      .edu-degree {
        font-weight: 500;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <aside class="sidebar">
        {{if .Data.Photo}}
        <img src="{{.Data.Photo}}" alt="Photo" class="photo" />
        {{end}}

        <section>
          <h2>Contact</h2>
          <p>{{.Data.Contact.Phone}}</p>
          <p>{{.Data.Contact.Email}}</p>
          {{if .Data.Contact.LinkedIn}}
          <p>{{.Data.Contact.LinkedIn}}</p>
          {{end}}
          <p style="margin-top: 8px;">{{.Data.Address.Street}}</p>
          <p>{{.Data.Address.City}}</p>
          {{if .Data.AdditionalInfo.WorkPermit}}
          <p style="margin-top: 8px;">
            Permis: {{.Data.AdditionalInfo.WorkPermit}}
          </p>
          {{end}}
        </section>

        <section>
          <h2>Compétences</h2>
          <ul>
            {{range .Data.Skills}}
            <li>{{.}}</li>
            {{end}}
          </ul>
        </section>

        <section>
          <h2>Langues</h2>
          {{range .Data.Languages}}
          <div class="lang-item">
            <span>{{.Name}}</span>
            <span class="lang-level">{{.Level}}</span>
          </div>
          {{end}}
        </section>

        {{if .Data.AdditionalInfo}}
        <section>
          <h2>Infos</h2>
          {{if .Data.AdditionalInfo.Availability}}
          <p>
            <strong>Disponibilité:</strong>
            {{.Data.AdditionalInfo.Availability}}
          </p>
          {{end}} {{if .Data.AdditionalInfo.ActivityRate}}
          <p><strong>Taux:</strong> {{.Data.AdditionalInfo.ActivityRate}}</p>
          {{end}} {{if .Data.AdditionalInfo.References}}
          <p style="margin-top: 8px; font-size: 9pt; opacity: 0.9;">
            {{.Data.AdditionalInfo.References}}
          </p>
          {{end}}
        </section>
        {{end}}
      </aside>

      <main class="main">
        <header class="header">
          <h1 class="name">{{.Data.FullName}}</h1>
          {{if .Data.JobTitle}}
          <p class="job-title">{{.Data.JobTitle}}</p>
          {{end}}
          <p class="summary">{{.Data.Summary}}</p>
        </header>

        <section>
          <h2>Expériences Professionnelles</h2>
          {{range .Data.Experiences}}
          <div class="experience">
            <p class="exp-title">{{.Title}}</p>
            <p class="exp-company">
              {{.Company}} – {{.Location}} ({{.Period}})
            </p>
            <ul class="exp-missions">
              {{range .Missions}}
              <li>{{.}}</li>
              {{end}}
            </ul>
          </div>
          {{end}}
        </section>

        <section>
          <h2>Formation</h2>
          {{range .Data.Education}}
          <p class="edu-item">
            <span class="edu-degree">{{.Degree}}</span> – {{.School}} {{if
            .Status}}({{.Status}}){{end}}
          </p>
          {{end}}
        </section>
      </main>
    </div>
  </body>
</html>
```

### CV Simple (1 colonne)

```html
<!DOCTYPE html>
<html lang="{{.Locale}}">
  <head>
    <meta charset="UTF-8" />
    <link
      href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap"
      rel="stylesheet"
    />
    <style>
      @page {
        size: A4;
        margin: 0;
      }
      * {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
      }
      html,
      body {
        width: 210mm;
        height: 297mm;
        font-family: "Inter", sans-serif;
        font-size: 10pt;
        line-height: 1.5;
        color: #1a1a1a;
      }
      .cv {
        padding: 40px 50px;
      }

      /* Header */
      .header {
        border-bottom: 2px solid #1a1a1a;
        padding-bottom: 16px;
        margin-bottom: 24px;
      }
      .name {
        font-size: 28pt;
        font-weight: 700;
        margin-bottom: 4px;
      }
      .job-title {
        font-size: 14pt;
        color: #666;
        margin-bottom: 12px;
      }
      .contact-row {
        display: flex;
        flex-wrap: wrap;
        gap: 16px;
        font-size: 9pt;
        color: #444;
      }
      .contact-row span {
        display: flex;
        align-items: center;
        gap: 4px;
      }

      /* Sections */
      section {
        margin-bottom: 20px;
      }
      h2 {
        font-size: 11pt;
        font-weight: 700;
        text-transform: uppercase;
        letter-spacing: 1px;
        color: #1a1a1a;
        margin-bottom: 10px;
        padding-bottom: 4px;
        border-bottom: 1px solid #ddd;
      }

      /* Summary */
      .summary {
        font-size: 10pt;
        color: #333;
        line-height: 1.6;
      }

      /* Experience */
      .exp-item {
        margin-bottom: 14px;
      }
      .exp-header {
        display: flex;
        justify-content: space-between;
        margin-bottom: 2px;
      }
      .exp-title {
        font-weight: 600;
      }
      .exp-period {
        color: #666;
        font-size: 9pt;
      }
      .exp-company {
        color: #666;
        font-size: 9pt;
        margin-bottom: 4px;
      }
      .exp-missions {
        padding-left: 18px;
        font-size: 9pt;
        color: #444;
      }
      .exp-missions li {
        margin-bottom: 2px;
      }

      /* Education */
      .edu-item {
        margin-bottom: 6px;
      }
      .edu-degree {
        font-weight: 500;
      }

      /* Skills & Languages */
      .tags {
        display: flex;
        flex-wrap: wrap;
        gap: 8px;
      }
      .tag {
        background: #f3f4f6;
        padding: 4px 10px;
        border-radius: 4px;
        font-size: 9pt;
      }
      .lang-grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: 8px;
      }
      .lang-item {
        font-size: 9pt;
      }
      .lang-name {
        font-weight: 500;
      }
      .lang-level {
        color: #666;
      }
    </style>
  </head>
  <body>
    <div class="cv">
      <header class="header">
        <h1 class="name">{{.Data.FullName}}</h1>
        {{if .Data.JobTitle}}
        <p class="job-title">{{.Data.JobTitle}}</p>
        {{end}}
        <div class="contact-row">
          <span>{{.Data.Contact.Email}}</span>
          <span>{{.Data.Contact.Phone}}</span>
          <span>{{.Data.Address.City}}</span>
          {{if .Data.AdditionalInfo.WorkPermit}}
          <span>{{.Data.AdditionalInfo.WorkPermit}}</span>
          {{end}}
        </div>
      </header>

      {{if .Data.Summary}}
      <section>
        <h2>Profil</h2>
        <p class="summary">{{.Data.Summary}}</p>
      </section>
      {{end}}

      <section>
        <h2>Expériences</h2>
        {{range .Data.Experiences}}
        <div class="exp-item">
          <div class="exp-header">
            <span class="exp-title">{{.Title}}</span>
            <span class="exp-period">{{.Period}}</span>
          </div>
          <p class="exp-company">{{.Company}}, {{.Location}}</p>
          <ul class="exp-missions">
            {{range .Missions}}
            <li>{{.}}</li>
            {{end}}
          </ul>
        </div>
        {{end}}
      </section>

      <section>
        <h2>Formation</h2>
        {{range .Data.Education}}
        <div class="edu-item">
          <span class="edu-degree">{{.Degree}}</span> – {{.School}} {{if
          .Period}}({{.Period}}){{end}}
        </div>
        {{end}}
      </section>

      <section>
        <h2>Compétences</h2>
        <div class="tags">
          {{range .Data.Skills}}
          <span class="tag">{{.}}</span>
          {{end}}
        </div>
      </section>

      <section>
        <h2>Langues</h2>
        <div class="lang-grid">
          {{range .Data.Languages}}
          <div class="lang-item">
            <span class="lang-name">{{.Name}}</span>
            <span class="lang-level"> – {{.Level}}</span>
          </div>
          {{end}}
        </div>
      </section>
    </div>
  </body>
</html>
```

### Cover Letter Classic

```html
<!DOCTYPE html>
<html lang="{{.Locale}}">
  <head>
    <meta charset="UTF-8" />
    <link
      href="https://fonts.googleapis.com/css2?family=EB+Garamond:ital,wght@0,400;0,600;1,400&display=swap"
      rel="stylesheet"
    />
    <style>
      @page {
        size: A4;
        margin: 0;
      }
      * {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
      }
      html,
      body {
        width: 210mm;
        height: 297mm;
        font-family: "EB Garamond", Georgia, serif;
        font-size: 12pt;
        line-height: 1.7;
        color: #1a1a1a;
      }
      .letter {
        padding: 50px 60px;
      }

      .header {
        text-align: center;
        margin-bottom: 40px;
      }
      .sender-name {
        font-size: 22pt;
        font-weight: 400;
        letter-spacing: 6px;
        text-transform: uppercase;
        margin-bottom: 10px;
      }
      .sender-contact {
        font-size: 10pt;
        font-style: italic;
        color: #555;
      }

      .date {
        text-align: right;
        font-variant: small-caps;
        margin-bottom: 30px;
      }

      .recipient {
        margin-bottom: 24px;
        font-size: 11pt;
      }
      .subject {
        font-weight: 600;
        margin-bottom: 20px;
      }
      .salutation {
        margin-bottom: 16px;
      }

      .paragraph {
        text-align: justify;
        text-indent: 2em;
        margin-bottom: 12px;
      }

      .closing {
        margin-top: 28px;
        margin-bottom: 50px;
      }
      .signature {
        font-weight: 600;
      }
    </style>
  </head>
  <body>
    <div class="letter">
      <header class="header">
        <h1 class="sender-name">{{.Data.SenderName}}</h1>
        <p class="sender-contact">
          {{.Data.SenderAddress}}<br />
          {{.Data.SenderPhone}} · {{.Data.SenderEmail}}
        </p>
      </header>

      <p class="date">{{.Data.Date}}</p>

      {{if or .Data.RecipientCompany .Data.RecipientAddress}}
      <div class="recipient">
        {{if .Data.RecipientCompany}}
        <p>{{.Data.RecipientCompany}}</p>
        {{end}} {{if .Data.RecipientAddress}}
        <p>{{.Data.RecipientAddress}}</p>
        {{end}}
      </div>
      {{end}}

      <p class="subject">{{.Data.Subject}}</p>
      <p class="salutation">{{.Data.Salutation}}</p>

      {{range .Data.Paragraphs}}
      <p class="paragraph">{{.}}</p>
      {{end}}

      <p class="closing">{{.Data.Closing}}</p>
      <p class="signature">{{.Data.Signature}}</p>
    </div>
  </body>
</html>
```

---

## Code Go MVP

### main.go

```go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/lucasnevespereira/candipack-pdf/internal/handlers"
	"github.com/lucasnevespereira/candipack-pdf/internal/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("POST /resume", handlers.HandleCV)
	mux.HandleFunc("POST /cover-letter", handlers.HandleCoverLetter)
	mux.HandleFunc("GET /templates", handlers.HandleTemplates)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// Middleware
	handler := middleware.CORS(middleware.APIKey(mux))

	log.Printf("Server running on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
```

### Middleware

```go
// internal/middleware/middleware.go
package middleware

import (
	"net/http"
	"os"
)

func APIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := os.Getenv("API_KEY")
		if key != "" && r.Header.Get("X-API-Key") != key {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
```

---

## Dockerfile

```dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o candipack-pdf ./cmd/server

FROM alpine:3.19
RUN apk add --no-cache chromium font-noto
WORKDIR /app
COPY --from=builder /app/candipack-pdf .
COPY --from=builder /app/templates ./templates
ENV CHROME_PATH=/usr/bin/chromium-browser
ENV PORT=9000
EXPOSE 9000
CMD ["./candipack-pdf"]
```

---

## Variables d'environnement

```bash
PORT=9000
API_KEY=sk_candipack_xxxxxxxx  # Optionnel
TEMPLATES_PATH=./templates
CHROME_PATH=/usr/bin/chromium-browser
```

---

## MVP Checklist

### Setup

- [ ] Créer le repo `candipack-pdf`
- [ ] Initialiser `go mod init github.com/lucasnevespereira/candipack-pdf`
- [ ] Créer la structure de dossiers

### Core

- [ ] Implémenter `POST /resume`
- [ ] Implémenter `POST /cover-letter`
- [ ] Implémenter `GET /templates`
- [ ] Implémenter génération PDF (chromedp)

### Templates

- [ ] Template CV `professional` (2 colonnes)
- [ ] Template CV `simple` (1 colonne)
- [ ] Template Cover Letter `classic`

### Infra

- [ ] Dockerfile
- [ ] docker-compose.yml (dev)
- [ ] README.md

### Tests

- [ ] Fichier exemple `examples/resume.json`
- [ ] Fichier exemple `examples/cover-letter.json`
- [ ] Tester génération CV
- [ ] Tester génération Cover Letter

---

## Commandes

```bash
# Dev
go run ./cmd/server

# Build
go build -o candipack-pdf ./cmd/server

# Docker
docker build -t candipack-pdf .
docker run -p 9000:9000 candipack-pdf

# Test CV
curl -X POST http://localhost:9000/resume \
  -H "Content-Type: application/json" \
  -d @examples/resume.json \
  -o test-cv.pdf

# Test Cover Letter
curl -X POST http://localhost:9000/cover-letter \
  -H "Content-Type: application/json" \
  -d @examples/cover-letter.json \
  -o test-letter.pdf
```
