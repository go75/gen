package dao

import (
	"errors"

	"github.com/go75/gen/common"
	"github.com/go75/gen/utils"
)

type Config struct {
	common.Location
	common.FiltHook
	// 默认继承FiltHook
	NotExtendFiltHook bool
	DisableC          bool
	DisableR          bool
	DisableU          bool
	DisableD          bool
	DisableInitDB     bool
	OrmPackageName    string
	SqlPackageName    string
}

func (c *Config) Fill() error {
	
	c.Location.Fill("./dao", "dao")
	if c.NotExtendFiltHook {
		c.FiltHook.Fill()
	} else {
		c.FiltHook = common.Parser.ModelDependence.FiltHook
	}
	
	if c.SqlPackageName == "" {
		c.SqlPackageName = "gorm.io/driver/mysql"
	}

	if c.OrmPackageName == "" {
		c.OrmPackageName = "gorm.io/gorm"
	}

	// 创建dao层文件夹
	err := utils.MakeDir(c.SavePath)
	if err != nil {
		return errors.New("make directory error: " + err.Error())
	}

	common.Parser.DaoDependence = common.NewDaoDependence(c.FiltHook)
	return nil
}