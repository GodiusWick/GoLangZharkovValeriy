package main

import (
	"encoding/json"
	"fmt"

	"DataStructures"
	"PostgresPars"
	"XMLReader"

	_ "github.com/lib/pq"
)

func main() {

	//Читаю XML-файл
	data := XMLReader.UnmarshalXMLFile("xmlFiles//feed_example.xml")

	Buildings := PostgresPars.GetDataFromBuilding()
	Lots := PostgresPars.GetDataFromLot()
	Projects := PostgresPars.GetDataFromProject()

	//Начинаем добирать данные из файла XML и добавлять в BD недостающие
	var F bool

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
			PostgresPars.AddDataToBuilding(elem.IdBuilding, elem.NameBuilding)
			Buildings = append(Buildings, DataStructures.Building{Id: elem.IdBuilding, Name: elem.NameBuilding, Lot: nil})
		}
		if !F {
			F = true
		}
		for _, elemL := range Lots {
			if elemL.Id == elem.Id {
				F = false
				break
			}
		}
		if F {
			//Запись лотов в БД
			PostgresPars.AddDataToLot(elem.Id, elem.Floor, elem.TotalSquare,
				elem.LocalSquare, elem.KitchenSquare, elem.Price,
				elem.LotType, elem.RoomType)
			Lots = append(Lots, DataStructures.Lot{Id: elem.Id, Floor: elem.Floor,
				TotalSquare: elem.TotalSquare, LocalSquare: elem.LocalSquare,
				KitchenSquare: elem.KitchenSquare, Price: elem.Price,
				LotType: elem.LotType, RoomType: elem.RoomType})
		}
		//Одновременно дополняем таблицу связывающие здания и лоты
		PostgresPars.AddDataToLotBuilding(elem.Id, elem.IdBuilding)

		F = true
		for _, elemP := range Projects {
			if elemP.Id == elem.Id {
				F = false
				break
			}
		}
		if F {
			//Добавляем заказы в БД
			PostgresPars.AddDataToProject(elem.Id, elem.Name, elem.Description,
				elem.Address)
			Projects = append(Projects, DataStructures.Project{Id: elem.Id,
				Name: elem.Name, Description: elem.Description,
				Address: elem.Address, Building: nil})
		}
		//Одновременно дополняем данные в таблицу, связывающую заказы и здания
		PostgresPars.AddDataToProjectBuilding(elem.Id, elem.IdBuilding)
	}

	Projects = XMLReader.MarshalToJson(Projects, Buildings, Lots)

	//Маршелим в json-формат
	jsonData, _ := json.Marshal(Projects)
	fmt.Println(string(jsonData))

}
