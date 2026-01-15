package articles;

type Article struct {
	Title string `json:"title"`
	Tags []string `json:"tags,omitempty"`
	EditorComment string `json:"-"`
}