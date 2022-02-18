package repository

import (
	"RPSBackend/model"
	"RPSBackend/utility"
)

type RPSrepository struct{}

func (rps *RPSrepository) ScoreBoard() ([]model.User, error) {

	utility.DbConnect()
	defer utility.DbClose()
	var err error
	score := []model.User{}
	if err = utility.Db.Order("score desc nulls last").Find(&score).Error; err != nil {
		return score, err
	}

	return score, err

}

func (rps *RPSrepository) CreateUser(user model.User) (int, error) {

	utility.DbConnect()
	defer utility.DbClose()
	err := utility.Db.Create(&user).Error
	return user.Id, err

}

func (rps *RPSrepository) GetUser(id int) (model.User, error) {

	var err error
	utility.DbConnect()
	defer utility.DbClose()
	user := model.User{}
	if err = utility.Db.Where("id", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, err

}

func (rps *RPSrepository) UpdateScore(id, score int) (int, error) {

	var err error
	utility.DbConnect()
	defer utility.DbClose()
	currentScore := model.User{}
	if err = utility.Db.Model(&currentScore).Where("id", id).Update("score", score).Error; err != nil {
		return currentScore.Score, err
	}
	return currentScore.Score, err

}
