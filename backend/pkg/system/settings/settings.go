package settings

type Settings struct {
	BatteryMonitorStatus BatteryMonitorStatus `json:"battery_monitor"`
}

func NewSettings() *Settings {
	// Poptulate the settings with default values.
	s := &Settings{
		BatteryMonitorStatus: BatteryMonitorStatus{
			Enabled:        new(bool),
			MaxChargeLimit: new(uint8),
			MinChargeLimit: new(uint8),
		},
	}

	*s.BatteryMonitorStatus.Enabled = false
	*s.BatteryMonitorStatus.MaxChargeLimit = 80
	*s.BatteryMonitorStatus.MinChargeLimit = 30

	return s
}
