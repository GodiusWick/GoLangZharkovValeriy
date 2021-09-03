package main

import (
	"os"
	"structures"
	"testing"

	_ "github.com/lib/pq"
)

// Тесты для файла xmlreader

func TestUnmarshalXMLFile(t *testing.T) {
	// Набор данных
	_, err := os.Open("xmlfiles//feed_example.xml")

	// Действия

	// Сверка результата

	if err != nil {
		t.Error(err)
	}
}

func TestMarshalToJson(t *testing.T) {
	// Набор данных

	var Projects []structures.Project
	var Buildings []structures.Building
	var Sections []structures.Section
	var Lots []structures.Lot
	var IDBuildings []int
	var IDSections []int
	var IDProjects []int

	Projects = append(Projects, structures.Project{})
	Buildings = append(Buildings, structures.Building{})
	Sections = append(Sections, structures.Section{})
	Lots = append(Lots, structures.Lot{})
	IDBuildings = append(IDBuildings, 0)
	IDSections = append(IDSections, 0)
	IDProjects = append(IDProjects, 0)

	// Действия

	for i, elemS := range Sections {
		for j, elem := range IDSections {
			if elemS.ID == elem {
				Sections[i].Lot = append(Sections[i].Lot, Lots[j])
			}
		}

	}

	for i, elemB := range Buildings {
		for j, elem := range IDBuildings {
			if elemB.ID == elem {
				Buildings[i].Section = append(Buildings[i].Section, Sections[j])
			}
		}
	}

	for i, elemP := range Projects {
		for j, elem := range IDProjects {
			if elemP.ID == elem {
				Projects[i].Building = append(Projects[i].Building, Buildings[j])
			}
		}
	}

	// Итоговый результат

	if Projects[0].Building[0].Section[0].Lot == nil {
		t.Error("Неправильное заполнение!")
	}
}

// Тесты для файла MainStruct не прописываются, т.к. этот файл содержит только главную структуру,
// и неисправности могут выявить линтеры!

// Тесты для проверки коннекта к БД виснут при неправильном IP, но при этом error = nil, что???
