package application

import (
	"os"

	"github.com/joho/godotenv"
)

/*Initialize configuration*/
func (goservice *GoService) Initialize() {
	godotenv.Load()

	var appconfig ConfigParameters = handleConfiguration()
	goservice.Config = &appconfig
}

func handleConfiguration() ConfigParameters {
	var appconfig ConfigParameters

	appconfig.DbUser = os.Getenv("db_user")
	appconfig.DbPass = os.Getenv("db_pass")
	appconfig.DbName = os.Getenv("db_name")
	appconfig.DbHost = os.Getenv("db_host")
	appconfig.DbPort = os.Getenv("db_port")
	appconfig.DbType = os.Getenv("db_type")
	return appconfig
}
