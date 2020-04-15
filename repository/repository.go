// Package repository is
package repository

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 新たにモデル(テーブル)を定義した場合はここに追加する事
var tables = []interface{}{
	User{},
	Room{},
	Group{},
	Event{},
	Tag{},
	EventTag{},
	GroupUsers{},
}

type Repository interface {
	UserRepository
	GroupRepository
	RoomRepository
	EventRepository
	TagRepository
}

// GormRepository implements Repository interface
type GormRepository struct {
	DB *gorm.DB
}

type TraQVersion int64

const (
	TraQv1 TraQVersion = iota
	TraQv3
)

type TraQRepository struct {
	Version TraQVersion
	Host    string
	Token   string
}

type GoogleAPIRepository struct {
	Config     *jwt.Config
	Client     *http.Client
	CalendarID string
}

var (
	MARIADB_HOSTNAME = os.Getenv("MARIADB_HOSTNAME")
	MARIADB_DATABASE = os.Getenv("MARIADB_DATABASE")
	MARIADB_USERNAME = os.Getenv("MARIADB_USERNAME")
	MARIADB_PASSWORD = os.Getenv("MARIADB_PASSWORD")

	DB        *gorm.DB
	logger, _ = zap.NewDevelopment()
)

func (repo *GoogleAPIRepository) Setup() {
	repo.Client = repo.Config.Client(oauth2.NoContext)
}

func (repo *TraQRepository) getBaseURL() string {
	var traQEndPointVersion = [2]string{
		"/1.0",
		"/v3",
	}

	return repo.Host + traQEndPointVersion[repo.Version]
}

func (repo *TraQRepository) getRequest(path string) ([]byte, error) {
	if repo.Token == "" {
		return nil, ErrForbidden
	}
	req, err := http.NewRequest(http.MethodGet, repo.getBaseURL()+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+repo.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 {
		// TODO consider 300
		switch res.StatusCode {
		case 401:
			return nil, ErrForbidden
		case 403:
			return nil, ErrForbidden
		case 404:
			return nil, ErrNotFound
		default:
			return nil, errors.New(http.StatusText(res.StatusCode))
		}
	}
	return ioutil.ReadAll(res.Body)
}

// SetupDatabase set up DB and crate tables
func SetupDatabase() (*gorm.DB, error) {
	var err error
	//tmp
	if MARIADB_HOSTNAME == "" {
		MARIADB_HOSTNAME = ""
	}
	if MARIADB_DATABASE == "" {
		MARIADB_DATABASE = "room"
	}
	if MARIADB_USERNAME == "" {
		MARIADB_USERNAME = "root"
	}

	if MARIADB_PASSWORD == "" {
		MARIADB_PASSWORD = "password"
	}

	// データベース接続
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", MARIADB_USERNAME, MARIADB_PASSWORD, MARIADB_HOSTNAME, MARIADB_DATABASE))
	if err != nil {
		return DB, err
	}
	if err := initDB(DB); err != nil {
		return DB, err
	}
	return DB, nil
}

// initDB データベースのスキーマを更新
func initDB(db *gorm.DB) error {
	// gormのエラーの上書き
	gorm.ErrRecordNotFound = ErrNotFound

	// テーブルが無ければ作成
	if err := db.AutoMigrate(tables...).Error; err != nil {
		return err
	}
	return nil
}
