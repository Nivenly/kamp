package local

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

type KampConfig struct {
	ProjectName string
}

func LocalGit() (*git.Repository, error) {
	r, err := git.PlainOpen("./")
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetLocal() (*KampConfig, error) {
	_, err := LocalGit()
	if err != nil {
		fmt.Println("not at root of a git repository")
	}
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return &KampConfig{ProjectName: wd}, nil
}
