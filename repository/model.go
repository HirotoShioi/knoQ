package repository

import "time"

// User traQユーザー情報構造体
type User struct {
	// TRAQID traQID
	TRAQID string `json:"traq_id" gorm:"type:varchar(32);primary_key"`
	// Admin 管理者かどうか
	Admin bool `gorm:"not null"`
}

// Tag Room Group Event have tags
type Tag struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Official bool   `json:"official"`
	Locked   bool   `json:"locked" gorm:"-"`
	ForRoom  bool   `json:"for_room"`
	ForGroup bool   `json:"for_group"`
	ForEvent bool   `json:"for_event"`
}

// EventTag is many to many table
type EventTag struct {
	TagID   int
	EventID int
	Locked  bool
}

// Room 部屋情報
type Room struct {
	ID        int       `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Place     string    `json:"place" gorm:"type:varchar(16);unique_index:idx_room_unique"`
	Date      string    `json:"date" gorm:"type:DATE; unique_index:idx_room_unique"`
	TimeStart string    `json:"time_start" gorm:"type:TIME; unique_index:idx_room_unique"`
	TimeEnd   string    `json:"time_end" gorm:"type:TIME; unique_index:idx_room_unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Group グループ情報
type Group struct {
	ID             int       `json:"id" gorm:"primary_key; AUTO_INCREMENT"`
	Name           string    `json:"name" gorm:"type:varchar(32);unique;not null"`
	Description    string    `json:"description" gorm:"type:varchar(1024)"`
	Members        []User    `json:"members" gorm:"many2many:group_users; save_associations:false"`
	CreatedBy      User      `json:"created_by" gorm:"foreignkey:CreatedByRefer; not null"`
	CreatedByRefer string    `json:"created_by_refer" gorm:"type:varchar(32);"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Event 予約情報
type Event struct {
	ID            int       `json:"id" gorm:"AUTO_INCREMENT"`
	Name          string    `json:"name" gorm:"type:varchar(32); not null"`
	Description   string    `json:"description" gorm:"type:varchar(1024)"`
	GroupID       int       `json:"group_id,omitempty" gorm:"not null"`
	Group         Group     `json:"group" gorm:"foreignkey:group_id; save_associations:false"`
	RoomID        int       `json:"room_id,omitempty" gorm:"not null"`
	Room          Room      `json:"room" gorm:"foreignkey:room_id; save_associations:false"`
	TimeStart     string    `json:"time_start" gorm:"type:TIME"`
	TimeEnd       string    `json:"time_end" gorm:"type:TIME"`
	CreatedBy     string    `json:"created_by" gorm:"type:varchar(32);"`
	AllowTogether bool      `json:"allow_together"`
	Tags          []Tag     `json:"tags" gorm:"many2many:event_tags; save_associations:false"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
