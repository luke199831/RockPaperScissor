package service

import (
	"RPSBackend/model"
	"RPSBackend/repository"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type RPSService struct{}

var sysRPS = []string{"rock", "paper", "scissor"}

var rpsRepository = repository.RPSrepository{}

func (rps *RPSService) ScoreBoard() ([]model.User, error) {
	var (
		err  error
		user []model.User
	)
	user, err = rpsRepository.ScoreBoard()
	if err != nil {
		err = errors.New("Something went wrong " + err.Error())
		return user, err
	}
	return user, err
}

func (rps *RPSService) CreateUser(user model.User) (int, error) {

	id, err := rpsRepository.CreateUser(user)
	if err != nil {
		err = errors.New("Something went wrong " + err.Error())
		return id, err
	}
	return id, err

}

func (rps *RPSService) UpdateScore(id int, userRPS string) (int, error) {

	user, err := rpsRepository.GetUser(id)
	if err != nil {
		err = errors.New("Something went wrong " + err.Error())
		return user.Score, err
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(sysRPS))
	fmt.Println(i)
	if userRPS == "rock" || userRPS == "paper" || userRPS == "scissor" {

		if userRPS == sysRPS[i] {
			user.Score = user.Score
		} else if userRPS == "rock" && sysRPS[i] == "scissor" {
			user.Score += 1
		} else if userRPS == "paper" && sysRPS[i] == "rock" {
			user.Score += 1
		} else if userRPS == "scissor" && sysRPS[i] == "paper" {
			user.Score += 1
		} else {
			user.Score -= 1
		}

	} else {
		return user.Score, errors.New("select rock, paper or scissor")
	}

	updateScore, updateErr := rpsRepository.UpdateScore(id, user.Score)
	if updateErr != nil {
		updateErr = errors.New("Error in update " + updateErr.Error())
		return updateScore, updateErr
	}

	return updateScore, updateErr

}
