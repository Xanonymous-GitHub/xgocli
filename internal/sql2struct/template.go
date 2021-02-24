package sql2struct

import (
	``
	"fmt"
	"os"
	"text/template"
	"xgocli/internal/word"
)

// define the struct template as a string.
const structTpl = `type {{ .TableName | ToCamelCase }} struct {
		{{ range .Columns }} {{ $length := len .Comment }} {{ if gt $length 0 }}//
		{{ .Comment }} {{ else }}// {{ .Name }} {{ end }}
		{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }} {{ .Name | ToCamelCase }}
		{{ .Type }} {{ .Tag }} {{ else }} {{ .Name }} {{ end }}
	{{ end }}
	
	func (model {{ .TableName | ToCamelCase }}) TableName() string {
		return "{{ .TableName }}"
	}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

// get the golang struct by the template.
func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))

	for _, column := range tbColumns {
		// json struct tag
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name: column.ColumnName,
			// convert db-type to golang type.
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}

// generate the golang struct from template and print to stdout.
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	// create a text-template-consumer (convert the field to real text)
	tpl := template.Must(template.New("sql2-struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	// combine the tplColumns and its name.
	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}

	// print to screen.
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}

	return nil
}
