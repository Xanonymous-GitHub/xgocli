package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // add to prevent error in build.
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

var DBTypeToStructType = map[string]string{
	"int":       "int32",
	"tinyint":   "int8",
	"smallint":  "int",
	"mediumint": "int64",
	"bigint":    "int64",
	"bit":       "int",
	"bool":      "bool",
	"enum":      "string",
	"set":       "string",
	"varchar":   "string",
}

// use db-info to create a new model.
func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

// connect to database.
func (m *DBModel) Connect() error {
	var err error

	// connection script.
	s := "%s:%s@tcp(%s)/information_schema?" + "charset=%s&parseTime=True&loc=Local"

	// fill value to the connection script.
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)

	// start the connection to database use the script.
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

// get columns from the specific table in specific database.
func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	// query string. '?' is a place which will be filled something.
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?"

	// send query string to database.
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("no data")
	}
	defer func() { _ = rows.Close() }()

	// fill the query result(rows) to 'columns'
	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn

		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}
