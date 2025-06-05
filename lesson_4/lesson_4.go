package main

// Подключение MySQL - https://youtu.be/UE2y_3onSeY?si=cuRANgc4REywMhps

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	db, err := sql.Open("mysql", "root:Ar11042008@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Установка данных
	insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES ('Bob', 19)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	// Выборка данных
	res, err := db.Query("SELECT `name`, `age` FROM `users`")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("User: %s with age: %d\n", user.Name, user.Age)
	}

	fmt.Println("Подключено к MySQL")
}
