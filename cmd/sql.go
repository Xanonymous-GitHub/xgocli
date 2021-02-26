package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"xgocli/internal/sql2struct"
)

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql conversion and processes",
	Long:  "sql conversion and processes",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql transfer",
	Long:  "sql transfer",
	Run: func(cmd *cobra.Command, args []string) {
		// create a db info from user command.
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}

		// use the db info to create a db model.
		dbModel := sql2struct.NewDBModel(dbInfo)

		// connect to db.
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}

		// get columns.
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}

		// paint template.
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	// register 'struct' to 'sql' as its subcommand.
	sqlCmd.AddCommand(sql2structCmd)

	// add all flags to command that read user-input data to variables.
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "please provide db account name")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "please provide db account password")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "please provide db host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "please provide db charset")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "please provide which db you want to use")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "please provide your db name")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "please provide table name")
}
