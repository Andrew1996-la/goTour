package time

import (
	"fmt"
	"time"
)

func calculateInterval() {
	firstVersion := time.Date(2023, time.February, 1, 0, 0, 0, 0, time.UTC)
	secondVersion := time.Date(2023, time.August, 8, 0, 0, 0, 0, time.UTC)

	interval := secondVersion.Sub(firstVersion)

	days := int(interval.Hours() / 24)
	fmt.Printf("Между выходами версий прошло %d дней", days)
}
