package LogModel

import "time"

type LogImpl struct {
	ID         string    `gorm:"column:log_id;primary_key;"`
	Name       string    `gorm:"column:log_name"`
	UpdateTime time.Time `gorm:"column:log_time"`
}

func New(name string, updateTime time.Time) *LogImpl {
	return &LogImpl{Name: name, UpdateTime: updateTime}
}

func (LogImpl) TableName() string {
	return "t_log"
}
