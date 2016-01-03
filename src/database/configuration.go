package database

import (
	"time"
)



type Configuration struct  {
	ID int   			`json:"id"`
	Name string			`json:"name"`
	Version string		`json:"version"`
	Port string			`json:"port"`
	GitHubToken string	`json:"github_token"`
	Workspace string	`json:"workspace"`
	Created int64		`json:"created"`
	Updated int64		`json:"updated"`
}

func GetConfigurationByName(name string) Configuration {
	 cfg := Configuration{}
	 DB.Where(&Configuration{Name: name}).First(&cfg)

	return cfg
}



func(a *Configuration) SaveConfiguration() bool{
	a.Created = time.Now().Unix()
	a.Updated = time.Now().Unix()

	if err := DB.Save(a).Error; err != nil {
		return false
	}

	return true
}

func (a *Configuration) UpdateConfiguration() bool {
	a.Updated = time.Now().Unix()

	if err := DB.Save(a).Error; err != nil {
		return false
	}


	return true
}

func (a *Configuration) DeleteConfiguration() bool {
	if err := DB.Delete(a).Error; err != nil {
		return false
	}

	return true
}