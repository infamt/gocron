package models

import "time"

// 行为日志
type ActionLog struct {
	BaseModel
	Id         int       `json:"id" xorm:"pk autoincr notnull "`
	Username   string    `json:"username" xorm:"varchar(32) notnull"`
	Ip         string    `json:"ip" xorm:"varchar(15) not null"`
	TaskId     int       `json:"task_id" xorm:"int(11)"`
	TaskName   string    `json:"task_name" xorm:"varchar(255)"`
	ActionType string    `json:"task_type" xorm:"varchar(20)"`
	TaskHost   string    `json:"task_host" xorm:"varchar(255)"`
	Action     string    `json:"action" xorm:"varchar(255) notnull"`
	Created    time.Time `json:"created" xorm:"datetime notnull created"`
}

func (log *ActionLog) Create() (insertId int, err error) {
	_, err = Db.Insert(log)
	if err == nil {
		insertId = log.Id
	}

	return
}

func (log *ActionLog) List(params CommonMap) ([]ActionLog, error) {
	log.parsePageAndPageSize(params)
	list := make([]ActionLog, 0)
	err := Db.Desc("id").Limit(log.PageSize, log.pageLimitOffset()).Find(&list)

	return list, err
}

func (log *ActionLog) Total() (int64, error) {
	return Db.Count(log)
}
