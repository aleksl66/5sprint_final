package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("некорректный формат данных: ожидается 2 элемента, получено %d", len(parts))
	}

	// TrimSpace не используется, чтобы пробелы вызывали ошибку
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка преобразования количества шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным, получено %d", steps)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка преобразования длительности: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", fmt.Errorf("количество шагов должно быть положительным")
	}
	if ds.Duration <= 0 {
		return "", fmt.Errorf("продолжительность должна быть положительной")
	}

	dist := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка расчёта калорий: %w", err)
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		dist,
		calories,
	), nil
}
