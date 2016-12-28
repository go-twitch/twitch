package twitch

type TopGamesResult struct {
	Total int       `json:"_total"`
	Top   []TopGame `json:"top"`
}

type TopGame struct {
	Channels int  `json:"channels"`
	Viewers  int  `json:"viewers"`
	Game     Game `json:"game"`
}

type Game struct {
	ID  uint64 `json:"_id"`
	Box struct {
		Large    string `json:"large"`
		Medium   string `json:"medium"`
		Small    string `json:"small"`
		Template string `json:"template"`
	} `json:"box"`
	GiantbombID uint64 `json:"giantbomb_id"`
	Logo        struct {
		Large    string `json:"large"`
		Medium   string `json:"medium"`
		Small    string `json:"small"`
		Template string `json:"template"`
	} `json:"logo"`
	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
}
