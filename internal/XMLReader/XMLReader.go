package XMLReader

import (
	"encoding/xml"
	"fmt"
	"internal/DataStructures"
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

func MarshalToJson(Projects []DataStructures.Project, Buildings []DataStructures.Building,
	Sections []DataStructures.Section, Lots []DataStructures.Lot, IdProjects []int,
	IdBuildings []int, IdSections []int) []DataStructures.Project {

	for i, elem := range IdSections {
		for j, elemS := range Sections {
			if elemS.Id == elem {
				Sections[j].Lot = append(Sections[j].Lot, Lots[i])
			}
		}

	}

	for i, elem := range IdBuildings {
		for j, elemB := range Buildings {
			if elemB.Id == elem {
				Buildings[j].Section = append(Buildings[j].Section, Sections[i])
			}
		}
	}

	for i, elem := range IdProjects {
		for j, elemP := range Projects {
			if elemP.Id == elem {
				Projects[j].Building = append(Projects[j].Building, Buildings[i])
			}
		}
	}

	return Projects
}

// Вспомогательная структура для начала парсинга xml-документа
type DocumentData struct {
	Project []Offer `xml:"offer" json:"project"`
}

// Вспомогательная структура, которая собирает всю необходимую информацию из файла
type Offer struct {
	Name          string  `xml:"deal-status" json:"name"`
	Description   string  `xml:"description" json:"description"`
	Address       string  `xml:"location>address" json:"address"`
	IdBuilding    int     `xml:"yandex-building-id" json:"id-building"`
	NameBuilding  string  `xml:"building-name" json:"name-building"`
	Floor         int     `xml:"floor" json:"floor"`
	TotalSquare   float64 `xml:"area>value" json:"total_square"`
	LocalSquare   float64 `xml:"living-space>value" json:"living_square"`
	KitchenSquare float64 `xml:"kitchen-space>value" json:"kitchen_square"`
	Price         int     `xml:"price>value" json:"price"`
	LotType       string  `xml:"type" json:"lot_type"`
	RoomType      string  `xml:"category" json:"room_type"`
	IdSection     int     `xml:"building-section" json:"id-section"`
	IdLot         int     `xml:"internal-id,attr" json:"id"`
}
