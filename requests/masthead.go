package requests

type CreateMastheadInput struct {
	ImageURL string `json:"image_url" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Order    int    `json:"order"`
	Status   int    `json:"status"`
}

type UpdateMastheadInput struct {
	ImageURL *string `json:"image_url,omitempty"`
	Link     *string `json:"link,omitempty"`
	Order    *int    `json:"order,omitempty"`
	Status   *int    `json:"status,omitempty"`
}
