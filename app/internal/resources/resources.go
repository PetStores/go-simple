package resources

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

type R struct {
	Config  Config
	PetFile *csv.Writer
	fd      *os.File
	DB      *reform.DB
	conn    *sql.DB
}

type Config struct {
	DiagPort    int    `envconfig:"DIAG_PORT" default:"8081" required:"true"`
	RESTAPIPort int    `envconfig:"PORT" default:"8080" required:"true"`
	DBURL       string `envconfig:"DB_URL" default:"postgres://user:password@localhost:5432/petstore?sslmode=disable" required:"true"`
	PetFileName string `envconfig:"PET_FILE" default:"~/petfile.csv"`
}

func New(logger *zap.SugaredLogger) (*R, error) {
	conf := Config{}
	err := envconfig.Process("", &conf)
	if err != nil {
		return nil, fmt.Errorf("can't process the config: %w", err)
	}

	conn, err := sql.Open("pgx", conf.DBURL)
	if err != nil {
		return nil, err
	}

	db := reform.NewDB(conn, postgresql.Dialect, reform.NewPrintfLogger(logger.Infof))

	fd, _ := os.Open(conf.PetFileName)

	return &R{
		Config:  conf,
		DB:      db,
		PetFile: csv.NewWriter(fd),
	}, nil
}

func (r *R) Release() error {
	r.fd.Close()
	return r.conn.Close()
}
