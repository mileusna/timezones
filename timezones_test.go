package timezones_test

import (
	"fmt"
	"testing"

	"github.com/mileusna/timezones"
)

func TestTimezones(t *testing.T) {
	tzones := timezones.List()
	if len(tzones) == 0 {
		t.Fatal("No zones loaded")
	}

	for _, tz := range tzones {
		fmt.Println(tz)
	}
}
