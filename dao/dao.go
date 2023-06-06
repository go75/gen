package dao

import (
	"fmt"
	"github.com/go75/gen/common"
	"github.com/go75/gen/utils"
)

type DaoEngine struct {
	*Config
}

func New(config *Config) *DaoEngine {
	return &DaoEngine{
		Config: config,
	}
}

func (e *DaoEngine) Run() error {
	var err error

	if !e.DisableInitDB {
		err = e.genDB()
		if err != nil {
			return err
		}
	}

	err = e.genCRUDs()

	return err
}

func (e *DaoEngine) genDB() error {

	content := fmt.Sprintf(`
package %s

import (
	"%s"
	"%s"
)

var db *gorm.DB

func InitDB() (err error) {
	dsn := "%s"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
`, 
	e.PackageName, e.Config.SqlPackageName, e.Config.OrmPackageName, common.Parser.DSN)

	filepath := e.SavePath + "/init.go"
	
	err := utils.WriteGoFile(filepath, content)
	
	if err != nil {
		return err
	}
	
	return nil
}

func (e *DaoEngine) genCRUDs() error {
	for structName, fields := range common.Parser.Structs {
		if e.FiltHook(structName) {
			continue
		}

		// generate crud
		content := e.generateCRUD(structName, fields)
		
		filepath := e.SavePath + "/" + fields[0].TableName + ".go"
	
		err := utils.WriteGoFile(filepath, content)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *DaoEngine) generateCRUD(structName string, fields []*common.Field) string {
	content := common.NewContent(e.Location)
	
	content.WriteString(`import "`)
	content.WriteString(common.Parser.ModelDependence.PackagePath)
	content.WriteString("\"\n\n")

	if !e.DisableC {
		content.WriteString(e.c(structName))
	}

	if !e.DisableR {
		content.WriteString(e.r(structName))
	}

	if !e.DisableU {
		content.WriteString(e.u(structName))

	}

	if !e.DisableD {
		content.WriteString(e.d(structName))
	}

	return content.String()
}

func (e *DaoEngine) c(structName string) string {
	fullName := common.Parser.ModelDependence.PackageName + "." + structName
	return fmt.Sprintf(`

func Create%s(info *%s) error {
	return db.Create(info).Error
}

func CreateAll%s(infos *[]%s) error {
	return db.Create(infos).Error
}

`	,structName, fullName,
	structName, fullName)
}

func (e *DaoEngine) r(structName string) string {
	fullName := common.Parser.ModelDependence.PackageName + "." + structName
	return fmt.Sprintf(`

func Query%s(cond, result *%s) error {
	return db.Model(cond).First(result).Error
}

func QueryAll%s(cond, results *[]%s) error {
	return db.Model(cond).Find(results).Error
}

`,
	structName, fullName,
	structName, fullName)
}

func (e *DaoEngine) u(structName string) string {
	fullName := common.Parser.ModelDependence.PackageName + "." + structName
	return fmt.Sprintf(`
	
func Update%s(old, new *%s) error {
	return db.Model(old).Updates(*new).Error
}

`, 
	structName, fullName)
}

func (e *DaoEngine) d(structName string) string {
	fullName := common.Parser.ModelDependence.PackageName + "." + structName
	return fmt.Sprintf(`

func Delete%s(cond *%s) error {
	return db.Delete(cond).Error
}

`, 
	structName, fullName)
}