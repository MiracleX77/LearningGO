package user

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Err struct {
	Message string `json:"message"`
}

func CreateUserHandler(c echo.Context) error {
	u := User{}
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO users (name) values ($1)  RETURNING id", u.Name)
	err = row.Scan(&u.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, u)
}

func GetUsersHandler(c echo.Context) error {
	stmt, err := db.Prepare("SELECT id, name FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query all users statment:" + err.Error()})
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all users:" + err.Error()})
	}

	users := []User{}

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})
		}
		users = append(users, u)
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT id, name FROM users WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query user statment:" + err.Error()})
	}

	row := stmt.QueryRow(id)
	u := User{}
	err = row.Scan(&u.ID, &u.Name)
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "user not found"})
	case nil:
		return c.JSON(http.StatusOK, u)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})
	}
}

var db *sql.DB

func InitDB() {
	log.Fatal("start init db")
	var err error
	db, err = sql.Open("postgres", "postgres://nweznkyg:RAg4JxPS1_QM-AYFd49UOVpyx8UFbn1z@rain.db.elephantsql.com/nweznkyg")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	//defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS users ( id SERIAL PRIMARY KEY, name TEXT);
	`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}
}
