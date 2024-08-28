package main

const (
	speedLimitMin = -3
	speedLimitMax = 3
)

type speed struct {
	x, y int
}

// GetSpeeds generates all posibile new speeds for the current speed
func GetSpeeds(currentSpeed speed) []speed {
	speeds := make([]speed, 0, 9)

	for xInc := -1; xInc <= 1; xInc++ {
		xx := currentSpeed.x + xInc
		if xx < speedLimitMin || xx > speedLimitMax {
			continue
		}

		for yInc := -1; yInc <= 1; yInc++ {
			yy := currentSpeed.y + yInc
			if yy < speedLimitMin || yy > speedLimitMax {
				continue
			}

			if yy == 0 && xx == 0 {
				// skip zero speed because no hop happens
				continue
			}

			speeds = append(speeds, speed{xx, yy})
		}
	}

	return speeds
}
