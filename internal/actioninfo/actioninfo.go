package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string)(err error)
	ActionInfo()(string, error)
}

func Info(dataset []string, dp DataParser) {

	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Println("Ошибка: ", err )
			continue
		}
	}
dp.ActionInfo()
}
