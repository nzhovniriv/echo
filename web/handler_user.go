package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	User struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}
)

func SaveUser(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func UpdateUser(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func DeleteUser(c echo.Context) error {
	_ = c.Param("id")
	return c.NoContent(http.StatusNoContent)
}
