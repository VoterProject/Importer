package importer

import "github.com/voterproject/importer/pkg/parser/pa"

func Start(path string) {
	pa_parser.ParseDirectory(path)
}
