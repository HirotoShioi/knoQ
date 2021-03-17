// Code generated by gotypeconverter; DO NOT EDIT.
package presentation

import (
	"github.com/gofrs/uuid"
	"github.com/traPtitech/knoQ/domain"
)

func ConvertEventReqWriteTodomainWriteEventParams(src EventReqWrite) (dst domain.WriteEventParams) {
	dst.Name = src.Name
	dst.Description = src.Description
	dst.GroupID = src.GroupID
	dst.RoomID = src.RoomID
	dst.TimeStart = src.TimeStart
	dst.TimeEnd = src.TimeEnd
	dst.Admins = src.Admins
	dst.AllowTogether = src.AllowTogether
	dst.Tags = make([]domain.EventTagParams, len(src.Tags))
	for i := range src.Tags {
		dst.Tags[i].Name = src.Tags[i].Name
		dst.Tags[i].Locked = src.Tags[i].Locked
	}
	return
}

func ConvertGroupReqTodomainWriteGroupParams(src GroupReq) (dst domain.WriteGroupParams) {
	dst.Name = src.Name
	dst.Description = src.Description
	dst.JoinFreely = src.JoinFreely
	dst.Members = src.Members
	dst.Admins = src.Admins
	return
}

func ConvertRoomReqTodomainWriteRoomParams(src RoomReq) (dst domain.WriteRoomParams) {
	dst.Place = src.Place
	dst.TimeStart = src.TimeStart
	dst.TimeEnd = src.TimeEnd
	return
}
func ConvertdomainEventTagToEventTagRes(src domain.EventTag) (dst EventTagRes) {
	dst.Locked = src.Locked
	return
}

func ConvertdomainEventToEventResMulti(src domain.Event) (dst EventResMulti) {
	dst.ID = src.ID
	dst.Name = src.Name
	dst.Description = src.Description
	dst.AllowTogether = src.AllowTogether
	dst.TimeStart = src.TimeStart
	dst.TimeEnd = src.TimeEnd
	dst.RoomID = ConvertdomainRoomTouuidUUID(src.Room)
	dst.GroupID = ConvertdomainGroupTouuidUUID(src.Group)
	dst.Place = src.Room.Place
	dst.GroupName = src.Group.Name
	dst.Admins = make([]uuid.UUID, len(src.Admins))
	for i := range src.Admins {
		dst.Admins[i] = ConvertdomainUserTouuidUUID(src.Admins[i])
	}
	dst.Tags = src.Tags
	dst.CreatedBy = ConvertdomainUserTouuidUUID(src.CreatedBy)
	dst.Model.CreatedAt = src.Model.CreatedAt
	dst.Model.UpdatedAt = src.Model.UpdatedAt
	dst.Model.DeletedAt = src.Model.DeletedAt
	return
}
func ConvertdomainEventToEventResOne(src domain.Event) (dst EventResOne) {
	dst.ID = src.ID
	dst.Name = src.Name
	dst.Description = src.Description
	dst.Room = ConvertdomainRoomToRoomRes(src.Room)
	dst.Group = ConvertdomainGroupToGroupResOne(src.Group)
	dst.TimeStart = src.TimeStart
	dst.TimeEnd = src.TimeEnd
	dst.CreatedBy = ConvertdomainUserToUserRes(src.CreatedBy)
	dst.Admins = make([]UserRes, len(src.Admins))
	for i := range src.Admins {
		dst.Admins[i] = ConvertdomainUserToUserRes(src.Admins[i])
	}
	dst.Tags = make([]EventTagRes, len(src.Tags))
	for i := range src.Tags {
		dst.Tags[i] = ConvertdomainEventTagToEventTagRes(src.Tags[i])
	}
	dst.AllowTogether = src.AllowTogether
	dst.Model.CreatedAt = src.Model.CreatedAt
	dst.Model.UpdatedAt = src.Model.UpdatedAt
	dst.Model.DeletedAt = src.Model.DeletedAt
	return
}

func ConvertdomainGroupToGroupResOne(src domain.Group) (dst GroupResOne) {
	dst.ID = src.ID
	dst.GroupReq.Name = src.Name
	dst.GroupReq.Description = src.Description
	dst.GroupReq.JoinFreely = src.JoinFreely
	dst.GroupReq.Members = make([]uuid.UUID, len(src.Members))
	for i := range src.Members {
		dst.GroupReq.Members[i] = ConvertdomainUserTouuidUUID(src.Members[i])
	}
	dst.GroupReq.Admins = make([]uuid.UUID, len(src.Admins))
	for i := range src.Admins {
		dst.GroupReq.Admins[i] = ConvertdomainUserTouuidUUID(src.Admins[i])
	}
	dst.CreatedBy = ConvertdomainUserTouuidUUID(src.CreatedBy)
	dst.Model.CreatedAt = src.Model.CreatedAt
	dst.Model.UpdatedAt = src.Model.UpdatedAt
	dst.Model.DeletedAt = src.Model.DeletedAt
	return
}
func ConvertdomainGroupTouuidUUID(src domain.Group) (dst uuid.UUID) {
	dst = src.ID
	return
}

func ConvertdomainRoomToRoomRes(src domain.Room) (dst RoomRes) {
	dst.ID = src.ID
	dst.Verified = src.Verified
	dst.RoomReq.Place = src.Place
	dst.RoomReq.TimeStart = src.TimeStart
	dst.RoomReq.TimeEnd = src.TimeEnd
	dst.CreatedBy = ConvertdomainUserToUserRes(src.CreatedBy)
	dst.Model.CreatedAt = src.Model.CreatedAt
	dst.Model.UpdatedAt = src.Model.UpdatedAt
	dst.Model.DeletedAt = src.Model.DeletedAt
	return
}
func ConvertdomainRoomTouuidUUID(src domain.Room) (dst uuid.UUID) {
	dst = src.ID
	return
}

func ConvertdomainTagToTagRes(src domain.Tag) (dst TagRes) {
	dst.ID = src.ID
	dst.Name = src.Name
	dst.Model.CreatedAt = src.Model.CreatedAt
	dst.Model.UpdatedAt = src.Model.UpdatedAt
	dst.Model.DeletedAt = src.Model.DeletedAt
	return
}
func ConvertdomainUserToUserRes(src domain.User) (dst UserRes) {
	dst.ID = src.ID
	dst.Name = src.Name
	dst.DisplayName = src.DisplayName
	dst.Privileged = src.Privileged
	dst.IsTrap = src.IsTrap
	return
}

func ConvertdomainUserTodomainWriteUserParams(src domain.User) (dst domain.WriteUserParams) {
	dst.Name = src.Name
	dst.DisplayName = src.DisplayName
	dst.Icon = src.Icon
	return
}
func ConvertdomainUserTouuidUUID(src domain.User) (dst uuid.UUID) {
	dst = src.ID
	return
}
