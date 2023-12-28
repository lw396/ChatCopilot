package repository

type Classify struct {
	Id       uint64 `json:"id"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
	Describe string `json:"describe"`
	Types    uint8  `json:"types"`
	Basic    string `json:"basic"`
	Weigh    uint64 `json:"weigh"`
	Top      uint8  `json:"top"`
	IsDel    uint8  `json:"is_del"`
}

func (Classify) TableName() string {
	return "pw_classify"
}

type SetMeal struct {
	Id       uint64  `json:"id"`
	ClId     uint64  `json:"cl_id"`
	Type     uint8   `json:"type"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Describe string  `json:"describe"`
	Num      uint64  `json:"num"`
	Weigh    uint64  `json:"weigh"`
	Top      uint8   `json:"top"`
	IsDel    uint8   `json:"is_del"`
}

func (SetMeal) TableName() string {
	return "pw_set_meal"
}
