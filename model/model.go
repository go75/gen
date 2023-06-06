package model

import (
	"database/sql"
	"errors"
	"github.com/go75/gen/common"
	"github.com/go75/gen/utils"
)

type ModelEngine struct {
	*Config
}

func New(config *Config) *ModelEngine {
	return &ModelEngine{
		Config: config,
	}
}

func (e *ModelEngine) Run() error {
	var err error

	// 获取数据连接
	e.db, err = sql.Open("mysql", common.Parser.DSN)
	if err != nil {
		return errors.New("dial mysql failed: " + err.Error())
	}

	defer e.db.Close()

	// 获取数据库中过滤后的所有表
	common.Parser.Structs, err = e.getTables()
	if err != nil {
		return err
	}

	// 创建model层文件
	err = e.createStructs()
	if err != nil {
		return err
	}

	return nil
}

func (e *ModelEngine) createStructs() (err error) {

	for structName, tableColumn := range common.Parser.Structs {
		if e.FiltHook(structName) {
			continue
		}

		err = e.createModel(structName, tableColumn)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *ModelEngine) getTables() (map[string][]*common.Field, error) {
	sqlStr := `SELECT COLUMN_NAME,DATA_TYPE,IS_NULLABLE,TABLE_NAME,COLUMN_COMMENT FROM information_schema.COLUMNS 
	WHERE table_schema = DATABASE() order by TABLE_NAME asc, ORDINAL_POSITION asc`

	rows, err := e.db.Query(sqlStr)
	if err != nil {
		return nil, errors.New("Failed to read table information: " + err.Error())
	}

	defer rows.Close()

	tables := make(map[string][]*common.Field)

	i := 0

	for rows.Next() {
		i++
		col := common.Field{}
		err = rows.Scan(&col.Name, &col.Type, &col.Nullable, &col.TableName, &col.Comment)

		if err != nil {
			return nil, errors.New("Failed to read table struct: " + err.Error())
		}

		structName := utils.CamelName(col.TableName)
		if _, ok := tables[structName]; !ok {
			tables[structName] = []*common.Field{&col}
		} else {
			tables[structName] = append(tables[structName], &col)
		}
	}

	return tables, nil
}

func (e *ModelEngine) createModel(structName string, tableColumns []*common.Field) error {

	content := common.NewContent(e.Location)

	// 组装struct
	content.WriteString("type ")
	content.WriteString(structName)
	content.WriteString(" struct {\n")

	for i := 0; i < len(tableColumns); i++ {
		col := tableColumns[i]
		camelName := utils.CamelName(col.Name)
		content.WriteString(camelName)
		content.WriteByte(' ')
		content.WriteString(utils.TypeMap[col.Type])

		if !e.DisableJsonTag {
			content.WriteString(" `json:\"")
			if e.JsonTagToUpper {
				content.WriteString(camelName)
			} else {
				content.WriteString(col.Name)
			}
			content.WriteString("\"`")
		}
		
		if !e.DisableComment {
			if col.Comment != "" {
				content.WriteString(" //")
				content.WriteString(col.Comment)
			}
		}
		
		content.WriteByte('\n')
	}

	content.WriteByte('}')

	filepath := e.SavePath + "/" + tableColumns[0].TableName + ".go"
	
	err := utils.WriteGoFile(filepath, content.String())
	return err
}
