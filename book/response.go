package book

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	AuthorID    int    `json:"author_id"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
}

type BookResponseView struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	AuthorID    int    `json:"author_id"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	AuthorName  string `json:"author_name"`
}

type AuthorResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
