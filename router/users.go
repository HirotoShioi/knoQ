package router

import (
	"net/http"
	"room/middleware"
	repo "room/repository"

	"github.com/labstack/echo"
)

// HandleGetUserMe ヘッダー情報からuser情報を取得
func HandleGetUserMe(c echo.Context) error {
	requestUser := middleware.GetRequestUser(c)
	return c.JSON(http.StatusOK, requestUser)
}

// HandleGetUsers ユーザーすべてを取得
func HandleGetUsers(c echo.Context) error {
	users := []repo.User{}
	if err := repo.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, users)
}
