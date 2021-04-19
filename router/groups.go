package router

import (
	"net/http"

	"github.com/traPtitech/knoQ/presentation"

	"github.com/labstack/echo/v4"
)

// HandlePostGroup グループを作成
func (h *Handlers) HandlePostGroup(c echo.Context) error {
	var req presentation.GroupReq
	if err := c.Bind(&req); err != nil {
		return badRequest(err, message(err.Error()))
	}
	groupParams := presentation.ConvGroupReqTodomainWriteGroupParams(req)

	group, err := h.Repo.CreateGroup(groupParams, getConinfo(c))
	if err != nil {
		return judgeErrorResponse(err)
	}
	res := presentation.ConvdomainGroupToGroupRes(*group)
	return c.JSON(http.StatusCreated, res)
}

// HandleGetGroup グループを一件取得
func (h *Handlers) HandleGetGroup(c echo.Context) error {
	groupID, err := getPathGroupID(c)
	if err != nil {
		return notFound(err)
	}

	group, err := h.Repo.GetGroup(groupID, getConinfo(c))
	if err != nil {
		return judgeErrorResponse(err)
	}

	res := presentation.ConvdomainGroupToGroupRes(*group)
	return c.JSON(http.StatusOK, res)
}

// HandleGetGroups グループを取得
func (h *Handlers) HandleGetGroups(c echo.Context) error {
	groups, err := h.Repo.GetAllGroups(getConinfo(c))
	if err != nil {
		return judgeErrorResponse(err)
	}
	res := presentation.ConvSPdomainGroupToSPGroupRes(groups)
	return c.JSON(http.StatusOK, res)
}

// HandleDeleteGroup グループを削除
func (h *Handlers) HandleDeleteGroup(c echo.Context) error {
	groupID, err := getPathGroupID(c)
	if err != nil {
		return notFound(err)
	}

	if err := h.Repo.DeleteGroup(groupID, getConinfo(c)); err != nil {
		return judgeErrorResponse(err)
	}

	return c.NoContent(http.StatusNoContent)
}

// HandleUpdateGroup 変更できるものはpostと同等
func (h *Handlers) HandleUpdateGroup(c echo.Context) error {
	groupID, err := getPathGroupID(c)
	if err != nil {
		return notFound(err)
	}

	var req presentation.GroupReq
	if err := c.Bind(&req); err != nil {
		return badRequest(err, message(err.Error()))
	}
	groupParams := presentation.ConvGroupReqTodomainWriteGroupParams(req)

	res, err := h.Repo.UpdateGroup(groupID, groupParams, getConinfo(c))
	if err != nil {
		return judgeErrorResponse(err)
	}
	return c.JSON(http.StatusOK, res)
}

func (h *Handlers) HandleAddMeGroup(c echo.Context) error {
	groupID, err := getPathGroupID(c)
	if err != nil {
		return notFound(err)
	}

	err = h.Repo.AddMeToGroup(groupID, getConinfo(c))
	if err != nil {
		return judgeErrorResponse(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handlers) HandleDeleteMeGroup(c echo.Context) error {
	groupID, err := getPathGroupID(c)
	if err != nil {
		return notFound(err)
	}

	err = h.Repo.DeleteMeGroup(groupID, getConinfo(c))
	if err != nil {
		return judgeErrorResponse(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handlers) HandleGetMeGroupIDs(c echo.Context) error {
	userID, _ := getRequestUserID(c)

	groupIDs, err := h.Repo.GetUserBelongingGroupIDs(userID, getConinfo(c))
	if err != nil {
		return judgeErrorResponse(err)
	}
	return c.JSON(http.StatusOK, groupIDs)
}

func (h *Handlers) HandleGetGroupIDsByUserID(c echo.Context) error {
	userID, err := getPathUserID(c)
	if err != nil {
		return notFound(err, message(err.Error()))
	}

	groupIDs, err := h.Repo.GetUserBelongingGroupIDs(userID, getConinfo(c))
	if err != nil {
		return judgeErrorResponse(err)
	}

	return c.JSON(http.StatusOK, groupIDs)
}
