package settings

import "fmt"

type BatteryMonitorStatus struct {
	Enabled        *bool  `json:"enabled"`
	MaxChargeLimit *uint8 `json:"max_charge_limit"`
	MinChargeLimit *uint8 `json:"min_charge_limit"`
}

func (s *BatteryMonitorStatus) SetMaxChargeLimit(newMaxLimit uint8) error {
	// Restrict having the max limit above the minimum charge limit.
	if newMaxLimit <= *s.MinChargeLimit {
		return fmt.Errorf("maximum charge limit must be larger than the minimum charge limit of %d", s.MinChargeLimit)
	}

	// Restring to having a max value of 100.
	if newMaxLimit > 100 {
		return fmt.Errorf("max charge limit, %d, cannot surpass 100%%", newMaxLimit)
	}

	*s.MaxChargeLimit = newMaxLimit
	return nil
}

func (s *BatteryMonitorStatus) SetMinChargeLimit(newMinLimit uint8) error {
	// Restrict having the min limit less than the maximum charge limit.
	if newMinLimit >= *s.MaxChargeLimit {
		return fmt.Errorf("minimum charge limit must be less than the maximum charge limit of %d", s.MaxChargeLimit)
	}

	// Restring to having a max value of 100.
	if newMinLimit > 100 {
		return fmt.Errorf("min charge limit, %d, cannot surpass 100%%", newMinLimit)
	}

	*s.MinChargeLimit = newMinLimit
	return nil
}
