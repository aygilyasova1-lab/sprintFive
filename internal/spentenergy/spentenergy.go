package spentenergy

import (
	"time"
	"errors"
)
var (
	invalidStepsError = errors.New("некорректное количество шагов")
	invalidHeightError = errors.New("некорректный рост")
	invalidWeightError = errors.New("некорректный вес")
	invalidDurationError = errors.New("некорректное время")
)
// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		return 0, invalidStepsError
	}
	if height <= 0 {
		return 0, invalidHeightError
	}
	if weight <= 0 {
		return 0, invalidWeightError
	}
	if duration <= 0 {
		return 0, invalidDurationError
	}
	
	meanSpeed := MeanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	result := (weight * meanSpeed * minutes) / minInH

	return result * walkingCaloriesCoefficient, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {

		return 0, invalidStepsError
	}
	if height <= 0 {

		return 0, invalidHeightError
	}
	if weight <= 0 {

		return 0, invalidWeightError
	}
	if duration <= 0 {

		return 0, invalidDurationError
	}
	meanSpeed := MeanSpeed(steps, height, duration)

	minutes := duration.Minutes()

	return (weight * meanSpeed * minutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {

	if height <= 0 {
		return 0
	}
	if steps <= 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)

	hours := duration.Hours()

	return distance / hours

}

func Distance(steps int, height float64) float64 {

	if steps <= 0 {
		return 0
	}
	if height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient

	return stepLength * float64(steps) / mInKm
}
