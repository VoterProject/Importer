package importer

import (
	"github.com/voterproject/importer/pkg/config"
	"github.com/voterproject/importer/pkg/parser/wa"
	"github.com/voterproject/importer/pkg/sql"
)

func Start(configPath string) {
	c := config.ConfigFromFile(configPath)
	db := sql.NewSQL(c.DSN)
	//pa_parser.ParseDirectory(c.PA, db)
	wa_parser.ParseDirectory(c.WA, db)
}
