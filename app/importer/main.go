package importer

import "github.com/voterproject/importer/pkg/parser/pa"

func Start() {
	pa_parser.ParseDirectory("/home/amir/Downloads/Voting/Data/")
}
