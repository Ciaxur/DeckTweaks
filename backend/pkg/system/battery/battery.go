package battery

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/**
 * Helper function which attempts to parse an int64 value from the given file.
 *
 * @param filepath File path to read from.
 * @returns int64 Resulting Int64 value.
 * @returns error Error when a failure occurs during parsing or reading file.
 */
func read_int64_from_file(filepath string) (int64, error) {
	b_data, err := os.ReadFile(filepath)
	if err != nil {
		return 0, fmt.Errorf("failed to read file %s: %v", filepath, err)
	}

	// Parse the file's contents, expecting an integer value.
	s_data := strings.TrimSpace(string(b_data))
	i_val, err := strconv.ParseInt(s_data, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse int64 from %s with content '%s'", filepath, s_data)
	}

	return i_val, nil
}

/**
 * Obtain the current amps drawn by the battery.
 *
 * @returns float64 Resulting floating-point value.
 * @returns error Error when a failure occurs during parsing or reading file.
 */
func GetCurrentDraw() (float64, error) {
	i_val_int64, err := read_int64_from_file(BAT1_CURRENT_NOW_FILEPATH)
	if err != nil {
		return 0, err
	}

	// Convert the read in value to amps.
	i_val := float64(i_val_int64) * math.Pow10(-7)
	return i_val, nil
}

/**
 * Obtain the current voltage drawn by the battery.
 *
 * @returns float64 Resulting floating-point value.
 * @returns error Error when a failure occurs during parsing or reading file.
 */
func GetVoltageDraw() (float64, error) {
	uV_val_int64, err := read_int64_from_file(BAT1_VOLTAGE_NOW_FILEPATH)
	if err != nil {
		return 0, err
	}

	// Convert the read in value from uV -> V.
	v_val := float64(uV_val_int64) * math.Pow10(-6)
	return v_val, nil
}

/**
 * Obtain the current battery percentage.
 *
 * @returns int64 Current battery capacity.
 * @returns error Error when a failure occurs during parsing or reading file.
 */
func GetPercentage() (int64, error) {
	return read_int64_from_file(BAT1_CAPACITY_FILEPATH)
}

/**
 * Obtain the current charge status of the battery.
 *
 * @returns int8 Resulting ChargeStatus enum value.
 * @returns error Error when a failure occurs during parsing or reading file.
 */
func GetChargeStatus() (int8, error) {
	b_data, err := os.ReadFile(BAT1_STATUS_FILEPATH)
	if err != nil {
		return 0, fmt.Errorf("failed to read file %s", BAT1_STATUS_FILEPATH)
	}

	// Map the string value read with the appropriate enum.
	s_data := string(b_data)
	status_enum_val := STATUS_UNKNOWN

	switch s_data {
	case "Discharging":
		status_enum_val = STATUS_DISCHARGING
	case "Charging":
		status_enum_val = STATUS_CHARGING
	}

	return int8(status_enum_val), nil
}
