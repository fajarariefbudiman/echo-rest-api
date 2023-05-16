package service

import (
	"echo-api/db"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Users struct {
	Id    int64  `json:"id"`
	Name  string `validate:"required" json:"name"`
	Email string `validate:"required,email" json:"email"`
	Class int    `validate:"required" json:"class"`
}

func GetAllUsers() (Response, error) {
	var get Users
	var getAll []Users
	var res Response

	cond := db.NewData()
	sqlstmt := "select * from users"

	rows, err := cond.Query(sqlstmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Name, &get.Email, &get.Class)
		if err != nil {
			return res, err
		}

		getAll = append(getAll, get)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = getAll

	return res, nil
}

func CreateUsers(name string, email string, class int) (Response, error) {
	var response Response
	v := validator.New()

	Use := Users{
		Name:  name,
		Email: email,
		Class: class,
	}

	err := v.Struct(Use)
	if err != nil {
		return response, err
	}

	cond := db.NewData()
	sqlstmnt := "insert into users (name,email,class) values(?,?,?)"

	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}

	result, err := stmt.Exec(name, email, class)
	if err != nil {
		return response, err
	}

	lastinsert, err := result.LastInsertId()
	if err != nil {
		return response, err
	}

	response.Status = http.StatusOK
	response.Message = "Succes Create"
	response.Data = map[string]int64{
		"lastinsertId": lastinsert,
	}
	return response, err

}

func DeletFromUsersById(Id int) (Response, error) {

	var response Response

	cond := db.NewData()
	sqlstmnt := "delete from users where id=?"

	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}

	result, err := stmt.Exec(Id)
	if err != nil {
		return response, err
	}

	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return response, err
	}

	response.Status = http.StatusOK
	response.Message = "Succes Delete"
	response.Data = map[string]int64{
		"rowsaffected": rowsaffected,
	}
	return response, err

}

func UpdateUsers(Id int, Name string, Email string, Class int) (Response, error) {
	var response Response

	cond := db.NewData()
	sqlstmnt := "update users set name=?, email=?, class=? where id=?"

	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}

	result, err := stmt.Exec(Name, Email, Class, Id)
	if err != nil {
		return response, err
	}

	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return response, err
	}

	response.Status = http.StatusOK
	response.Message = "Succes Update"
	response.Data = map[string]int64{
		"rowsaffected": rowsaffected,
	}
	return response, err

}


