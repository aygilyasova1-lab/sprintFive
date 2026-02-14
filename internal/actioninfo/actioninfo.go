package actioninfo

import (
	"log"
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
		}	
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("Ошибка: ", err)
		}
		fmt.Println(info)
	}

