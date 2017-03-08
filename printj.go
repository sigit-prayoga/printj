package printj

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	prefix = "PrintJ says: "
)

// Print prints any object into pretty json
func Print(v interface{}, pretty bool, label string) {
	if v == nil {
		log.Println(prefix, "The object you passed is nil.")
		return
	}
	var data []byte
	var err error
	if pretty {
		data, err = json.MarshalIndent(v, "", "    ")
	} else {
		data, err = json.Marshal(v)
	}

	if err != nil {
		log.Printf("%s: %v", prefix, err)
		return
	}

	fmt.Printf("** %s ** \n %s \n \n", label, string(data))
}
