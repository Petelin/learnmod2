package learnmod2

import (
    "time"

    "golang.org/x/time/rate"
)

func x()  {
    _ = rate.Every(time.Hour)
}
