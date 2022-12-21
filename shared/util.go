package shared

import "fmt"

func CompleteWithZero(value int64) string {
	return fmt.Sprintf("%06d", value)
}
