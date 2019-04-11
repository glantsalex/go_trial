package main

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"io/ioutil"
	"log"
	"innovi-event-server/model"
)

func getInnoviEventFromFile( fullFileName string ) ( ev *model.InnoviEvent ) {

	content, err := ioutil.ReadFile( fullFileName  )
	if err != nil {
		log.Printf("Error reading from file %s/\n%s\nDefault dataset will be used.", fullFileName , err )
		return
	}

	ev = &model.InnoviEvent{}
	if err = json.Unmarshal( content, ev );err != nil {
		seelog.Errorf("Error unmarshal event from file:\n%s", err )
		ev =nil
	}
	return
}
