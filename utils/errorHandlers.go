package utils

import "log"

// HandleTemplateErr handles error executing a template
func HandleTemplateErr(err error) {
	if err != nil {
		log.Fatalln("error executing template:", err)
	}
}

// HandleFileErr handles error opening a file
func HandleFileErr(err error) {
	if err != nil {
		log.Fatalln("error opening file:", err)
	}
}
