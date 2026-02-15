package actioninfo

import (
	"log"
	"fmt"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
if len(dataset) == 0 {
	return
}
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Println("Ошибка: ", err )
			continue
		}	
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("Ошибка: ", err)
			continue
		}
		fmt.Println(info)
	}
	}

