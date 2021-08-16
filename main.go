package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"postgres"
	"structures"
	"xmlreader"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

//homePage запускает главную страницу и выполняет код расшифровки xml документа в json формат
func homePage(w http.ResponseWriter, r *http.Request) {
	//Читаю XML-файл
	data := xmlreader.UnmarshalXMLFile("xmlFiles//feed_example.xml")

	Buildings, IDProjects := postgres.GetDataFromBuilding()
	Sections, IDBuildings := postgres.GetDataFromSection()
	Lots, IDSections := postgres.GetDataFromLot()
	Projects := postgres.GetDataFromProject()

	if len(Projects) == 0 {
		Projects = append(Projects, structures.Project{ID: 0, Name: "Some name", Description: "Some description", Building: nil})
		postgres.AddDataToProject(0, "Some name", "Some description", "Some address")
	}

	//Начинаем добирать данные из файла XML и добавлять в BD недостающие
	var F bool
	postgres.AddDataToProject(len(Projects), "SomeName", "SomeDescription", "SomeAddress")

	for _, elem := range data.Project {
		F = true
		for _, elemB := range Buildings {
			if elemB.ID == elem.IDBuilding {
				F = false
				break
			}
		}
		if F {
			//Запись зданий в БД
			postgres.AddDataToBuilding(elem.IDBuilding, elem.NameBuilding, len(Projects)-1)
			Buildings = append(Buildings, structures.Building{ID: elem.IDBuilding, Name: elem.NameBuilding, Section: nil})
			IDProjects = append(IDProjects, len(Projects)-1)
		}

		F = true
		for _, elemS := range Sections {
			if elemS.ID == elem.IDSection {
				F = false
				break
			}
		}
		if F {
			//Добавляем секции в БД
			postgres.AddDataToSection(elem.IDSection, elem.Name, elem.IDBuilding)
			Sections = append(Sections, structures.Section{ID: elem.IDSection,
				Name: elem.Name, Lot: nil})
			IDBuildings = append(IDBuildings, elem.IDBuilding)
		}

		if !F {
			F = true
		}
		for _, elemL := range Lots {
			if elemL.ID == elem.IDLot {
				F = false
				break
			}
		}
		if F {
			//Запись лотов в БД
			postgres.AddDataToLot(elem.IDLot, elem.Floor, elem.TotalSquare,
				elem.LocalSquare, elem.KitchenSquare, elem.Price,
				elem.LotType, elem.RoomType, elem.IDSection)
			Lots = append(Lots, structures.Lot{ID: elem.IDLot, Floor: elem.Floor,
				TotalSquare: elem.TotalSquare, LocalSquare: elem.LocalSquare,
				KitchenSquare: elem.KitchenSquare, Price: elem.Price,
				LotType: elem.LotType, RoomType: elem.RoomType})
			IDSections = append(IDSections, elem.IDSection)
		}
	}

	Projects = xmlreader.MarshalToJson(Projects, Buildings, Sections, Lots, IDProjects, IDBuildings, IDSections)

	//Маршелим в json-формат
	jsonData, err := json.Marshal(Projects)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(jsonData))
}
