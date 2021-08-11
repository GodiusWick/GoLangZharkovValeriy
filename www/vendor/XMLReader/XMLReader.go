package XMLReader

import (
	"DataStructures"
	"PostgresPars"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

func UnmarshalXMLFile(url string) DocumentData {
	openXmlFile, err := os.Open(url)
	if err != nil {
		fmt.Println("Ошибка в открытии файла:", err)
	}
	defer openXmlFile.Close()

	xmlData, err := ioutil.ReadAll(openXmlFile)
	if err != nil {
		fmt.Println("Ошибка в чтении файла:", err)
	}

	var data DocumentData
	xml.Unmarshal([]byte(xmlData), &data)

	return data
}

func MarshalToJson(Project []DataStructures.Project, Building []DataStructures.Building, Lot []DataStructures.Lot) []DataStructures.Project {
	for i, elem := range Project {
		for _, elemB := range Building {
			//Достраиваем структуру Project, привязывая здания
			if PostgresPars.CheckStringProjectBuilding(true, true, elem.Id, elemB.Id) {
				Project[i].Building = append(Project[i].Building, elemB)
			}
		}
	}

	for i, elem := range Project {
		for j, elemPB := range Project[i].Building {
			//Достариваем структуры Building в Project, привязывая лоты
			if PostgresPars.CheckStringLotBuilding(true, true, elem.Id, elemPB.Id) {
				for _, elemL := range Lot {
					if elemL.Id == elem.Id {
						Project[i].Building[j].Lot = append(Project[i].Building[j].Lot, elemL)
						break
					}
				}

			}
		}
	}

	return Project
}

// Вспомогательная структура для начала парсинга xml-документа
type DocumentData struct {
	Project []Offer `xml:"offer" json:"project"`
}

// Вспомогательная структура, которая собирает всю необходимую информацию из файла
type Offer struct {
	Id            int     `xml:"internal-id,attr" json:"id"`
	Name          string  `xml:"deal-status" json:"name"`
	Description   string  `xml:"description" json:"description"`
	Address       string  `xml:"location>address" json:"address"`
	IdBuilding    int     `xml:"yandex-building-id" json:"id-building"`
	NameBuilding  string  `xml:"building-name" json:"name-building"`
	Floor         int     `xml:"floor" json:"floor"`
	TotalSquare   float64 `xml:"area>value" json:"total_square"`
	LocalSquare   float64 `xml:"living-space>value" json:"living_square"`
	KitchenSquare float64 `xml:"kitchen-space>value" json:"kitchen_square"`
	Price         float64 `xml:"price>value" json:"price"`
	LotType       string  `xml:"type" json:"lot_type"`
	RoomType      string  `xml:"category" json:"room_type"`
}
