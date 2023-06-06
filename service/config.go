package service

import (
	"errors"
	"fmt"

	"github.com/go75/gen/common"
	"github.com/go75/gen/utils"
)

type Config struct {
	common.Location
	common.FiltHook
	// 默认继承FiltHook
	NotExtendFiltHook bool
	DisableVO      bool
	DisableTO      bool
	FormTagToUpper bool
	// 是否禁用生成form tag, 默认不禁用
	DisableFormTag bool
	// json tag的字段名是否大写, 默认小写
	JsonTagToUpper bool
	// 是否禁用生成json tag, 默认不禁用
	DisableJsonTag bool
	// 是否禁用生成comment, 默认不禁用
	DisableComment bool
	MethodName        string
}

func (c *Config) Fill() error {
	c.Location.Fill("./service", "service")
	if c.NotExtendFiltHook {
		c.FiltHook.Fill()
	} else {
		c.FiltHook = common.Parser.DaoDependence.FiltHook
	}

	if c.MethodName == "" {
		c.MethodName = "Xxx"
	}

	// 创建service层文件夹
	err := utils.MakeDir(c.SavePath)
	if err != nil {
		fmt.Println("make service dir error: " + err.Error())
		return errors.New("make directory error: " + err.Error())
	}

	if !c.DisableVO {
		dirPath := c.SavePath + "/vo"
		err = utils.MakeDir(dirPath)
		if err != nil {
			fmt.Printf("make vo dir %s error: %s\n", dirPath, err.Error())
			return err
		}
	}

	if !c.DisableTO {
		dirPath := c.SavePath + "/to"
		err = utils.MakeDir(dirPath)
		if err != nil {
			fmt.Printf("make to dir %s error: %s\n", dirPath, err.Error())
			return err
		}
	}

	common.Parser.ServiceDependence = common.NewServiceDependence(c.FiltHook)

	return nil
}