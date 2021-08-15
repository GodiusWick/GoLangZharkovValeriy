package PostgresPars

import (
	"database/sql"
	"fmt"
	"internal/DataStructures"

	_ "github.com/lib/pq"
)

var connectToDB bool = false
var db *sql.DB

//Вспомогательная функция, для установления соединения с БД на локальной машине
func makeConnectToDB() {
	PostSQLConnStr := "host=127.0.0.1 port=49161 user=root password=12345 dbname=ProjectDB sslmode=disable"

	DB, err := sql.Open("postgres", PostSQLConnStr)

	if err != nil {
		fmt.Println(err)
		return
	}

	db = DB

	connectToDB = true
}

//Взять данные из таблицы зданий
func GetDataFromBuilding() ([]DataStructures.Building, []int) {
	if !connectToDB {
		makeConnectToDB()
	}

	var IdProject []int
	var Arr []DataStructures.Building
	rows, _ := db.Query("SELECT * FROM building")
	defer rows.Close()

	for rows.Next() {
		elem := DataStructures.Building{}
		var Id int
		err := rows.Scan(&elem.Id, &elem.Name, &Id)

		if err != nil {
			panic(err)
		}

		Arr = append(Arr, elem)
		IdProject = append(IdProject, Id)
	}

	return Arr, IdProject
}

//Взть данные из таблицы секций
func GetDataFromSection() ([]DataStructures.Section, []int) {
	if !connectToDB {
		makeConnectToDB()
	}

	var IdBuilding []int
	var Arr []DataStructures.Section
	rows, _ := db.Query("SELECT * FROM section")

	for rows.Next() {
		elem := DataStructures.Section{}
		var Id int
		err := rows.Scan(&elem.Id, &elem.Name, &Id)

		if err != nil {
			panic(err)
		}

		Arr = append(Arr, elem)
		IdBuilding = append(IdBuilding, Id)
	}

	return Arr, IdBuilding
}

//Взять данные из таблицы Lot
func GetDataFromLot() ([]DataStructures.Lot, []int) {
	if !connectToDB {
		makeConnectToDB()
	}

	var IdSection []int
	var Arr []DataStructures.Lot
	rows, _ := db.Query("SELECT * FROM lot")

	for rows.Next() {
		elem := DataStructures.Lot{}
		var Id int
		err := rows.Scan(&elem.Id, &elem.Floor, &elem.TotalSquare, &elem.LocalSquare, &elem.KitchenSquare,
			&elem.Price, &elem.LotType, &elem.RoomType, &Id)

		if err != nil {
			panic(err)
		}

		Arr = append(Arr, elem)
		IdSection = append(IdSection, Id)
	}

	return Arr, IdSection
}

//Взять данные из таблицы Project
func GetDataFromProject() []DataStructures.Project {
	if !connectToDB {
		makeConnectToDB()
	}

	var Arr []DataStructures.Project
	rows, _ := db.Query("SELECT * FROM project")

	for rows.Next() {
		elem := DataStructures.Project{}
		err := rows.Scan(&elem.Id, &elem.Name, &elem.Description, &elem.Address)

		if err != nil {
			fmt.Println(err)
			continue
		}

		Arr = append(Arr, elem)
	}

	return Arr
}

//Добавление данных в таблицу зданий
func AddDataToBuilding(Id int, Name string, Id_Project int) {
	db.Exec("INSERT INTO building (id, name, id_project) VALUES($1, $2, $3)", Id, Name, Id_Project)
}

func AddDataToSection(Id int, Name string, Id_Building int) {
	db.Exec("INSERT INTO section (id, name, id_building) VALUES($1, $2, $3)", Id, Name, Id_Building)
}

//Добавление данных в таблицу Lot
func AddDataToLot(Id int, Floor int, TotalSquare float64, LocalSquare float64,
	KitchenSquare float64, Price int, LotType string, RoomType string, Id_Section int) {
	db.Exec("INSERT INTO lot (id, floor, total_square, living_square, kitchen_square, price, lot_type, room_type, id_section) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)", Id, Floor, TotalSquare, LocalSquare, KitchenSquare, Price, LotType, RoomType, Id_Section)
}

//Добавление данных в таблицу Project
func AddDataToProject(Id int, Name string, Description string, Address string) {
	db.Exec("INSERT INTO project (id, name, description, address) VALUES($1, $2, $3, $4)", Id, Name, Description, Address)
}
