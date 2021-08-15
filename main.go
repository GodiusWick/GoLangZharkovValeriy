package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"internal/DataStructures"
	"internal/PostgresPars"
	"internal/XMLReader"

	_ "github.com/lib/pq"
)

func main() {

	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)

}

func home_page(w http.ResponseWriter, r *http.Request) {
	//Читаю XML-файл
	data := XMLReader.UnmarshalXMLFile("xmlFiles//feed_example.xml")

	Buildings, IdProjects := PostgresPars.GetDataFromBuilding()
	Sections, IdBuildings := PostgresPars.GetDataFromSection()
	Lots, IdSections := PostgresPars.GetDataFromLot()
	Projects := PostgresPars.GetDataFromProject()

	if len(Projects) == 0 {
		Projects = append(Projects, DataStructures.Project{Id: 0, Name: "Some name", Description: "Some description", Building: nil})
		PostgresPars.AddDataToProject(0, "Some name", "Some description", "Some address")
	}

	//Начинаем добирать данные из файла XML и добавлять в BD недостающие
	var F bool
	PostgresPars.AddDataToProject(len(Projects), "SomeName", "SomeDescription", "SomeAddress")

	for _, elem := range data.Project {
		F = true
		for _, elemB := range Buildings {
			if elemB.Id == elem.IdBuilding {
				F = false
				break
			}
		}
		if F {
			//Запись зданий в БД
			PostgresPars.AddDataToBuilding(elem.IdBuilding, elem.NameBuilding, len(Projects)-1)
			Buildings = append(Buildings, DataStructures.Building{Id: elem.IdBuilding, Name: elem.NameBuilding, Section: nil})
			IdProjects = append(IdProjects, len(Projects)-1)
		}

		F = true
		for _, elemS := range Sections {
			if elemS.Id == elem.IdSection {
				F = false
				break
			}
		}
		if F {
			//Добавляем секции в БД
			PostgresPars.AddDataToSection(elem.IdSection, elem.Name, elem.IdBuilding)
			Sections = append(Sections, DataStructures.Section{Id: elem.IdSection,
				Name: elem.Name, Lot: nil})
			IdBuildings = append(IdBuildings, elem.IdBuilding)
		}

		if !F {
			F = true
		}
		for _, elemL := range Lots {
			if elemL.Id == elem.IdLot {
				F = false
				break
			}
		}
		if F {
			//Запись лотов в БД
			PostgresPars.AddDataToLot(elem.IdLot, elem.Floor, elem.TotalSquare,
				elem.LocalSquare, elem.KitchenSquare, elem.Price,
				elem.LotType, elem.RoomType, elem.IdSection)
			Lots = append(Lots, DataStructures.Lot{Id: elem.IdLot, Floor: elem.Floor,
				TotalSquare: elem.TotalSquare, LocalSquare: elem.LocalSquare,
				KitchenSquare: elem.KitchenSquare, Price: elem.Price,
				LotType: elem.LotType, RoomType: elem.RoomType})
			IdSections = append(IdSections, elem.IdSection)
		}
	}

	Projects = XMLReader.MarshalToJson(Projects, Buildings, Sections, Lots, IdProjects, IdBuildings, IdSections)

	//Маршелим в json-формат
	jsonData, err := json.Marshal(Projects)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(jsonData))
}
