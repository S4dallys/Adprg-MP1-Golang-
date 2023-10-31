package empdata

type Time struct {
	Hour		int
	Minute		int
}

type Day struct {
	In_Time		Time 	
	Out_Time	Time 	
}

type Employee struct {
	Name	    string
	Week	    [7]Day
}

// must be init before invoking builder functions
type Config struct {
	Name          	string
	Daily_Rate	  	float64
	Max_Reg_Hours 	int
	Day_Types		[7]Bitfield
	In_Time			Time
	Out_Time		Time
}