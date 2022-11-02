package requests

type CreateMastheadInput struct {
	ImageURL string `json:"image_url" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Order    int    `json:"order"`
	Status   int    `json:"status"`
}

type UpdateMastheadInput struct {
	ImageURL string `json:"image_url"`
	Link     string `json:"link"`
	Order    int    `json:"order"`
	Status   int    `json:"status"`
}
