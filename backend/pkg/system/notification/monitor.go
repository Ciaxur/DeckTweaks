package notification

import (
	"fmt"
	"time"

	"steamdeckhomebrew.decktweaks/pkg/devtools/icon"
	"steamdeckhomebrew.decktweaks/pkg/devtools/toaster"
	"steamdeckhomebrew.decktweaks/pkg/system/battery"
	"steamdeckhomebrew.decktweaks/pkg/system/settings"
)

const (
	LOW_BATTERY_PERCENTAGE = 5
)

type NotificationMonitor struct {
	isAlive    bool
	pollRate_s uint8
	m_settings *settings.Settings
}

func NewMonitor(s *settings.Settings) *NotificationMonitor {
	return &NotificationMonitor{
		isAlive:    false,
		pollRate_s: 60,
		m_settings: s,
	}
}

func (n *NotificationMonitor) Stop() {
	n.isAlive = false
}

func (n *NotificationMonitor) Start() {
	if !n.isAlive {
		n.isAlive = true
		go n.RunMonitor(n.m_settings)
	}
}

// Maps the given battery percentage value to an apporpriate battery icon.
// Returns the icon matching the appropriate battery percentage.
func getBatteryChargeIcon(percentage int64) string {
	if percentage > 90 {
		return icon.FaBatteryFull
	} else if percentage > 50 {
		return icon.FaBatteryThreeQuarters
	} else if percentage > 40 {
		return icon.FaBatteryHalf
	} else if percentage > 15 {
		return icon.FaBatteryQuarter
	} else {
		return icon.FaBatteryEmpty
	}
}

// RunMonitor - Monitors the current state of the SteamDeck's battery,
// notifying the user based on the given settings.
func (n *NotificationMonitor) RunMonitor(s *settings.Settings) {
	fmt.Printf("[%v] NotificationMonitor: Starting...\n", time.Now())
	n.isAlive = true

	// Notification invokation state.
	notifiedOfMaxLimit, notifiedOfMinLimit, notifiedLowBattery := false, false, false

	// Use a poll rate to watch over files.
	for n.isAlive {
		time.Sleep(time.Duration(n.pollRate_s) * time.Second)
		fmt.Printf("[%v] NotificationMonitor: Polling at %ds...\n", time.Now(), n.pollRate_s)

		// Monitor the battery's state compared to the settings.
		bat_perc, err := battery.GetPercentage()
		if err != nil {
			fmt.Printf("[%v] NotificationMonitor: Failed to read battery percentage: %v\n", time.Now(), err)
		}

		// Snapshot the current state.
		maxChargeLimit := int64(*s.BatteryMonitorStatus.MaxChargeLimit)
		minChargeLimit := int64(*s.BatteryMonitorStatus.MinChargeLimit)

		// Notify the user of the battery charge reaching max limit.
		if bat_perc >= maxChargeLimit && !notifiedOfMaxLimit {
			notifiedOfMaxLimit = true
			toast := toaster.NewToast("Battery Monitor", fmt.Sprintf("Battery charge reached max charge limit of %d%%.", maxChargeLimit))
			toast.SetIcon(getBatteryChargeIcon(bat_perc))
			if err := toast.Toast(); err != nil {
				fmt.Printf("[%v] NotificationMonitor: Failed invoke toast: %v\n", time.Now(), err)
			}
		}

		// Notify the user of the battery charge reaching min limit.
		if bat_perc <= minChargeLimit && !notifiedOfMinLimit {
			notifiedOfMinLimit = true
			toast := toaster.NewToast("Battery Monitor", fmt.Sprintf("Battery charge reached min charge limit of %d%%.", minChargeLimit))
			toast.SetIcon(getBatteryChargeIcon(bat_perc))
			if err := toast.Toast(); err != nil {
				fmt.Printf("[%v] NotificationMonitor: Failed invoke toast: %v\n", time.Now(), err)
			}
		}

		// Notify the user when the battery is at extremely low charge.
		if bat_perc <= LOW_BATTERY_PERCENTAGE && !notifiedLowBattery {
			notifiedLowBattery = true
			toast := toaster.NewToast("Low Battery!", fmt.Sprintf("Charge your battery asap! Charge reached %d%%.", bat_perc))
			toast.SetIcon(getBatteryChargeIcon(bat_perc))
			if err := toast.Toast(); err != nil {
				fmt.Printf("[%v] NotificationMonitor: Failed invoke toast: %v\n", time.Now(), err)
			}
		}

		// Reset notification state.
		if bat_perc > minChargeLimit {
			notifiedOfMinLimit = false
		}
		if bat_perc < maxChargeLimit {
			notifiedOfMaxLimit = false
		}
		if bat_perc > LOW_BATTERY_PERCENTAGE {
			notifiedLowBattery = false
		}
	}

	fmt.Printf("[%v] NotificationMonitor: Exiting...\n", time.Now())
}
