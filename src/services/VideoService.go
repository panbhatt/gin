package services

import (
	"fmt"
)

type VideoService struct {
	name string
}

func GetVideoService() *VideoService {
	return &VideoService{
		name: "Pnkaj Bhatt. ",
	}
}

func (vService *VideoService) AddVideo(l int) error {
	if l != -1 {
		return nil
	} else {
		return fmt.Errorf("New Error occured. ")
	}
}
