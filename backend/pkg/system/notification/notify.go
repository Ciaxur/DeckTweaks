package notification

type NotificationState struct {
	// Notification state when charge limit reached, so that we don't spam the
	// user.
	MaxChargeLimitNotfied bool
	MinChargeLimitNotfied bool
}
