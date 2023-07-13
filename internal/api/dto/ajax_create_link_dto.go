package dto

type AjaxCreateLinkDto struct {
	Link string `json:"link" validate:"required,url"`
	// TODO: rest of the fields
}
