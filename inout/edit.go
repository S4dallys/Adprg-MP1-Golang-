package inout

import (
	emp "app/empdata"
	sc "strconv"
)

// negate for opposite value
func Already_Exists(choice int, availability emp.Bitfield) bool {
	return ((emp.Bitfield) (0b1) << (choice - 1) & availability != 0)
}

// returns true if edits saved
func Run_Config_Editor(cfg * emp.Config, availability * emp.Bitfield, index int) {
	new_cfg := emp.Create_Default_Config()
	new_cfg.Name = "Cfg_" + sc.Itoa(index + 1)

	editing := true

	for editing {
		var choice int
		Offer_Edit_Config(new_cfg)
		Get_Int(&choice)

		switch choice {
			case 1:
				var temp = new_cfg.Name // in case fail
				Get_String(&temp)
				new_cfg.Name = temp
			case 2:
				var temp = new_cfg.Daily_Rate 
				Get_Float(&temp)
				new_cfg.Daily_Rate = temp
			case 3:
				var temp = new_cfg.Max_Reg_Hours
				Get_Int(&temp)
				new_cfg.Max_Reg_Hours = temp
			case 4:
				Run_Day_Type_Editor(new_cfg.Day_Types[:])
			case 5:
				var temp string
				Get_String(&temp)
				new_time, err := emp.String_To_MT(temp)
				if err == nil {
					new_cfg.In_Time = new_time
				} else {
					Print_Error(INVALID_INPUT_ERROR)
				}
			case 6:
				var temp string
				Get_String(&temp)
				new_time, err := emp.String_To_MT(temp)
				if err == nil {
					new_cfg.Out_Time = new_time
				} else {
					Print_Error(INVALID_INPUT_ERROR)
				}
			case 7:
				// if saved
				*cfg = new_cfg
				*availability = *availability | (emp.Bitfield) (0b1 << index)
				editing = false
			case 8:
				editing = false
			default:
				Print_Error(INVALID_INPUT_ERROR)
		}
	}
}

func Run_Day_Type_Editor(day_types []emp.Bitfield) {
	editing := true

	for editing {
		var choice int
		Offer_Edit_Day_Types(day_types)
		err := Get_Int(&choice)

		if choice == 8 {
			editing = false
		} else if choice > 0 && choice < 8 && err == nil {
			cur := &day_types[choice - 1]
			*cur += 1
			*cur %= 6
		} else {
			Print_Error(INVALID_INPUT_ERROR)
		}
	}
}

func Run_Employee_Editor(em* emp.Employee) bool {
	editing := true

	var input_mode = 1

	for editing {
		var choice int
		Offer_Edit_Employee(*em, input_mode)
		err := Get_Int(&choice)

		if err != nil {
			Print_Error(INVALID_INPUT_ERROR)
			continue
		}

		if choice == 0 {
			input_mode = 1 - input_mode
		} else if choice == 8 {
			return true
		} else if choice == 9 {
			return false
		} else if choice > 0 && choice < 8 {
			// edit
			var temp string
			Get_String(&temp)
			new_time, err := emp.String_To_MT(temp)
			if err == nil {
				if input_mode == 1 {
					em.Week[choice - 1].Out_Time = new_time
				} else {
					em.Week[choice - 1].In_Time = new_time
				}
			} else {
				Print_Error(INVALID_INPUT_ERROR)
			}
		} else {
			Print_Error(INVALID_INPUT_ERROR)
		}
	}

	return false
}
	