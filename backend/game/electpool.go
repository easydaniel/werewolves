package game

type ElectRound struct {
	Name   string `json:"name"`
	Elect  []int  `json:"elect"`
	Weight []int  `json:"weight"`
	Result []int  `json:"result"`
	End    bool
}

type ElectPool struct {
	Pool []*ElectRound `json:"pool"`
}
