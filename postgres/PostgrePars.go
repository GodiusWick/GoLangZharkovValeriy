package postgres

import (
	"database/sql"
	"fmt"
	"structures"
)

var connectToDB bool = false
var db *sql.DB

//makeConnectToDB Вспомогательная функция, для установления соединения с БД на локальной машине
func makeConnectToDB() {
	PostSQLConnStr := "host=host.docker.internal port=49156 user=root password=12345 dbname=ProjectDB sslmode=disable"

	DB, err := sql.Open("postgres", PostSQLConnStr)

	if err != nil {
		fmt.Println(err)
		return
	}

	db = DB

	connectToDB = true
}

//GetDataFromBuilding достаёт данные из таблицы Building
func GetDataFromBuilding() ([]structures.Building, []int) {
	if !connectToDB {
		makeConnectToDB()
	}

	var IDProject []int
	var Arr []structures.Building
	rows, _ := db.Query("SELECT * FROM building")
	defer rows.Close()

	for rows.Next() {
		elem := structures.Building{}
		var ID int
		err := rows.Scan(&elem.ID, &elem.Name, &ID)

		if err != nil {
			panic(err)
		}

		Arr = append(Arr, elem)
		IDProject = append(IDProject, ID)
	}

	return Arr, IDProject
}

//GetDataFromSection Взть данные из таблицы секций
func GetDataFromSection() ([]structures.Section, []int) {
	if !connectToDB {
		makeConnectToDB()
	}

	var IDBuilding []int
	var Arr []structures.Section
	rows, _ := db.Query("SELECT * FROM section")

	for rows.Next() {
		elem := structures.Section{}
		var ID int
		err := rows.Scan(&elem.ID, &elem.Name, &ID)

		if err != nil {
			panic(err)
		}

		Arr = append(Arr, elem)
		IDBuilding = append(IDBuilding, ID)
	}

	return Arr, IDBuilding
}

//GetDataFromLot Взять данные из таблицы Lot
func GetDataFromLot() ([]structures.Lot, []int) {
	if !connectToDB {
		makeConnectToDB()
	}

	var IDSection []int
	var Arr []structures.Lot
	rows, _ := db.Query("SELECT * FROM lot")

	for rows.Next() {
		elem := structures.Lot{}
		var ID int
		err := rows.Scan(&elem.ID, &elem.Floor, &elem.TotalSquare, &elem.LocalSquare, &elem.KitchenSquare,
			&elem.Price, &elem.LotType, &elem.RoomType, &ID)

		if err != nil {
			panic(err)
		}

		Arr = append(Arr, elem)
		IDSection = append(IDSection, ID)
	}

	return Arr, IDSection
}

//Взять данные из таблицы Project
func GetDataFromProject() []structures.Project {
	if !connectToDB {
		makeConnectToDB()
	}

	var Arr []structures.Project
	rows, _ := db.Query("SELECT * FROM project")

	for rows.Next() {
		elem := structures.Project{}
		err := rows.Scan(&elem.ID, &elem.Name, &elem.Description, &elem.Address)

		if err != nil {
			fmt.Println(err)
			continue
		}

		Arr = append(Arr, elem)
	}

	return Arr
}

//AddDataToBuilding Добавление данных в таблицу зданий
func AddDataToBuilding(ID int, Name string, IDProject int) {
	db.Exec("INSERT INTO building (id, name, id_project) VALUES($1, $2, $3)", ID, Name, IDProject)
}

func AddDataToSection(ID int, Name string, IDBuilding int) {
	db.Exec("INSERT INTO section (id, name, id_building) VALUES($1, $2, $3)", ID, Name, IDBuilding)
}

//Добавление данных в таблицу Lot
func AddDataToLot(ID int, Floor int, TotalSquare float64, LocalSquare float64,
	KitchenSquare float64, Price int, LotType string, RoomType string, IDSection int) {
	db.Exec("INSERT INTO lot (id, floor, total_square, living_square, kitchen_square, price, lot_type, room_type, id_section) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)", ID, Floor, TotalSquare, LocalSquare, KitchenSquare, Price, LotType, RoomType, IDSection)
}

//Добавление данных в таблицу Project
func AddDataToProject(ID int, Name string, Description string, Address string) {
	db.Exec("INSERT INTO project (id, name, description, address) VALUES($1, $2, $3, $4)", ID, Name, Description, Address)
}
