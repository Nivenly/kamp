package local

import (
	"fmt"
	"os"
	"gopkg.in/src-d/go-git.v4"
)

type KampConfig struct {
	ProjectName       string
}

func LocalGit() (*KampConfig, error) {
	r, err := git.PlainOpen("./")
	if err != nil {
		return nil, err
	}
	fmt.Printf("r: %+v \n", r)

	return nil, nil
}

func GetLocal() (*KampConfig, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return &KampConfig{ProjectName: wd,}, nil
}
