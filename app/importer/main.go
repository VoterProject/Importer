package importer

import (
	"github.com/voterproject/importer/pkg/parser/pa"
	"github.com/voterproject/importer/pkg/sql"
)

func Start(path, configPath string) {
	config := sql.ConfigFromFile(configPath)
	db := sql.NewSQL(config.DSN)
	pa_parser.ParseDirectory(path, db)
}
