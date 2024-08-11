package main

import (
	"context"
	"gf2gacha/logic"
	"gf2gacha/model"
	"gf2gacha/util"
	"github.com/sirupsen/logrus"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetGameInfo() model.Info {
	info, err := util.GetGameInfo()
	if err != nil {
		logrus.Error(err)
		return model.Info{}
	}
	return info
}

func (a *App) GetUserList() []string {
	userList, err := logic.GetUserList()
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return userList
}

func (a *App) GetPoolInfo(uid string, poolType int64) model.Pool {
	pool, err := logic.GetPoolInfo(uid, poolType)
	if err != nil {
		logrus.Error(err)
		return model.Pool{}
	}
	return pool
}

func (a *App) IncrementalUpdatePoolInfo() string {
	uid, err := logic.IncrementalUpdatePoolInfo()
	if err != nil {
		logrus.Error(err)
		return ""
	}
	return uid
}
