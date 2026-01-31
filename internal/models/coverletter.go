package models

type CoverLetter struct {
	Meta       CoverLetterMeta `json:"meta"`
	Sender     Sender          `json:"sender"`
	Recipient  Recipient       `json:"recipient"`
	Date       string          `json:"date"`
	Subject    string          `json:"subject"`
	Salutation string          `json:"salutation"`
	Paragraphs []string        `json:"paragraphs"`
	Closing    string          `json:"closing"`
	Signature  string          `json:"signature"`
}

type CoverLetterMeta struct {
	Template string `json:"template"`
	Lang     string `json:"lang"`
}

type Sender struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type Recipient struct {
	Company string `json:"company"`
	Address string `json:"address"`
}
