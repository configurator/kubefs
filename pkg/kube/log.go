package kube

import (
	"fmt"
	"log"
)

func logDouble(before string, success string, err string) func(string, ...interface{}) func(error) {
	return func(what string, params ...interface{}) func(error) {
		log.Printf(fmt.Sprintf(before, what), params...)
		return func(err error) {
			if err != nil {
				log.Printf("Error loading "+what+": %#v\n", append(params, err)...)
			} else {
				log.Printf("Successfully loaded "+what+"\n", params...)
			}
		}
	}
}

// Logs "loading xxx" when entering a function and returns a function that logs on exit
// Use defer to invoke the return value with the error result:
// func LoadSomething(param int) (err error) {
//     defer LogLoading("something(%d)", param)(err)
// }
var LogLoading = logDouble(
	"Loading %s...\n",
	"Successfully loaded %s\n",
	"Error loading %s\n",
)

var LogDeleting = logDouble(
	"Deleting %s...\n",
	"Successfully deleted %s\n",
	"Error deleting %s\n",
)
