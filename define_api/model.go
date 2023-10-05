package define_api

type APICounter struct {
	OptionsAPI         int     `json:"OptionsAPI"`
	CompositionAPI     int     `json:"CompositionAPI"`
	OptionsPercent     float64 `json:"OptionsPercent"`
	CompositionPercent float64 `json:"CompositionPercent"`
	TotalFiles         int     `json:"TotalFiles"`
}
