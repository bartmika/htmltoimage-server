package dtos

type ScreenshotRequestDTO struct {
	WebsiteURL   string `json:"website_url"`
	ImageType    string `json:"image_type"`
	ImageQuality int    `json:"image_quality"`
}
type ScreenshotResponseDTO struct {
	FileName string `json:"file_name"`
	Content  []byte `json:"content"`
}
