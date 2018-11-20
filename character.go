package mangarock

// Character contains the information related to a Manga character
type Character struct {
	Name      string `json:"name"`
	Oid       string `json:"oid"`
	Thumbnail string `json:"thumbnail"`
}
