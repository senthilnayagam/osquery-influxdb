package main

import (
	"fmt"
	"os/exec"
)

func main() {

	fmt.Println("hello")

	sql := "\"SELECT pid, parent, name, resident_size, phys_footprint, user_time, system_time FROM processes;\""
	data := runQuery(sql)
	fmt.Println(data)

}

func runQuery(sql string) string {
	app := "osqueryi"
	//	sql =  +  sql
	fmt.Println(sql)
	//cmd := exec.Command(app," --json ", sql)
	cmd := app + " --json " + sql

	out, err := exec.Command("sh", "-c", cmd).Output()
	//out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return "error"
	}
	return string(out)

}
