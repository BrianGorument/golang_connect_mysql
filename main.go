package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type King struct {
	Name          string `json:"name"`
	Date_of_born  string `json:"date_of_born"`
	Date_of_death string `json:"date_of_death"`
	Title         string `json:"title"`
}

func main() {
	fmt.Println("Go mySQl")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/verycooldatabase")
	defer db.Close()
	//panic handler
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connect to database")

	//insert
	insert, err := db.Query("Insert into the_king VALUES('Heculless','2024-01-01','2050-01-01','the mighty')")
	defer insert.Close()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Suceessfully insert data")

	//get
	getAlldata, err := db.Query("Select name,date_of_born,date_of_death,title FROM `the_king`")
	defer getAlldata.Close()
	if err != nil {
		panic(err.Error())
	}

	for getAlldata.Next() {
		var king King
		err = getAlldata.Scan(&king.Name, &king.Date_of_born, &king.Date_of_death, &king.Title)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("----------------------")
		fmt.Println(king.Name)
		fmt.Println(king.Date_of_born)
		fmt.Println(king.Date_of_death)
		fmt.Println(king.Title)
		fmt.Println("----------------------")
	}

	//Update
	//fmt.Print("Input a king name : ")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// inputText1 := scanner.Text()

	// fmt.Print("Replace a new name for king : ")
	// scanner.Scan()
	// inputText2 := scanner.Text()

	// udpate, err := db.Prepare("Update `the_king` SET `name`= ? Where `name`= ? ")
	// defer udpate.Close()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// udpate.Exec(&inputText2, &inputText1)
	// fmt.Println("Successfully Update data")

	//delete

	fmt.Print("id to delete ? = ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	deletetext1 := scanner.Text()

	delete, err := db.Prepare("DELETE From `the_king` where `id`= ?")
	defer delete.Close()
	delete.Exec(deletetext1)

	fmt.Println("Successfully Delete item")
}
