package internal

import (
	"fmt"

	"github.com/npvu1510/en-vocab-server/internal/migration"
	"github.com/npvu1510/en-vocab-server/pkg/config"
	"github.com/npvu1510/en-vocab-server/pkg/utils"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(lc fx.Lifecycle, conf *config.Config) *gorm.DB {
	// CONNECT DATABASE
	host := conf.Postgres.Host
	port := conf.Postgres.Port
	user := conf.Postgres.User
	password := conf.Postgres.Password
	dbname := conf.Postgres.DbName

	connectStr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	utils.PanicError(err)

	// MIGRATE DATABASE
	migration.Migrations(db)

	fmt.Println("✅ Connected to database")

	return db
}
