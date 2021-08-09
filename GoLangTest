package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	//Беру файл из папки xmlFiles
	xmlFilePath := "xmlFiles//feed_example.xml"
	openXmlFile, err := os.Open(xmlFilePath)

	if err != nil {
		fmt.Println("Ошибка при открытии файла (скорее всего вы неправильно указали путь):", err)
		return
	}
	defer openXmlFile.Close()

	//Устанановка соединения с базой данных
	PostSQLConnStr := "host=localhost port=5432 user=Valeriy password=10IMclass dbname=TestGo sslmode=disable"

	db, err := sql.Open("postgres", PostSQLConnStr)

	if err != nil {
		fmt.Println(err)
		return
	}

	//Помещаем содержимое xml-файла в перменную
	XMLdata, _ := ioutil.ReadAll(openXmlFile)

	//Парсим xmlфайл
	var data DocumentData
	xml.Unmarshal([]byte(XMLdata), &data)

	rowsBuild, _ := db.Query("SELECT * FROM building ")
	//Объект, в который записываем необходимые данные по зданиям из базы данных
	Buildings := []Building{}

	for rowsBuild.Next() {
		elem := Building{}
		err := rowsBuild.Scan(&elem.Id, &elem.Name)

		if err != nil {
			fmt.Println(err)
			continue
		}

		Buildings = append(Buildings, elem)
	}

	//Объект, в который записываем необходимые данные по лотам из базы данных
	Lots := []Lot{}

	rowsLot, _ := db.Query("SELECT * FROM lot")

	for rowsLot.Next() {
		elem := Lot{}
		err := rowsLot.Scan(&elem.Id, &elem.Floor, &elem.TotalSquare, &elem.LocalSquare,
			&elem.KitchenSquare, &elem.Price, &elem.LotType, &elem.RoomType)

		if err != nil {
			fmt.Println(err)
			continue
		}

		Lots = append(Lots, elem)
	}

	//Объект, в который записываем необходимые данные по заказам(как понял) из базы данных
	Projects := []Project{}

	rowsProject, _ := db.Query("SELECT * FROM project")

	for rowsProject.Next() {
		elem := Project{}
		err := rowsProject.Scan(&elem.Id, &elem.Name, &elem.Description, &elem.Address)

		if err != nil {
			fmt.Println(err)
			continue
		}

		Projects = append(Projects, elem)
	}

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
			db.Exec("INSERT INTO building (id, name) VALUES($1, $2)", elem.IdBuilding, elem.NameBuilding)
			Buildings = append(Buildings, Building{elem.IdBuilding, elem.NameBuilding, nil})
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
			db.Exec("INSERT INTO lot (id,floor, total_square, local_square,"+
				"kitchen_square, price, lot_type, room_type)"+
				"VALUES($1, $2, $3, $4, $5, $6, $7, $8)", elem.Id, elem.Floor, elem.TotalSquare,
				elem.LocalSquare, elem.KitchenSquare, elem.Price,
				elem.LotType, elem.RoomType)
			Lots = append(Lots, Lot{elem.Id, elem.Floor, elem.TotalSquare,
				elem.LocalSquare, elem.KitchenSquare, elem.Price,
				elem.LotType, elem.RoomType})
		}
		//Одновременно дополняем таблицу связывающие здания и лоты
		row, _ := db.Query("SELECT * FROM lotbuilding WHERE idlot = $1 LIMIT 1", elem.Id)
		if !row.Next() {
			db.Exec("INSERT INTO lotbuilding (idlot, idbuilding) VALUES($1, $2)", elem.Id, elem.IdBuilding)
		}
		defer row.Close()

		F = true
		for _, elemP := range Projects {
			if elemP.Id == elem.Id {
				F = false
				break
			}
		}
		if F {
			//Добавляем заказы в БД
			db.Exec("INSERT INTO project (id,name, description, address)"+
				"VALUES($1, $2, $3, $4)", elem.Id, elem.Name, elem.Description,
				elem.Address)
			Projects = append(Projects, Project{elem.Id, elem.Name, elem.Description,
				elem.Address, nil})
		}
		//Одновременно дополняем данные в таблицу, связывающую заказы и здания
		rowP, _ := db.Query("SELECT * FROM projectbuilding WHERE idproject = $1 LIMIT 1", elem.Id)

		if !rowP.Next() {
			db.Exec("INSERT INTO projectbuilding (idproject, idbuilding) VALUES($1, $2)", elem.Id, elem.IdBuilding)
		}
		defer rowP.Close()
	}

	for i, elem := range Projects {
		for _, elemB := range Buildings {
			//Достраиваем структуру Project, привязывая здания
			rowS, _ := db.Query("SELECT * FROM projectbuilding WHERE idproject=$1 AND idbuilding=$2", elem.Id, elemB.Id)
			if rowS.Next() {
				Projects[i].Building = append(Projects[i].Building, elemB)
			}
			defer rowS.Close()
		}
	}

	for i, elem := range Projects {
		for j, elemPB := range Projects[i].Building {
			//Достариваем структуры Building в Project, привязывая лоты
			rowSS, _ := db.Query("SELECT * FROM lotbuilding WHERE idlot=$1 AND idbuilding=$2", elem.Id, elemPB.Id)
			if rowSS.Next() {
				for _, elemL := range Lots {
					if elemL.Id == elem.Id {
						Projects[i].Building[j].Lot = append(Projects[i].Building[j].Lot, elemL)
						break
					}
				}

			}
			defer rowSS.Close()
		}
	}

	//Маршелим в json-формат
	jsonData, _ := json.Marshal(Projects)
	fmt.Println(string(jsonData))

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

// Структура по заданию
type Project struct {
	Id          int        `xml:"internal-id,attr" json:"id"`
	Name        string     `xml:"deal-status" json:"name"`
	Description string     `xml:"description" json:"description"`
	Address     string     `xml:"location>address" json:"address"`
	Building    []Building `xml:"building" json:"building"`
}

// Структура по заданию
type Building struct {
	Id   int    `xml:"yandex-building-id" json:"id"`
	Name string `xml:"building-name" json:"name"`
	Lot  []Lot  `xml:"lot" json:"lot"`
}

// Структура по заданию
type Lot struct {
	Id            int     `xml:"new-flat" json:"id"`
	Floor         int     `xml:"floor" json:"floor"`
	TotalSquare   float64 `xml:"area>value" json:"total_square"`
	LocalSquare   float64 `xml:"living-space>value" json:"living_square"`
	KitchenSquare float64 `xml:"kitchen-space>value" json:"kitchen_square"`
	Price         float64 `xml:"price>value" json:"price"`
	LotType       string  `xml:"type" json:"lot_type"`
	RoomType      string  `xml:"category" json:"room_type"`
}
