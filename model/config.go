package model

import (
	"database/sql"
	"errors"

	"github.com/go75/gen/common"
	"github.com/go75/gen/utils"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	common.Location
	common.FiltHook
	// json tag的字段名是否大写, 默认小写
	JsonTagToUpper bool
	// 是否禁用生成json tag, 默认不禁用
	DisableJsonTag bool
	// 是否禁用生成comment, 默认不禁用
	DisableComment bool
	db       *sql.DB
}

func (c *Config) Fill() error {
	c.Location.Fill("./model", "model")
	c.FiltHook.Fill()

	// 创建model层文件夹
	err := utils.MakeDir(c.SavePath)
	if err != nil {
		return errors.New("make directory error: " + err.Error())
	}

	common.Parser.ModelDependence = common.NewModelDependence(common.Parser.ProjectPrefix + "/" + c.PackageName, c.PackageName, c.FiltHook)

	return nil
}