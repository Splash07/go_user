package handler

import (
	"net/http"

	"github.com/Splash07/go_user/db"
	"github.com/Splash07/go_user/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// UpdateUser godoc
// @Summary Updater a User record
// @Description accept JSON User payload (excluding ID) and JWT token to update the user information
// @Param userUpdateRequest body model.UserUpdateReq true "User model excluding ID"
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Failure 400 {object} model.Error
// @Router /private/user [put]
func UpdateUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.UserClaims)

	var req model.UserUpdateReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	req.ID = claims.UserID

	res, err := db.UpdateUser(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}
