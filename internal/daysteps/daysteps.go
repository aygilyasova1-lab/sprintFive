package daysteps

import (
	"strings"
	"fmt"
	"strconv"
	"time"
	"log"
	"github.com/aygilyasova1-lab/sprintFive"
)

var (
	parsingError = errors.New("ошибка парсинга")
	conversionError = errors.New("ошибка преобразования типа")
)
type DaySteps struct {
	Steps int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {

	dataSlice := strings.Split(datastring, ",")

	if len(dataSlice) != 2 {
		log.Println("Ошибка: ", err)

		return parsingError
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		log.Println("Ошибка: ", err)

		return conversionError
	}
	duration, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		log.Println("Ошибка: ", err)

		return parsingError
	}
	ds.Steps = steps

	ds.Duration = duration

	return nil

}

func (ds DaySteps) ActionInfo() (string, error) {

	distance := Distance(ds.Steps, ds.Height)

	calories, err := WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		log.Println("Ошибка: ", err)
		
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)

	return result, nil

}
