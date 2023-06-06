package gen

import (
	"fmt"
	"github.com/go75/gen/common"
	"github.com/go75/gen/controller"
	"github.com/go75/gen/dao"
	"github.com/go75/gen/model"
	"github.com/go75/gen/service"
	"github.com/go75/gen/utils"
)

type Generator struct {
	ModelConfig model.Config
	DaoConfig dao.Config
	ServiceConfig service.Config
	ControllerConfig controller.Config
}

func New(projectPrefix string, dsn string) *Generator {
	common.Parser.ProjectPrefix = projectPrefix
	common.Parser.DSN = dsn
	return &Generator{}
}

func (g *Generator) Run() error {
	g.fill()

	modelEngine := model.New(&g.ModelConfig)
	err := modelEngine.Run()
	if err != nil {
		return fmt.Errorf("model engine run error: %s", err)
	}

	daoEngine := dao.New(&g.DaoConfig)
	err = daoEngine.Run()
	if err != nil {
		return fmt.Errorf("dao engine run error: %s", err)
	}

	serviceEngine := service.New(&g.ServiceConfig)
	err = serviceEngine.Run()
	if err != nil {
		return fmt.Errorf("service engine run error: %s", err)
	}

	controllerEngine := controller.New(&g.ControllerConfig)
	err = controllerEngine.Run()
	if err != nil {
		return fmt.Errorf("controller engine run error: %s", err)
	}

	err = utils.GoModTidy()
	if err != nil {
		fmt.Println("go mod tidy error: ", err)
		return fmt.Errorf("go mod tidy error: %s", err)
	}

	return nil
}

func (g *Generator) fill() {
	g.ModelConfig.Fill()
	g.DaoConfig.Fill()
	g.ServiceConfig.Fill()
	g.ControllerConfig.Fill()
}