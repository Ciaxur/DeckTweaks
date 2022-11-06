package battery

import (
	"fmt"
	"math"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
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

type BatteryState struct {
	Error      error
	CurrentNow int64
	VoltageNow int64
	Status     string
	Capacity   int64
}

/**
 * Stream data from the battery sysfs files into the given state channel.
 *
 * @param battChan Go-routine channel which contains written data from sysfs.
 * @param quit Go-routine channel to sync exiting the go-routine.
 */
func StreamBatteryState(battChan chan BatteryState, quit chan bool) {
	// Initial battery state.
	state := BatteryState{
		Error:      nil,
		VoltageNow: 0,
	}

	// Helper function: Open file in Read-Only mode.
	openFdRdOnly := func(filePath string) (*os.File, error) {
		fd, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
		basename := path.Base(filePath)
		if err != nil {
			state.Error = fmt.Errorf("failed to open %s file: %v", basename, err)
			battChan <- state
			return nil, err
		}
		return fd, nil
	}

	// Helper function: Reading integer values from sysfs.
	buffer := make([]byte, 255)
	numRe := regexp.MustCompile(`\d+`)
	strRe := regexp.MustCompile(`\w+`)
	readStringFromFd := func(fd *os.File) (string, error) {
		_, err := fd.Read(buffer)
		if err != nil {
			state.Error = fmt.Errorf("failed to read data from system file: %v", err)
			battChan <- state
			return "", state.Error
		}

		// Clean up the data.
		str := string(strRe.Find(buffer))
		return str, nil
	}

	readInt64FromFd := func(fd *os.File) (int64, error) {
		strBuffer, err := readStringFromFd(fd)
		if err != nil {
			return 0, err
		}

		// Convert voltage value to an integer.
		strValue := string(numRe.FindString(strBuffer))
		i_val, err := strconv.ParseInt(strValue, 10, 64)
		if err != nil {
			state.Error = fmt.Errorf("failed to parse integer value from system file: %v", err)
			battChan <- state
			return 0, state.Error
		}

		return i_val, nil
	}

	// Open all file descriptors for the battery files.
	currentNow_fd, err := openFdRdOnly(BAT1_CURRENT_NOW_FILEPATH)
	if err != nil {
		fmt.Printf("[%s] Battery State Go-routine failure: %v", time.Now(), err)
		return
	}
	voltageNow_fd, err := openFdRdOnly(BAT1_VOLTAGE_NOW_FILEPATH)
	if err != nil {
		fmt.Printf("[%s] Battery State Go-routine failure: %v", time.Now(), err)
		return
	}
	chargeState_fd, err := openFdRdOnly(BAT1_STATUS_FILEPATH)
	if err != nil {
		fmt.Printf("[%s] Battery State Go-routine failure: %v", time.Now(), err)
		return
	}
	capacity_fd, err := openFdRdOnly(BAT1_CAPACITY_FILEPATH)
	if err != nil {
		fmt.Printf("[%s] Battery State Go-routine failure: %v", time.Now(), err)
		return
	}
	defer currentNow_fd.Close()
	defer voltageNow_fd.Close()
	defer chargeState_fd.Close()
	defer capacity_fd.Close()

	for {
		time.Sleep(POLL_RATE)

		// Reset file index.
		currentNow_fd.Seek(0, 0)
		voltageNow_fd.Seek(0, 0)
		chargeState_fd.Seek(0, 0)
		capacity_fd.Seek(0, 0)

		fmt.Println("DEBUG: Looping")
		switch {
		case <-quit:
			fmt.Printf("[%s] Battery State Go-routine: quitting", time.Now())
			return
		default:
			// Read values from sysfs.
			fmt.Println("DEBUG: Reading from buffer")
			state.CurrentNow, err = readInt64FromFd(currentNow_fd)
			if err != nil {
				fmt.Printf("[%s] Battery State Go-routine failed to read from Current Now sysfs\n", time.Now())
				return
			}
			state.VoltageNow, err = readInt64FromFd(voltageNow_fd)
			if err != nil {
				fmt.Printf("[%s] Battery State Go-routine failed to read from Voltage Now sysfs\n", time.Now())
				return
			}
			state.Status, err = readStringFromFd(chargeState_fd)
			if err != nil {
				fmt.Printf("[%s] Battery State Go-routine failed to read from Charge State sysfs\n", time.Now())
				return
			}
			state.Capacity, err = readInt64FromFd(capacity_fd)
			if err != nil {
				fmt.Printf("[%s] Battery State Go-routine failed to read from Battery Capacity sysfs\n", time.Now())
				return
			}

			// Write the data to the channel.
			battChan <- state
		}
	}
}
