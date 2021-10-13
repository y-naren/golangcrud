package database

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type UserData struct {
	Id          int
	Name        string
	DOB         string
	Address     string
	Description string
	Created_At  string
}

func GetUserData() []UserData {
	// mysql -u root -p			-- open database
	// show database;
	// show tables;
	// use [database_name]
	// select * from [table_name];
	/*
		create table user (id INT unsigned NOT NULL Primary Key AUTO_INCREMENT,name varchar(100) NOT NULL, dob date NOT NULL, address varchar(255),description varchar(500),createdAt datetime NOT NULL);
	*/
	//decribe [tablename]
	/*
		Insert into user (name,dob,address,description,createdAt) values ('amrendra','1998-07-10','Ballia','Good Boy','2015-01-03 00:00:00');
	*/

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goRestService")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	var allUser []UserData

	se, err := db.Query("select * from user") //.Scan(&user.id, &user.name, &user.dob, &user.address, &user.description, &user.created_at)

	for se.Next() {
		var user UserData
		// for each row, scan the result into our tag composite object
		err = se.Scan(&user.Id, &user.Name, &user.DOB, &user.Address, &user.Description, &user.Created_At)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		allUser = append(allUser, user)
	}
	return allUser
}
func GetUserById(id string) UserData {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goRestService")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	var user UserData

	err = db.QueryRow("select * from user where id=?", i).Scan(&user.Id, &user.Name, &user.DOB, &user.Address, &user.Description, &user.Created_At)
	if err != nil {
		log.Println(err)
	}
	return user
}
func CreateUserData(user UserData) int64 {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goRestService")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	insForm, err := db.Prepare("INSERT INTO user(name,dob,address,description,createdAt) VALUES(?,?,?,?,?)")
	res, err := insForm.Exec(user.Name, user.DOB, user.Address, user.Description, user.Created_At)
	rows, _ := res.LastInsertId()
	return rows
}
func UpdateUserData(user UserData) int64 {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goRestService")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	insForm, err := db.Prepare("UPDATE user SET name=? WHERE id=?")
	res, err := insForm.Exec(user.Name, user.Id)
	rows, _ := res.RowsAffected()
	return rows
}
func DeleteUserData(user UserData) int64 {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goRestService")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	insForm, err := db.Prepare("DELETE FROM user WHERE id=?")
	res, err := insForm.Exec(user.Id)
	rows, _ := res.RowsAffected()
	return rows
}
