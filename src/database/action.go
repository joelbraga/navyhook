package database

import (
	"time"
)


const(
	PROCESSING = "PROCESSING"
	ERROR = "ERROR"
	FINISHED = "FINISHED"
)

type Action struct  {
	Id int64 			`json:"id"`
	UserName string 	`json:"user_name"`
	AvatarURL string	`json:"avatar_url"`
	Event string		`json:"event"`
	Repository string	`json:"repository"`
	State string		`json:"state"`
	Logger string		`json:"logger"`
	Info string			`json:"info"`
	Created int64		`json:"created"`
	Updated int64		`json:"updated"`
}


func (a *Action) GetById() bool {
	err := DB.Where(&Action{Id: a.Id}).First(a).Error

	if err != nil {
		return false
	}

	return true
}

func (a *Action) GetByRepository() []Action {
	var actions []Action

	DB.Where(&Action{Repository: a.Repository}).Limit(10).Order("id desc").Find(&actions)

	return actions
}

func (a *Action) GetAll() []Action {
	var actions []Action

	DB.Where(&Action{}).Limit(10).Order("id desc").Find(&actions)

	return actions
}


func(a *Action) Save() bool{
	a.Created = time.Now().Unix()
	a.Updated = time.Now().Unix()
	a.State = PROCESSING

	if err := DB.Save(a).Error; err != nil {
		return false
	}

	return true
}

func (a *Action) Update() bool {
	a.Updated = time.Now().Unix()

	if err := DB.Save(a).Error; err != nil {
		return false
	}

	return true
}

func (a *Action) Delete() bool {
	if err := DB.Delete(a).Error; err != nil {
		return false
	}

	return true
}

func (a *Action) OnError(err string) bool{
	a.State = ERROR
	a.Logger = err
	return a.Update()
}

func (a *Action) OnProcessing() bool{
	a.State = PROCESSING
	return a.Update()
}

func (a *Action) OnSuccess(log string) bool{
	a.State = FINISHED
	a.Logger = log
	return a.Update()
}