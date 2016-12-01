// Leap stub file
package leap

const testVersion = 2

// IsLeadYear determines if a year is a leap year.
func IsLeapYear(y int) bool {
	switch {
	// Every year divisible by 400 is a leap year
	case y%400 == 0:
		return true
	// Every year divisible by 100 is NOT a leap year
	case y%100 == 0:
		return false
	// Every year divisible by 4 (and not 100) is a leap year
	case y%4 == 0:
		return true
	}
	// Otherwise not a leap year
	return false
}
