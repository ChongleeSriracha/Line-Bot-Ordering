package model

type RichMenu struct {
	Size        Size   `json:"size"`
	Selected    bool   `json:"selected"`
	Name        string `json:"name"`
	ChatBarText string `json:"chatBarText"`
	Areas       []Area `json:"areas"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Area struct {
	Bounds Bound  `json:"bounds"`
	Action Action `json:"action"`
}

type Bound struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Action struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
