package trainings

import (
	"fmt"
	"strings"
	"errors"
	"strconv"
	"time"
	"log"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)
var (
	parsingError = errors.New("ошибка парсинга")
	conversionError = errors.New("ошибка преобразования типа")
	invalidTrainingType = errors.New("неизвестный тип тренировки")
)

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {

	dataSlice := strings.Split(datastring, ",")

	if len(dataSlice) != 3 {
		return parsingError
	}
	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return conversionError
	}
	t.Steps = steps

	t.TrainingType = dataSlice[1]

	duration, err := time.ParseDuration(dataSlice[2])
	if err != nil {
		return parsingError
	}
	t.Duration = duration

	return nil	
}

func (t Training) ActionInfo() (string, error) {

	distance := spentenergy.Distance(t.Steps, t.Height)

	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64

	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			log.Println("Ошибка: ", err)

			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			log.Println("Ошибка: ", err)

			return "", err
		}
	default:
		return "", invalidTrainingType
	}
	hours := t.Duration.Hours()

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %v ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, hours, 
	distance, meanSpeed, calories)

	return result, nil
}
