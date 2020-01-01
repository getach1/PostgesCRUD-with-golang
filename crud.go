package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var conn sql.DB

func init() {
	conn, err := sql.Open("postgres", "user=postgres dbname=first_time password=gechman sslmode=disable")
	checkErr(err)
	defer conn.Close()
}
func checkErr(err error) {

	if err == sql.ErrNoRows {
		fmt.Println("No rows were returned!")
		return
	}
	if err != nil {
		panic(err)
	}
}

type Users struct {
	Name  string
	ID    int
	Phone string
	Email string
}

func main() {
	user1 := Users{Name: "get", Phone: "0922222222222", Email: "dsd@fjdf.com"}
	user2 := Users{Name: "get", Phone: "0922222222222", Email: "dsd@fjdf.com"}
	user3 := Users{Name: "get", Phone: "0922222222222", Email: "dsd@fjdf.com"}
	user4 := Users{Name: "get", Phone: "0922222222222", Email: "dsd@fjdf.com"}
	user5 := Users{Name: "get", Phone: "0922222222222", Email: "dsd@fjdf.com"}
	user6 := Users{Name: "get", Phone: "0922222222222", Email: "dsd@fjdf.com"}
	user1.Insert()
	user2.Insert()
	user3.Insert()
	user4.Insert()
	user5.Insert()
	user6.Insert()
	//user1.Delete()
	//user1.UpdateEmail()
	//user1.UpdateName()
	//user1.UpdatePhone()

	//userList:=UsersList()
	//for user:=range userList{
	//fmt.Print("Name=%s Id=%s Email=%s Phone=%s",user.Name,string(user.ID),user.Email,user.Phone)
	//}
}
func (user Users) Delete() {
	stmt, err := conn.Prepare("delete from users where id=$1;")
	checkErr(err)
	_, err = stmt.Exec(user.ID)
	checkErr(err)
}
func (user *Users) Insert() {
	statement := "insert into users values($1,$2,$3)  returing id;"
	//Note the following
	//I am not using stmt.Exec() because i need the row num to assign to my user.ID
	err := conn.QueryRow(statement, user.Name, user.Phone, user.Email).Scan(&user.ID) ///////////^^^^^^^^^^
	checkErr(err)
}
func (user *Users) UpdateName() {
	stmt, err := conn.Prepare("UPDATE Users set name =$1 where id=$2 ;")
	checkErr(err)
	_, err = stmt.Exec(user.Name, user.ID)
	checkErr(err)
}
func (user *Users) UpdateEmail() {
	stmt, err := conn.Prepare("UPDATE Users set email =$1 where id=$2 ;")
	checkErr(err)
	_, err = stmt.Exec(user.Email, user.ID)
	checkErr(err)
}
func (user *Users) UpdatePhone() {
	stmt, err := conn.Prepare("UPDATE Users set phone =$1 where id=$2 ;")
	checkErr(err)
	_, err = stmt.Exec(user.Name, user.ID)
	checkErr(err)
}

func UsersList() (ul []Users) {
	stmt, err := conn.Prepare("SELECT * from users;")
	checkErr(err)
	users, err := stmt.Query()
	var user []Users
	for users.Next() {
		user1 := Users{}
		err := users.Scan(user1.Name, user1.ID, user1.Phone, user1.Email)
		checkErr(err)
		user = append(user, user1)
	}

	return user
}
