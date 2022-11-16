package requests

import "mime/multipart"

type CreateMastheadInput struct {
	ImageURL *multipart.FileHeader `form:"image_url" binding:"required"`
	Link     string                `form:"link" binding:"required"`
	Order    int                   `form:"order"`
	Status   int                   `form:"status"`
}

type UpdateMastheadInput struct {
	ImageURL *string `json:"image_url,omitempty"`
	Link     *string `json:"link,omitempty"`
	Order    *int    `json:"order,omitempty"`
	Status   *int    `json:"status,omitempty"`
}
