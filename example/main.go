package main

import (
	"github.com/go75/gen"
	"github.com/go75/gen/controller"
	"github.com/go75/gen/dao"
	"github.com/go75/gen/model"
	"github.com/go75/gen/service"
)

// 代码基本流程演示
func main() {
	// 创建生成器
	g := gen.New("项目路径", "DSN")
	
	// model层配置
	g.ModelConfig = model.Config{

	}

	// dao层配置
	g.DaoConfig = dao.Config{

	}

	// service层配置
	g.ServiceConfig = service.Config{

	}

	// controller层配置
	g.ControllerConfig = controller.Config{
		
	}

	// 生成代码
	err := g.Run()
	if err != nil {
		panic(err)
	}
}