package DataStructures

type Project struct {
	Id          int        `xml:"id" json:"id"`
	Name        string     `xml:"building-name" json:"name"`
	Description string     `xml:"description" json:"description"`
	Address     string     `xml:"location>address" json:"address"`
	Building    []Building `xml:"building" json:"building"`
}

// Структура по заданию
type Building struct {
	Id      int       `xml:"yandex-building-id" json:"id"`
	Name    string    `xml:"building-name" json:"name"`
	Section []Section `xml:"section" json:"section"`
}

type Section struct {
	Id   int    `xml:"building-section" json:"id"`
	Name string `xml:"building-name" json:"name"`
	Lot  []Lot  `xml:"lot" json:"lot"`
}

// Структура по заданию
type Lot struct {
	Id            int     `xml:"id" json:"id"`
	Floor         int     `xml:"floor" json:"floor"`
	TotalSquare   float64 `xml:"area>value" json:"total_square"`
	LocalSquare   float64 `xml:"living-space>value" json:"living_square"`
	KitchenSquare float64 `xml:"kitchen-space>value" json:"kitchen_square"`
	Price         int     `xml:"price>value" json:"price"`
	LotType       string  `xml:"type" json:"lot_type"`
	RoomType      string  `xml:"category" json:"room_type"`
}
