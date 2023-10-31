package inout

import (
	"fmt"
	emp "app/empdata"
	sc "strconv"
)

func Print_Startup() {
	fmt.Println(".  .  .  .  .  .  .  .  .  .  .")
	fmt.Println("Welcome To GoSalary v1.0")
	fmt.Println("By: Rendell Christian J. Ngo")
	fmt.Println("ADPRG MP1 2023")
	fmt.Println(".  .  .  .  .  .  .  .  .  .  .\n")
	fmt.Println("\nNote: For text input, use underscores (_) instead of spaces.\n")
}

func Print_End_Screen() {
	fmt.Println(".  .  .  .  .  .  .  .  .  .  .")
	fmt.Println("Thank you for using the app.")
	fmt.Println("ADPRG MP1 2023")
	fmt.Println(".  .  .  .  .  .  .  .  .  .  .\n")
}

func Print_Menu() {
	fmt.Println("MAIN MENU\n")
	fmt.Println("[1] Compute Salary")
	fmt.Println("[2] New Config")
	fmt.Println("[3] Exit App\n")
}

func Print_SC_Menu(configs [4]emp.Config, availability emp.Bitfield) {
	fmt.Println("SELECT CONFIG\n")
	for i, cfg := range configs {
		var display string
		if availability & (emp.Bitfield) (0b1 << i) != 0 {
			display = cfg.Name
		} else {
			continue
		}
		fmt.Printf("[%d] %v\n", i + 1, display)
	}

	fmt.Println()
}

func Print_New_Cfg(configs [4]emp.Config, availability emp.Bitfield) {
	fmt.Println("SELECT CONFIG SLOT\n")
	for i, cfg := range configs {
		var display string
		var ind string
		if availability & (emp.Bitfield) (0b1 << i) != 0 {
			display = cfg.Name
			ind = "/"
		} else {
			display = "- Create New -"
			ind = sc.Itoa(i + 1)
		}
		fmt.Printf("[%v] %v\n", ind, display)
	}

	fmt.Println()
}

func Offer_Edit_Config(cfg emp.Config) {
	fmt.Println("EDIT CONFIG MODE\n")
	fmt.Printf("[1] Name\t\t%v\n", cfg.Name)
	fmt.Printf("[2] Daily Rate\t\t%.2f\n", cfg.Daily_Rate)
	fmt.Printf("[3] Regular Hours\t%v\n", cfg.Max_Reg_Hours)
	fmt.Printf("[4] Day Types\t\t%v\n", "[Select to View]")
	fmt.Printf("[5] In Time\t\t%v\n", emp.MT_To_String(cfg.In_Time))
	fmt.Printf("[6] Out Time\t\t%v\n\n", emp.MT_To_String(cfg.Out_Time))
	fmt.Println("[7] Save Changes")
	fmt.Println("[8] Discard Changes\n")
}

func Offer_Edit_Day_Types(day_types []emp.Bitfield) {
	fmt.Println("EDIT DAYS MODE\n")
	fmt.Printf("[1] Monday\t\t%v\n", emp.Create_Day_Type_String(day_types[0]))
	fmt.Printf("[2] Tuesday\t\t%v\n", emp.Create_Day_Type_String(day_types[1]))
	fmt.Printf("[3] Wednesday\t\t%v\n", emp.Create_Day_Type_String(day_types[2]))
	fmt.Printf("[4] Thursday\t\t%v\n", emp.Create_Day_Type_String(day_types[3]))
	fmt.Printf("[5] Friday\t\t%v\n", emp.Create_Day_Type_String(day_types[4]))
	fmt.Printf("[6] Saturday\t\t%v\n", emp.Create_Day_Type_String(day_types[5]))
	fmt.Printf("[7] Sunday\t\t%v\n\n", emp.Create_Day_Type_String(day_types[6]))

	fmt.Println("[8] Go Back\n")
}

func Offer_Edit_Employee(em emp.Employee, mode int) {
	fmt.Println("EDIT EMPLOYEE MODE\n")
	if mode == 1 {
		fmt.Printf("[0] Mode: %v", "OUT_TIME\n")
	} else {
		fmt.Printf("[0] Mode: %v", "IN_TIME\n")
	}
	fmt.Println("\t\t\tIN\tOUT")
	fmt.Printf("[1] Monday\t\t%v\t%v\n", emp.MT_To_String(em.Week[0].In_Time), emp.MT_To_String(em.Week[0].Out_Time))
	fmt.Printf("[2] Tuesday\t\t%v\t%v\n", emp.MT_To_String(em.Week[1].In_Time), emp.MT_To_String(em.Week[1].Out_Time))
	fmt.Printf("[3] Wednesday\t\t%v\t%v\n", emp.MT_To_String(em.Week[2].In_Time), emp.MT_To_String(em.Week[2].Out_Time))
	fmt.Printf("[4] Thursday\t\t%v\t%v\n", emp.MT_To_String(em.Week[3].In_Time), emp.MT_To_String(em.Week[3].Out_Time))
	fmt.Printf("[5] Friday\t\t%v\t%v\n", emp.MT_To_String(em.Week[4].In_Time), emp.MT_To_String(em.Week[4].Out_Time))
	fmt.Printf("[6] Saturday\t\t%v\t%v\n", emp.MT_To_String(em.Week[5].In_Time), emp.MT_To_String(em.Week[5].Out_Time))
	fmt.Printf("[7] Sunday\t\t%v\t%v\n\n", emp.MT_To_String(em.Week[6].In_Time), emp.MT_To_String(em.Week[6].Out_Time))

	fmt.Println("[8] Start Computation")
	fmt.Println("[9] Cancel\n")
}

func Print_Report(report [7]string) {
	for _, s := range report {
		fmt.Println(s)
	}
}