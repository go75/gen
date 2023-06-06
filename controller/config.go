package controller

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
	DisableC   bool
	DisableR   bool
	DisableU   bool
	DisableD   bool
	DisableRes bool
}

func (c *Config) Fill() error {
	c.Location.Fill("./controller", "controller")
	if c.NotExtendFiltHook {
		c.FiltHook.Fill()
	} else {
		c.FiltHook = common.Parser.ServiceDependence.FiltHook
	}

	// 创建service层文件夹
	err := utils.MakeDir(c.SavePath)
	if err != nil {
		fmt.Println("make controller dir error: ", err.Error())
		return errors.New("make directory error: " + err.Error())
	}

	if !c.DisableRes {
		dirPath := c.SavePath + "/res"
		err = utils.MakeDir(dirPath)
		if err != nil {
			return err
		}
	}

	return nil
}