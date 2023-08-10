package model

type SwitchData struct {
	Name        string
	Type        string
	Actuation   string
	PreTravel   string
	TotalTravel string
}

type Switch struct {
	SwitchData
	MarketPrice string
	ImgUrl      string
	ProductUrl  string
}
