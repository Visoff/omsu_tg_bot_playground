package eservice

import (
	"errors"
	"time"
)

func gen_only_time(loc *time.Location) func(int, int) time.Time {
    return func (hours int, minutes int) time.Time {
        return time.Date(0, 0, 0, hours, minutes, 0, 0, loc)
    }
}

func TimeByClass(class int) (time.Time, time.Time, error) {
    only_time := gen_only_time(time.Now().Location())
    switch class {
        case 1:
            return only_time(8, 45), only_time(10, 20), nil
        case 2:
            return only_time(10, 30), only_time(12, 5), nil
        case 3:
            return only_time(12, 45), only_time(14, 20), nil
        case 4:
            return only_time(14, 30), only_time(16, 5), nil
        case 5:
            return only_time(16, 15), only_time(17, 50), nil
        case 6:
            return only_time(18, 0), only_time(19, 35), nil
        case 7:
            return only_time(19, 45), only_time(21, 20), nil
        case 8:
            return only_time(21, 30), only_time(13, 5), nil
    }
    return time.Time{}, time.Time{}, errors.New("Class does not exist")
}
