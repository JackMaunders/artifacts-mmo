package schema

type Destination struct {
	Name    string  `json:"name"`
	Skin    string  `json:"skin"`
	X       int     `json:"x"`
	Y       int     `json:"y"`
	Content Content `json:"content"`
}

type Content struct {
	Code string `json:"code"`
	Type string `json:"type"`
}
