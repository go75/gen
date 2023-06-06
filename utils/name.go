package utils

import "strings"

func CamelName(rawName string) string {
	var camelName string
	for _, p := range strings.Split(rawName, "_") {
		// 字段首字母大写的同时, 是否要把其他字母转换为小写
		switch len(p) {
		case 0:
		case 1:
			camelName += strings.ToUpper(p[0:1])
		default:
			camelName += strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return camelName
}