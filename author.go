package mangarock

// Author contains the info related to a manga author
type Author struct {
	Name string `json:"author"`
	Oid  string `json:"oid"`
	Role string `json:"role"`
}
