package author

type AuthorRequest struct {
	Name string `json:"name" binding:"required"`
}
