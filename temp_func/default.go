/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/4/18 0:47
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/4/18 0:47
 */

package temp_func

import (
	"strings"
	"text/template"
)

var DefaultFuncMap = template.FuncMap{
	"LeftUpper": LeftUpper,
	"LeftLower": LeftLower,
	"ToUpper": func(s string) string {
		return strings.ToUpper(s)
	},
	"ToLower": func(s string) string {
		return strings.ToLower(s)
	},
	"ToSnake": ToSnake,
	"ToCamel": ToCamel,
	"ToPath": func(s string) string {
		return strings.ReplaceAll(ToSnake(s), "_", "-")
	},
}
