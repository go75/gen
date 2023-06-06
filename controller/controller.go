package controller

import (
	"fmt"
	"strings"

	"github.com/go75/gen/common"
	"github.com/go75/gen/utils"
)

type ControllerEngine struct {
	*Config
}

func New(config *Config) *ControllerEngine {
	return &ControllerEngine{
		Config: config,
	}
}

func (e *ControllerEngine) Run() error {
	err := e.gen()
	if err != nil {
		return err
	}

	return nil
}

func (e *ControllerEngine) gen() error {

	var err error

	if !e.DisableRes {
		err = e.genRes()
		if err != nil {
			fmt.Println("gen controller res error: ", err.Error())
			return fmt.Errorf("gen controller res error: %s", err)
		}
	}

	for structName, fields := range common.Parser.Structs {
		err = e.genCRUD(structName, fields)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *ControllerEngine) genRes() error {
	resContent := fmt.Sprintf(`
package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	OK = iota
	ERR
)

type Res struct {
	Code int %s
	Msg string %s
	Data any %s
}

func Ok(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &Res{
		Code: OK,
		Msg: "success",
		Data: data,
	})
}

func Err(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &Res{
		Code: ERR,
		Msg: "failed",
	})
}
`, 
	"`json:\"code\"`", "`json:\"msg\"`", "`json:\"data\"`")

	err := utils.WriteGoFile(e.SavePath + "/res/res.go", resContent)
	if err != nil {
		return err
	}

	return nil
}

func (e *ControllerEngine) genCRUD(structName string, fields []*common.Field) error {
	content := new(strings.Builder)
	content.WriteString("package ")
	content.WriteString(e.PackageName)
	content.WriteString("\n\n")

	if !(e.DisableC && e.DisableR && e.DisableU && e.DisableD) {
		content.WriteString("import (\n\t")
		content.WriteString(`"github.com/gin-gonic/gin"`)
		content.WriteString("\n)\n\n")
	}

	if !e.DisableC {
		content.WriteString("func Create")
		content.WriteString(structName)
		content.WriteString("(c *gin.Context) {\n\n}\n\n")
	}

	if !e.DisableR {
		content.WriteString("func Query")
		content.WriteString(structName)
		content.WriteString("(c *gin.Context) {\n\n}\n\n")
	}

	if !e.DisableU {
		content.WriteString("func Update")
		content.WriteString(structName)
		content.WriteString("(c *gin.Context) {\n\n}\n\n")
	}

	if !e.DisableD {
		content.WriteString("func Delete")
		content.WriteString(structName)
		content.WriteString("(c *gin.Context) {\n\n}\n\n")
	}

	controllerPath := e.SavePath + "/" +fields[0].TableName + ".go"
	err := utils.WriteGoFile(controllerPath, content.String())
	if err != nil {
		fmt.Println("gen crup error: ", err.Error())
		return err
	}

	return nil
}