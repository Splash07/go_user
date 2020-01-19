package handler

import (
	"net/http"

	"github.com/Splash07/go_user/db"
	"github.com/Splash07/go_user/model"
	"github.com/labstack/echo/v4"
)

// HealthCheck godoc
// @Summary Check API's health
// @Description send a small request and check the response
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 400 {object} model.Error
// @Router /public/health [get]
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// RegisterUser godoc
// @Summary Register for a new user
// @Description create new User model
// @Param user body model.User true "User model excluding ID"
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {object} model.Error
// @Router /public/user/register [post]
func RegisterUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.AddUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}

// LoginUser godoc
// @Summary Login and obtain JWT Token
// @Description get JWT Token with email and password as payload
// @Param loginRequest body model.LoginReq true "Login Request containing email and password"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.LoginResp
// @Failure 400 {object} model.Error
// @Router /public/user/login [patch]
func LoginUser(c echo.Context) error {
	var req model.LoginReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.LoginUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}
