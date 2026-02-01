# candipack-pdf

Open source PDF generation engine for resumes and cover letters using JSON Resume schema convention.

## Features

- Generate professional PDF resumes from JSON data
- Generate cover letters in PDF format
- Multiple customizable HTML/CSS templates
- Simple REST API
- Multi-language support (en, fr)
- Follows JSON Resume schema convention

## Quick Start

### Using Docker Compose (Recommended)

```bash
# Clone
git clone https://github.com/snowztech/candipack-pdf.git
cd candipack-pdf

# Create environment file (optional)
cp .env.example .env

# Launch with Docker Compose
docker-compose up

# Test health endpoint
curl http://localhost:9000/health
```

### Local Development

```bash
# Clone
git clone https://github.com/snowztech/candipack-pdf.git
cd candipack-pdf

# Install dependencies
go mod download

# Run
go run ./cmd/server

# Test health endpoint
curl http://localhost:9000/health
```

## API Endpoints

| Method | Endpoint             | Description                |
| ------ | -------------------- | -------------------------- |
| `POST` | `/resume`            | Generate PDF resume        |
| `POST` | `/resume/html`       | Generate HTML resume       |
| `POST` | `/cover-letter`      | Generate PDF cover letter  |
| `POST` | `/cover-letter/html` | Generate HTML cover letter |
| `GET`  | `/templates`         | List available templates   |
| `GET`  | `/health`            | Health check               |

### Headers

```
Content-Type: application/json
X-API-Key: your-api-key (optional, if configured)
```

## Available Templates

### CV/Resume

- `professional` - 2-column layout with sidebar (default)
- `simple` - 1-column minimalist layout

### Cover Letter

- `classic` - Traditional letter format

## Example Usage

### Generate Resume

```bash
curl -X POST http://localhost:9000/resume \
  -H "Content-Type: application/json" \
  -d @examples/resume.json \
  -o resume.pdf
```

### Generate Cover Letter

```bash
curl -X POST http://localhost:9000/cover-letter \
  -H "Content-Type: application/json" \
  -d @examples/cover-letter.json \
  -o cover-letter.pdf
```

### List Templates

```bash
curl http://localhost:9000/templates
```

## JSON Schema

The API accepts JSON following the [JSON Resume](https://jsonresume.org/) schema convention with some extensions.

### Resume Request Example

See [example](examples/resume.json)

```json
{
  "meta": {
    "template": "professional",
    "lang": "en"
  },
  "basics": {
    "name": "Jane Smith",
    "label": "Software Engineer",
    "email": "jane@example.com",
    "phone": "+1 555-123-4567",
    "summary": "Experienced software engineer...",
    "location": {
      "city": "San Francisco",
      "countryCode": "US"
    }
  },
  "work": [...],
  "education": [...],
  "skills": [...],
  "languages": [...]
}
```

### Cover Letter Request Example

See [example](examples/cover-letter.json)

```json
{
  "meta": {
    "template": "classic",
    "lang": "en"
  },
  "sender": {
    "name": "Jane Smith",
    "address": "123 Main St, San Francisco, CA",
    "phone": "+1 555-123-4567",
    "email": "jane@example.com"
  },
  "recipient": {
    "company": "Tech Corp",
    "address": "456 Market St, San Francisco, CA"
  },
  "date": "January 31, 2026",
  "subject": "Application for Software Engineer Position",
  "salutation": "Dear Hiring Manager,",
  "paragraphs": ["Paragraph 1...", "Paragraph 2..."],
  "closing": "Sincerely,",
  "signature": "Jane Smith"
}
```

## Environment Variables

| Variable  | Default | Description               |
| --------- | ------- | ------------------------- |
| `PORT`    | `9000`  | Server port               |
| `API_KEY` | -       | Optional API key for auth |

## Docker

### Docker Compose (Recommended)

```bash
# Build and run
docker-compose up

# Run in background
docker-compose up -d

# Stop
docker-compose down

# Rebuild after changes
docker-compose up --build
```

### Docker (Manual)

```bash
# Build
docker build -t candipack-pdf .

# Run
docker run -p 9000:9000 candipack-pdf

# Run with custom env
docker run -p 9000:9000 -e API_KEY=secret candipack-pdf
```

## Project Structure

```
candipack-pdf/
├── cmd/server/          # Entry point
├── internal/
│   ├── handlers/        # HTTP handlers
│   ├── generator/       # PDF generation (chromedp)
│   ├── middleware/      # Auth & CORS
│   ├── models/          # Data structures
│   ├── parser/          # HTML template parser
│   └── lang/            # i18n support
├── templates/           # HTML templates
│   ├── cv/
│   └── cover-letter/
├── examples/           # Sample JSON files
├── configs/            # Configuration
└── README.md
```

## Tech Stack

- **Go** - Backend language
- **Gin** - HTTP web framework
- **ChromeDP** - Headless Chrome for PDF generation
- **Go Templates** - HTML templating

## License

MIT License - See [license](LICENSE) file for details
