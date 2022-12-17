package models

type Post struct {
	Title   string
	Message string
}

func (p Post) GetTitle() string {
	return p.Title
}

func (p *Post) SetTitle(title string) {
	p.Title = title
}

func (p Post) GetMessage() string {
	return p.Message
}

func (p *Post) SetMessage(message string) {
	p.Message = message
}
