package controller

import (
	"echo-api/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	result, err := service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func CreateUsers(c echo.Context) error {
	nama := c.FormValue("name")
	email := c.FormValue("email")
	kelas := c.FormValue("class")
	strc, err := strconv.Atoi(kelas)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := service.CreateUsers(nama, email, strc)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUsersById(c echo.Context) error {

	Id := c.FormValue("id")
	strconv, err := strconv.Atoi(Id)
	if err != nil {
		return err
	}

	result, err := service.DeletFromUsersById(strconv)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUsers(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("name")
	email := c.FormValue("email")
	kelas := c.FormValue("class")

	strid, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	strclass, err := strconv.Atoi(kelas)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := service.UpdateUsers(strid, nama, email, strclass)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
