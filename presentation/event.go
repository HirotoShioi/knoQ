package presentation

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/traPtitech/knoQ/domain"

	"github.com/gofrs/uuid"
	"github.com/lestrrat-go/ical"
)

type ScheduleStatus int

const (
	Pending ScheduleStatus = iota + 1
	Attendance
	Absent
)

// EventReqWrite is
//
//go:generate gotypeconverter -s EventReqWrite -d domain.WriteEventParams -o converter.go .
type EventReqWrite struct {
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	AllowTogether bool        `json:"sharedRoom"`
	TimeStart     time.Time   `json:"timeStart"`
	TimeEnd       time.Time   `json:"timeEnd"`
	RoomID        uuid.UUID   `json:"roomId"`
	Place         string      `json:"place"`
	GroupID       uuid.UUID   `json:"groupId"`
	Admins        []uuid.UUID `json:"admins"`
	Tags          []struct {
		Name   string `json:"name"`
		Locked bool   `json:"locked"`
	} `json:"tags"`
	Open bool `json:"open"`
}

type EventTagReq struct {
	Name string `json:"name"`
}

type EventScheduleStatusReq struct {
	Schedule ScheduleStatus `json:"schedule"`
}

// EventDetailRes is experimental
//
//go:generate gotypeconverter -s domain.Event -d EventDetailRes -o converter.go .
type EventDetailRes struct {
	ID            uuid.UUID          `json:"eventId"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	Room          RoomRes            `json:"room"`
	Group         GroupRes           `json:"group"`
	Place         string             `json:"place" cvt:"Room"`
	GroupName     string             `json:"groupName" cvt:"Group"`
	TimeStart     time.Time          `json:"timeStart"`
	TimeEnd       time.Time          `json:"timeEnd"`
	CreatedBy     uuid.UUID          `json:"createdBy"`
	Admins        []uuid.UUID        `json:"admins"`
	Tags          []EventTagRes      `json:"tags"`
	AllowTogether bool               `json:"sharedRoom"`
	Open          bool               `json:"open"`
	Attendees     []EventAttendeeRes `json:"attendees"`
	Model
}

type EventTagRes struct {
	ID     uuid.UUID `json:"tagId" cvt:"Tag"`
	Name   string    `json:"name" cvt:"Tag"`
	Locked bool      `json:"locked"`
}

type EventAttendeeRes struct {
	ID       uuid.UUID      `json:"userId" cvt:"UserID"`
	Schedule ScheduleStatus `json:"schedule"`
}

// EventRes is for multiple response
//
//go:generate gotypeconverter -s domain.Event -d EventRes -o converter.go .
//go:generate gotypeconverter -s []*domain.Event -d []EventRes -o converter.go .
type EventRes struct {
	ID            uuid.UUID          `json:"eventId"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	AllowTogether bool               `json:"sharedRoom"`
	TimeStart     time.Time          `json:"timeStart"`
	TimeEnd       time.Time          `json:"timeEnd"`
	RoomID        uuid.UUID          `json:"roomId" cvt:"Room"`
	GroupID       uuid.UUID          `json:"groupId" cvt:"Group"`
	Place         string             `json:"place" cvt:"Room"`
	GroupName     string             `json:"groupName" cvt:"Group"`
	Admins        []uuid.UUID        `json:"admins"`
	Tags          []EventTagRes      `json:"tags"`
	CreatedBy     uuid.UUID          `json:"createdBy"`
	Open          bool               `json:"open"`
	Attendees     []EventAttendeeRes `json:"attendees"`
	Model
}

func iCalVeventFormat(e *domain.Event, host string) *ical.Event {
	timeLayout := "20060102T150405Z"
	vevent := ical.NewEvent()
	_ = vevent.AddProperty("uid", e.ID.String())
	_ = vevent.AddProperty("dtstamp", time.Now().UTC().Format(timeLayout))
	_ = vevent.AddProperty("dtstart", e.TimeStart.UTC().Format(timeLayout))
	_ = vevent.AddProperty("dtend", e.TimeEnd.UTC().Format(timeLayout))
	_ = vevent.AddProperty("created", e.CreatedAt.UTC().Format(timeLayout))
	_ = vevent.AddProperty("last-modified", e.UpdatedAt.UTC().Format(timeLayout))
	_ = vevent.AddProperty("summary", e.Name)
	e.Description += "\n\n"
	e.Description += "-----------------------------------\n"
	e.Description += "イベント詳細ページ\n"
	e.Description += fmt.Sprintf("%s/events/%v", host, e.ID)
	_ = vevent.AddProperty("description", e.Description)
	_ = vevent.AddProperty("location", e.Room.Place)
	_ = vevent.AddProperty("organizer", e.CreatedBy.DisplayName)

	return vevent
}

func ICalFormat(events []*domain.Event, host string) *ical.Calendar {
	c := ical.New()
	ical.NewEvent()
	tz := ical.NewTimezone()
	_ = tz.AddProperty("TZID", "Asia/Tokyo")
	std := ical.NewStandard()
	_ = std.AddProperty("TZOFFSETFROM", "+9000")
	_ = std.AddProperty("TZOFFSETTO", "+9000")
	_ = std.AddProperty("TZNAME", "JST")
	_ = std.AddProperty("DTSTART", "19700101T000000")
	_ = tz.AddEntry(std)
	_ = c.AddEntry(tz)

	for _, e := range events {
		vevent := iCalVeventFormat(e, host)
		_ = c.AddEntry(vevent)
	}
	return c
}

func GenerateEventWebhookContent(method string, e *EventDetailRes, nofiticationTargets []string, origin string, isMention bool) string {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	timeFormat := "01/02(Mon) 15:04"
	var content string
	switch method {
	case http.MethodPost:
		content = "## イベントが作成されました" + "\n"
	case http.MethodPut:
		content = "## イベントが更新されました" + "\n"
	}
	content += fmt.Sprintf("### [%s](%s/events/%s)", e.Name, origin, e.ID) + "\n"
	content += fmt.Sprintf("- 主催: [%s](%s/groups/%s)", e.GroupName, origin, e.Group.ID) + "\n"
	content += fmt.Sprintf("- 日時: %s ~ %s", e.TimeStart.In(jst).Format(timeFormat), e.TimeEnd.In(jst).Format(timeFormat)) + "\n"
	content += fmt.Sprintf("- 場所: %s", e.Room.Place) + "\n"
	content += "\n"

	if e.TimeStart.After(time.Now()) {
		content += "以下の方は参加予定の入力をお願いします:pray:" + "\n"
		prefix := "@"
		if !isMention {
			prefix = "@."
		}

		sort.Strings(nofiticationTargets)
		for _, nt := range nofiticationTargets {
			content += prefix + nt + " "
		}
		content += "\n\n\n"
	}

	// delete ">" if no description
	if strings.TrimSpace(e.Description) != "" {
		content += "> " + strings.ReplaceAll(e.Description, "\n", "\n> ")
	} else {
		content = strings.TrimRight(content, "\n")
	}

	return content
}
