package dao

import (
	_ "database/sql"
	log "github.com/Sirupsen/logrus"
	"github.com/ggoblin/goblin/platform/libs/model"
	"github.com/ggoblin/goblin/platform/libs/utils"
	_ "github.com/lib/pq"
)

func GetAllIterations() ([]model.Iteration, error) {
	db, err := utils.GetDefaultDb()
	log.Debug("Start Get all iterations.")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var iterations []model.Iteration
	db.Find(&iterations)
	log.Debug("Start Get all iterations.")
	return iterations, nil
}