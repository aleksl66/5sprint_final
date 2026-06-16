package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0.0 || height <= 0 || height <= 0.0 || duration <= 0 {
		return 0, errors.New("неверные данные")
	}
	speed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	baseCalories := (weight * speed * durationMinutes) / minInH
	calories := baseCalories * walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0.0 || height <= 0 || height <= 0.0 || duration <= 0 {
		return 0, errors.New("неверные данные")
	}
	speed := MeanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()
	calories := (weight * speed * durationMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}
	dist := Distance(steps, height) // расстояние в километрах
	hours := duration.Hours()
	if hours == 0 {
		return 0
	}
	return dist / hours
}

func Distance(steps int, height float64) float64 {
	stepLen := height * stepLengthCoefficient
	return float64(steps) * stepLen / mInKm
}
