package empdata

func Create_Default_Employee(config Config) Employee {
	var week [7]Day

	for i := 0; i < 7; i++ {
		week[i] = Day{config.In_Time, config.Out_Time}
	}

	return Employee{"Mr_Default", week}
}

func Create_Default_Config() Config {
	return Config{
		"Default_Cfg",
		500,
		8,
		[7]Bitfield{D_NORMAL, D_NORMAL, D_NORMAL, D_NORMAL, D_NORMAL, D_REST, D_REST}, // 5 weekdays, 2 rest days
		Time{9, 0},
		Time{9, 0},
	}
}