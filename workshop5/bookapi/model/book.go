package model

// Book struct model
type Book struct {
	Isbn     string `json:"isbn"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
