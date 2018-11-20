package mangarock

// Chapter contains the information related to a Manga chapter
type Chapter struct {
	Cid       int    `json:"cid"`
	Name      string `json:"name"`
	Oid       string `json:"oid"`
	Order     int    `json:"order"`
	UpdatedAt int    `json:"updated_at"`
}
