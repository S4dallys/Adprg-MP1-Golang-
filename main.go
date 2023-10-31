package main

import (
	io "app/inout"
	emp "app/empdata"
)

func main() {
	// collections
	var CONFIGS = [4]emp.Config{emp.Create_Default_Config()}
	var CFG_AVAILABILITY emp.Bitfield = 1

	io.Print_Startup()

	running := true

	for running {
		io.Print_Menu()

		var choice int
		io.Get_Int(&choice)

		switch choice {
			case 1: 
				io.Print_SC_Menu(CONFIGS, CFG_AVAILABILITY)

				err := io.Get_Int(&choice)
				valid := io.Already_Exists(choice, CFG_AVAILABILITY)

				var cur_cfg emp.Config

				if valid && err == nil {
					choice-- // 1 based
					cur_cfg = CONFIGS[choice]
					cur_emp := emp.Create_Default_Employee(cur_cfg)

					pass := io.Run_Employee_Editor(&cur_emp)

					if pass {
						report := emp.Generate_Week_Report(cur_cfg, cur_emp)
						io.Print_Report(report)
					} 
				} else if err == nil {
					io.Print_Error(io.INVALID_INPUT_ERROR)
				}
			case 2:
				io.Print_New_Cfg(CONFIGS, CFG_AVAILABILITY)

				err := io.Get_Int(&choice)
				valid := !io.Already_Exists(choice, CFG_AVAILABILITY)

				if valid && err == nil {
					choice-- // 1 based
					io.Run_Config_Editor(&CONFIGS[choice], &CFG_AVAILABILITY, choice) 
				} else if err == nil {
					io.Print_Error(io.ALREADY_EXISTS_ERROR)
				}
			case 3:
				running = false
			default:
				io.Print_Error(io.INVALID_INPUT_ERROR)
		}
	}

	io.Print_End_Screen()
}