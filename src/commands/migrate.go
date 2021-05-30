package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"src/models"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Short: "Sends migrations to configured database",
	Long: "Sends migrations that are defined inside migrate.go file to the configured database",
	Run: migrate,
}


func migrate(cmd *cobra.Command, args []string) {
	fmt.Println("Opening database connection")
	db := models.GetConnection()

	fmt.Println("Migrating model: Endpoint")
	if err := db.AutoMigrate(&models.EndPoint{}); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Closing database connection")
	models.CloseConnection(db)
}