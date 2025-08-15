package json

type Phrase struct {
	Phrase      string `json:"phrase"`
	Translation string `json:"translation"`
}

type Definition struct {
	Translation string `json:"translation"`
	Type        string `json:"type"`
}

type VocItem struct {
	Word         string       `json:"word"`
	Translations []Definition `json:"translations"`
	Phrases      []Phrase     `json:"phrases"`
}

//easyjson:json
type VocItemList []VocItem

// 实现 sort.Interface 接口
func (v VocItemList) Len() int           { return len(v) }
func (v VocItemList) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v VocItemList) Less(i, j int) bool { return v[i].Word < v[j].Word }
