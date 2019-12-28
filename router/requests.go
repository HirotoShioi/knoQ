package router

import (
	repo "room/repository"

	"github.com/gofrs/uuid"
)

// GroupReq is group request model
type GroupReq struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	ImageID     string      `json:"image_id"`
	JoinFreely  bool        `json:"join_freely"`
	Members     []uuid.UUID `json:"members"`
}

func formatGroup(req *GroupReq) (g *repo.Group, err error) {
	g = &repo.Group{
		Name:        req.Name,
		Description: req.Description,
		ImageID:     req.ImageID,
		JoinFreely:  req.JoinFreely,
	}

	g.Members = make([]repo.User, 0, len(req.Members))
	for _, v := range req.Members {
		g.Members = append(g.Members, repo.User{ID: v})
	}
	return
}
