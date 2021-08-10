package PostgresPars

import (
	"DataStructures"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var connectToDB bool = false
var db *sql.DB

func makeConnectToDB() {
	PostSQLConnStr := "host=localhost port=5432 user=Valeriy password=10IMclass dbname=TestGo sslmode=disable"

	DB, err := sql.Open("postgres", PostSQLConnStr)

	if err != nil {
		fmt.Println(err)
		return
	}

	db = DB

	connectToDB = true
}

func GetDataFromBuilding() []DataStructures.Building {
	if !connectToDB {
		makeConnectToDB()
	}

	var Arr []DataStructures.Building
	rows, _ := db.Query("SELECT * FROM building")

	for rows.Next() {
		elem := DataStructures.Building{}
		err := rows.Scan(&elem.Id, &elem.Name)

		if err != nil {
			fmt.Println(err)
			continue
		}

		Arr = append(Arr, elem)
	}

	return Arr
}

func GetDataFromLot() []DataStructures.Lot {
	if !connectToDB {
		makeConnectToDB()
	}

	var Arr []DataStructures.Lot
	rows, _ := db.Query("SELECT * FROM lot")

	for rows.Next() {
		elem := DataStructures.Lot{}
		err := rows.Scan(&elem.Id, &elem.Floor, &elem.TotalSquare, &elem.LocalSquare,
			&elem.Price, &elem.KitchenSquare, &elem.LotType, &elem.RoomType)

		if err != nil {
			fmt.Println(err)
			continue
		}

		Arr = append(Arr, elem)
	}

	return Arr
}

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

func AddDataToBuilding(Id int, Name string) {
	db.Exec("INSERT INTO building (id, name) VALUES($1, $2)", Id, Name)
}

func AddDataToLot(Id int, Floor int, TotalSquare float64, LocalSquare float64,
	KitchenSquare float64, Price float64, LotType string, RoomType string) {
	db.Exec("INSERT INTO building (id, floor, total_square, local_square, "+
		"kitchen_square, price, lot_type, room_type)"+
		" VALUES($1, $2, $3, $4, $5, $6, &7, &8)",
		Id, Floor, TotalSquare, LocalSquare, KitchenSquare, Price, LotType, RoomType)
}

func AddDataToProject(Id int, Name string, Description string, Address string) {
	db.Exec("INSERT INTO building (id, name, description, address)"+
		" VALUES($1, $2, $3, $4)", Id, Name, Description, Address)
}

func AddDataToLotBuilding(IdLot int, IdBuilding int) {
	row, err := db.Query("SELECT * FROM lotbuilding WHERE idlot = $1 LIMIT 1", IdLot)

	if err != nil {
		fmt.Println("Ошибка в таблице LotBuilding", err)
		return
	}

	if !row.Next() {
		db.Exec("INSERT INTO lotbuilding (idlot, idbuilding) VALUES($1, $2)", IdLot, IdBuilding)
	}
	defer row.Close()
}

func AddDataToProjectBuilding(IdProject int, IdBuilding int) {
	row, err := db.Query("SELECT * FROM lotbuilding WHERE idlot = $1 LIMIT 1", IdProject)

	if err != nil {
		fmt.Println("Ошибка в таблице LotBuilding", err)
		return
	}

	if !row.Next() {
		db.Exec("INSERT INTO lotbuilding (idlot, idbuilding) VALUES($1, $2)", IdProject, IdBuilding)
	}
	defer row.Close()
}

func CheckStringLotBuilding(IsLot bool, IsBuilding bool, Ids ...int) bool {
	if IsLot && IsBuilding {
		row, err := db.Query("SELECT * FROM lotbuilding WHERE idlot = $1 AND idbuilding = $2", Ids[0], Ids[1])
		if err != nil {
			panic(err)
		}

		if row.Next() {
			return true
		} else {
			return false
		}
	} else if IsLot {
		row, err := db.Query("SELECT * FROM lotbuilding WHERE idlot = $1", Ids[0])
		if err != nil {
			panic(err)
		}

		if row.Next() {
			return true
		} else {
			return false
		}
	} else if IsBuilding {
		row, err := db.Query("SELECT * FROM lotbuilding WHERE idbuilding = $1", Ids[0])
		if err != nil {
			panic(err)
		}

		if row.Next() {
			return true
		} else {
			return false
		}
	} else {
		panic("Запрос неправильно составлен!")
	}
}

func CheckStringProjectBuilding(IsProject bool, IsBuilding bool, Ids ...int) bool {
	if IsProject && IsBuilding {
		row, err := db.Query("SELECT * FROM projectbuilding WHERE idproject = $1 AND idbuilding = $2", Ids[0], Ids[1])
		if err != nil {
			panic(err)
		}

		if row.Next() {
			return true
		} else {
			return false
		}
	} else if IsProject {
		row, err := db.Query("SELECT * FROM projectbuilding WHERE idproject = $1", Ids[0])
		if err != nil {
			panic(err)
		}

		if row.Next() {
			return true
		} else {
			return false
		}
	} else if IsBuilding {
		row, err := db.Query("SELECT * FROM projectbuilding WHERE idbuilding = $1", Ids[0])
		if err != nil {
			panic(err)
		}

		if row.Next() {
			return true
		} else {
			return false
		}
	} else {
		panic("Запрос неправильно составлен!")
	}
}
