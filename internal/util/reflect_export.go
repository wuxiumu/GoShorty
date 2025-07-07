// internal/util/reflect_export.go
package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"

	"goshor/internal/core"
)

// ExportCSV 将链接数据导出为 CSV 文件
func ExportCSV(data map[string]*core.Link, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	t := reflect.TypeOf(core.Link{})
	headers := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		headers[i] = t.Field(i).Name
	}
	_ = w.Write(headers)

	for _, v := range data {
		r := reflect.ValueOf(*v)
		row := make([]string, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			row[i] = fmt.Sprintf("%v", r.Field(i).Interface())
		}
		_ = w.Write(row)
	}
	return nil
}
