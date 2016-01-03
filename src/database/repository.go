package database
import (
	"time"
	"fmt"
)


type Repository struct {
	ID int `json:"id"`
	Name string `json:"name,omitempty"`
	Hooks []Hook `json:"hooks,omitempty"`
	Created int64		`json:"created"`
	Updated int64		`json:"updated"`
}

type Repositories []Repository

func GetAllRepositories() []Repository {
	repos := []Repository{}
	DB.Where(&Repository{}).Find(&repos)

	i :=0
	for _, repo := range repos{
		repos[i].Hooks = GetHooksByRepo(repo.ID)
		i++
	}

	fmt.Println(repos)
	return repos
}

func GetRepositoryByName(name string) Repository {
	cfg := Repository{}
	DB.Where(&Repository{Name: name}).First(&cfg)

	return cfg
}

func(a *Repository) SaveRepository() bool{
	a.Created = time.Now().Unix()
	a.Updated = time.Now().Unix()

	if err := DB.Save(a).Error; err != nil {
		return false
	}

	return true
}

func (a *Repository) UpdateRepository() bool {
	a.Updated = time.Now().Unix()

	if err := DB.Save(a).Error; err != nil {
		return false
	}


	return true
}

func (a *Repository) DeleteRepository() bool {
	if err := DB.Delete(a).Error; err != nil {
		return false
	}

	return true
}



type Hook struct {
	ID int `json:"id"`
	Name string `json:"name,omitempty"`
	Exec bool `json:"exec,omitempty"`
	RemoveFolder bool `json:"remove_folder,omitempty"`
	NavyFolder string `json:"navy_folder,omitempty"`
	Created int64		`json:"created"`
	Updated int64		`json:"updated"`
	RepositoryId int `json:"repository_Id"`
}

type Hooks []Hook

func(a *Hook) SaveHook() bool{
	a.Created = time.Now().Unix()
	a.Updated = time.Now().Unix()

	if err := DB.Save(a).Error; err != nil {
		return false
	}

	return true
}

func (a *Hook) UpdateHook() bool {
	a.Updated = time.Now().Unix()

	if err := DB.Save(a).Error; err != nil {
		return false
	}


	return true
}

func (a *Hook) DeleteHook() bool {
	if err := DB.Delete(a).Error; err != nil {
		return false
	}

	return true
}

func GetHooksByRepo(id int) []Hook{
	var hks []Hook

	DB.Where(&Hook{RepositoryId: id}).Find(&hks)

	return hks
}