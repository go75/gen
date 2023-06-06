package service

import (
	"fmt"
	"strings"

	"github.com/go75/gen/common"
	"github.com/go75/gen/utils"
)

type ServiceEngine struct {
	*Config
}

func New(config *Config) *ServiceEngine {
	return &ServiceEngine{
		Config: config,
	}
}

func (e *ServiceEngine) Run() error {

	err := e.gen()
	if err != nil {
		return err
	}

	return nil
}

func (e *ServiceEngine) gen() error {
	var err error
	for structName, fields := range common.Parser.Structs {
		if !e.DisableVO {
			voContent := e.genVo(structName, fields)
			voFilePath := e.SavePath + "/vo/" + fields[0].TableName + ".go"
			err = utils.WriteGoFile(voFilePath, voContent)
			if err != nil {
				fmt.Printf("service vo file %s write error: %s\n", voFilePath, err.Error())
				return err
			}
		}

		if !e.DisableTO {
			toContent := e.genTo(structName, fields)
			toFilePath := e.SavePath + "/to/" + fields[0].TableName + ".go"
			err = utils.WriteGoFile(toFilePath, toContent)
			if err != nil {
				fmt.Printf("service to file %s write error: %s\n", toFilePath, err.Error())
				return err
			}
		}

		serviceContent := e.genService(structName)
		servicePath := e.SavePath + "/" + fields[0].TableName + ".go"
		err = utils.WriteGoFile(servicePath, serviceContent)
		if err != nil {
			fmt.Printf("service file %s write error: %s\n", servicePath, err.Error())
			return err
		}
	}

	return nil
}

func (e *ServiceEngine) genVo(structName string, fields []*common.Field) string {
	content := e.genStruct("VO", structName, fields)
	return content
}

func (e *ServiceEngine) genTo(structName string, fields []*common.Field) string {
	content := e.genStruct("TO", structName, fields)
	return content
}

func (e *ServiceEngine) genStruct(postfix, structName string, fields []*common.Field) string {

	return fmt.Sprintf(`
package %s

type %s%s struct {
	%s}
`,
		strings.ToLower(postfix), structName, postfix, e.genFields(fields))
}

func (e *ServiceEngine) genService(structName string) string {
	return fmt.Sprintf(`
package %s

import (
	"%s"
	"%s"
)

func %sService(vo *vo.%sVO) *%sServiceHelper {
	return &%sServiceHelper{vo: vo}
}

type %sServiceHelper struct {
	vo *vo.%sVO
	to *to.%sTO
}

func (u *%sServiceHelper) Do%s() (*to.%sTO, error) {

	return u.to, nil
}

`,
		e.PackageName, common.Parser.ProjectPrefix+"/"+e.PackageName+"/vo", common.Parser.ProjectPrefix+"/"+e.PackageName+"/to",
		structName, structName, structName, structName, structName,
		structName, structName, structName, e.MethodName, structName)
}

func (e *ServiceEngine) genFields(fields []*common.Field) string {
	content := new(strings.Builder)

	for i := 0; i < len(fields); i++ {
		col := fields[i]
		camelName := utils.CamelName(col.Name)
		content.WriteString(camelName)
		content.WriteByte(' ')
		content.WriteString(utils.TypeMap[col.Type])

		tag := ""

		if !e.DisableJsonTag {
			tag += `json:"`
			if e.JsonTagToUpper {
				tag += camelName
			} else {
				tag += col.Name
			}
			tag += `"`
		}

		if !e.DisableFormTag {
			tag += ` form:"`
			if e.FormTagToUpper {
				tag += camelName
			} else {
				tag += col.Name
			}
			tag += `"`
		}

		if len(tag) > 0 {
			content.WriteString(" `")
			content.WriteString(tag)
			content.WriteByte('`')
		}

		if !e.DisableComment {
			if col.Comment != "" {
				content.WriteString(" //")
				content.WriteString(col.Comment)
			}
		}

		content.WriteByte('\n')
	}

	return content.String()
}
