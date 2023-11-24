package empdata

import (
	sc "strconv"
	"fmt"
)

const (
	S_NORMAL	  float64 = 1.000
	M_NS 			      = 1.100
	M_NORM_OT			  = 1.250
	M_OT			 	  = 1.300
	S_REST 				  = 1.300
	S_SNW				  = 1.300
	S_SNW_REST		      = 1.500
	S_REGH			      = 2.000
	S_REGH_REST		      = 2.600
)

// returns reg, ot, reg_ns, ot_ns
func Calculate_Hours_Worked(in_hour int, out_hour int, max_reg_hours int) (int, int, int) {
	max_reg_hours++; // account for break hour

	if in_hour == out_hour {
		return 0, 0, 0 // absent
	}

	reg_hours := min(max_reg_hours, Hour_Diff(in_hour, out_hour))
	ot_hours := max(0, Hour_Diff(in_hour, out_hour) - max_reg_hours)

	ot_start := (in_hour + reg_hours) % 24

	reg_looped_around := in_hour > (in_hour + reg_hours) % 24
	ot_looped_around := ot_start > (ot_start + ot_hours) % 24

	var reg_ns int = 0
	var ot_ns int = 0

	if reg_looped_around {
		reg_ns += min(2, 24 - in_hour)
		reg_ns += min(6, ot_start)
	} else {
		reg_ns = max(0, ot_start - 22) + max(0, 6 - in_hour)
	}


	if ot_looped_around && ot_hours > 0 {
		ot_ns += min(2, 24 - ot_start)
		ot_ns += min(6, out_hour)
	} else if !ot_looped_around && ot_hours > 0 {
		ot_ns = max(0, out_hour - 22) + max(0, 6 - ot_start)
	}

	return ot_hours - ot_ns, reg_ns, ot_ns
}


func Generate_Day_Report(day_type Bitfield, day Day, daily_rate float64, max_reg_hours int) (string, float64) {
	in_hour := day.In_Time.Hour
	out_hour := day.Out_Time.Hour

	// account for minutes (2:01 will be less than an hour so get its ceiling)
	// unless end time compensates for it
	if in_hour == out_hour && day.Out_Time.Minute < day.In_Time.Minute {
		in_hour = (in_hour + 1) % 24
	} else if in_hour != out_hour && day.Out_Time.Minute + (60 - day.In_Time.Minute) < 60 {
		in_hour = (in_hour + 1) % 24
	}

	ot_hours, reg_ns_hours, ot_ns_hours := Calculate_Hours_Worked(in_hour, out_hour, max_reg_hours)

	hourly_rate := daily_rate / float64(max_reg_hours)

	is_rest, is_snw, is_regh := Get_Day_Type(day_type)

	var salary float64 = 0
	var m_base float64 = 1.0

	if is_regh {
		if is_rest {
			m_base = S_REGH_REST
		} else {
			m_base = S_REGH
		}
	} else if is_snw {
		if is_rest {
			m_base = S_SNW_REST
		} else {
			m_base = S_SNW
		}
	} else {
		if is_rest {
			m_base = S_REST
		} else {
			m_base = S_NORMAL
		}
	}


	if in_hour != out_hour {
		salary = m_base * float64(daily_rate) + hourly_rate * 
				m_base * float64(reg_ns_hours) * M_NS

		if day_type == D_NORMAL {
			salary += hourly_rate * (
				m_base * float64(ot_hours) * M_NORM_OT + 
				m_base * float64(ot_ns_hours) * M_NORM_OT * M_NS)
		} else {
			salary += hourly_rate * (
				m_base * float64(ot_hours) * M_OT +
				m_base * float64(ot_ns_hours) * M_OT * M_NS)
		}
	} else {
		if (day_type != D_NORMAL) {
			salary = daily_rate;
		}
	}

	// returning a "report"
	// could maybe move this to another file

	format := "Daily Rate\t\t%15v\n" + 
			 "IN Time\t\t\t%15v\n" + 
			 "OUT Time\t\t%15v\n" + 
			 "Day Type\t\t%15v\n" + 
			 "Night Shift Hrs\t\t%15v\n" +
			 "OT Regular Hrs\t\t%15v\n" + 
			 "OT NS Hrs\t\t%15v\n" + 
			 "Salary\t\t\t%15v\n"


	report := fmt.Sprintf(format, sc.FormatFloat(daily_rate, 'f', 2, 64), MT_To_String(day.In_Time), MT_To_String(day.Out_Time), 
		Create_Day_Type_String(day_type), sc.Itoa(reg_ns_hours), sc.Itoa(ot_hours), sc.Itoa(ot_ns_hours), 
		sc.FormatFloat(salary, 'f', 2, 64))

	return report, salary
}

func Generate_Week_Report(config Config, employee Employee) [7]string {
	var week_report [7]string
	day_names := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	var total_sal float64

	for i := 0; i < 7; i++ {
		rep, sal := Generate_Day_Report(config.Day_Types[i], employee.Week[i], config.Daily_Rate, config.Max_Reg_Hours)
		week_report[i] = "--- " + day_names[i] + " ---\n" + rep
		total_sal += sal
	}

	week_report[6] += "\n--- Total Salary: " + sc.FormatFloat(total_sal, 'f', 2, 64) + " ---\n"

	return week_report
}