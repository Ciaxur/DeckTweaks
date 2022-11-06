package battery

import "time"

const (
	BAT1_CLASS_BASEPATH = "/sys/class/power_supply/BAT1"
	// TODO: /sys/class/hwmon/hwmon5/maximum_battery_charge_rate
	BAT1_CURRENT_NOW_FILEPATH = "/sys/class/power_supply/BAT1/current_now"
	BAT1_VOLTAGE_NOW_FILEPATH = "/sys/class/power_supply/BAT1/voltage_now"
	BAT1_STATUS_FILEPATH      = "/sys/class/power_supply/BAT1/status"
	BAT1_CAPACITY_FILEPATH    = "/sys/class/power_supply/BAT1/capacity"
)

// Battery status Enum.
const (
	STATUS_UNKNOWN     uint8 = 0
	STATUS_DISCHARGING uint8 = 1
	STATUS_CHARGING    uint8 = 2
)

// Default values.
const (
	POLL_RATE = 250 * time.Millisecond
)
