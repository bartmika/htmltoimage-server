package dtos

type HTMLToImageRequestDTO struct {
	WebsiteURL string `json:"website_url"`
}
type HTMLToImageResponseDTO struct {
	FileName string `json:"file_name"`
	Content  []byte `json:"content"`
}
