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

	for i, elemS := range Sections {
		for j, elem := range IdSections {
			if elemS.Id == elem {
				Sections[i].Lot = append(Sections[i].Lot, Lots[j])
			}
		}

	}

	for i, elemB := range Buildings {
		for j, elem := range IdBuildings {
			if elemB.Id == elem {
				Buildings[i].Section = append(Buildings[i].Section, Sections[j])
			}
		}
	}

	for i, elemP := range Projects {
		for j, elem := range IdProjects {
			if elemP.Id == elem {
				Projects[i].Building = append(Projects[i].Building, Buildings[j])
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
