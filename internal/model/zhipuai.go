package model

import (
	"time"
)

// ZhipuaiMessage [...]
type ZhipuaiMessage struct {
	ID        int64     `gorm:"primaryKey;column:id" json:"id"`
	RoleName  string    `gorm:"column:role_name" json:"role_name"`
	Content   string    `gorm:"column:content" json:"content"`
	UUID      string    `gorm:"column:uuid" json:"uuid"`
	Deleted   int64     `gorm:"column:deleted" json:"deleted"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *ZhipuaiMessage) TableName() string {
	return "zhipuai_message"
}

// ZhipuaiMessageColumns get sql column name.获取数据库列名
var ZhipuaiMessageColumns = struct {
	ID        string
	RoleName  string
	Content   string
	UUID      string
	Deleted   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	RoleName:  "role_name",
	Content:   "content",
	UUID:      "uuid",
	Deleted:   "deleted",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// ZhipuaiSession [...]
type ZhipuaiSession struct {
	ID        int64     `gorm:"primaryKey;column:id" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	UUID      string    `gorm:"column:uuid" json:"uuid"`
	Deleted   int64     `gorm:"column:deleted" json:"deleted"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *ZhipuaiSession) TableName() string {
	return "zhipuai_session"
}

// ZhipuaiSessionColumns get sql column name.获取数据库列名
var ZhipuaiSessionColumns = struct {
	ID        string
	Title     string
	UUID      string
	Deleted   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Title:     "title",
	UUID:      "uuid",
	Deleted:   "deleted",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// ZhipuaiSendLog [...]
type ZhipuaiSendLog struct {
	ID        int64     `gorm:"primaryKey;column:id" json:"id"`
	Req       string    `gorm:"column:req" json:"req"`
	Resp      string    `gorm:"column:resp" json:"resp"`
	UUID      string    `gorm:"column:uuid" json:"uuid"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *ZhipuaiSendLog) TableName() string {
	return "zhipuai_send_log"
}

// ZhipuaiSendLogColumns get sql column name.获取数据库列名
var ZhipuaiSendLogColumns = struct {
	ID        string
	Req       string
	Resp      string
	UUID      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Req:       "req",
	Resp:      "resp",
	UUID:      "uuid",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}
