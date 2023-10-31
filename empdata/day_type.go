package empdata

import (
	sc "strconv"
	"fmt"
	"errors"
)

type Bitfield uint8

const (
	D_NORMAL Bitfield = 0
	D_REST   		  = 1
	D_SNW 		      = 2
	D_REGH 		      = 4
)

func MT_To_String(time Time) string {
	return fmt.Sprintf("%02d%02d", time.Hour, time.Minute)
}

func String_To_MT(time string) (Time, error) {
	int_time, err := sc.Atoi(time)

	hr := int_time / 100
	min := int_time % 100

	if err != nil || hr < 0 || hr > 23 || min < 0 || min > 59 {
		return Time{}, errors.New("Invalid Time.")
	} else {
		return Time{hr, min}, nil
	}
}

// returns is_rest, is_snw, is_regh
func Get_Day_Type(day_type Bitfield) (bool, bool, bool) {
	return day_type & D_REST != 0, 
	day_type & D_SNW != 0, 
	day_type & D_REGH != 0
}

func True_Mod(i int, n int) int {
	return ((i % n) + n) % n
}

func Hour_Diff(start int, end int) int {
	return True_Mod(end - start, 24)
}

func Create_Day_Type_String (day_type Bitfield) string {
	var day_type_fmt string 
	is_rest, is_snw, is_regh := Get_Day_Type(day_type)

	if day_type == D_NORMAL {
		day_type_fmt = "Normal Day"
	} else if is_snw {
		day_type_fmt = "SNWH"
	} else if is_regh {
		day_type_fmt = "RH"
	}

	if day_type == D_REST {
		day_type_fmt = "Rest Day"
	} else if is_rest {
		day_type_fmt += ", Rest Day"
	}

	return day_type_fmt
}