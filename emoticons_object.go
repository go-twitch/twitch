package twitch

type EmoticonSetsResponse struct {
	EmoticonSets EmoticonSets `json:"emoticon_sets"`
}

type EmoticonSets map[string][]Emoticon

type Emoticon struct {
	Code string `json:"code"`
	ID   int64  `json:"id"`
}
