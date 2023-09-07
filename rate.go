package interest

import (
	"fmt"
	"math"
)

type Rate struct {
	value float64
	days  float64
}

func NewRate(value float64, days int) (Rate, error) {
	if days <= 0 {
		return Rate{}, fmt.Errorf("days must be greater than zero")
	}
	r := Rate{}
	r.days = float64(days)
	r.value = value
	return r, nil
}

func (r Rate) Value() float64 {
	return r.value
}

func (r Rate) Days() int {
	return int(r.days)
}

// Resample returns a new effective rate por the provided period.
func (r Rate) Resample(days int) (Rate, error) {
	if r.days == 0 {
		return Rate{}, fmt.Errorf("rate wasn't initialized correctly")
	}
	if days <= 0 {
		return Rate{}, fmt.Errorf("days must be greater than zero")
	}

	out := Rate{days: float64(days)}
	out.value = math.Pow(1+r.value, out.days/r.days) - 1

	return out, nil
}

// It's very common to see interest rates expressed Nominal
// (monthly effective rate) x 12
func (r Rate) NominalYearly() (float64, error) {
	if r.days == 0 {
		return 0, fmt.Errorf("rate wasn't initialized correctly")
	}

	monthly, err := r.Resample(30)
	if err != nil {
		return 0, fmt.Errorf("resampling to 30 days")
	}

	return monthly.value * 12, nil
}
