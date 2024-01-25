package main

import (
	"fmt"
	"pkg/main/controller"
	"pkg/main/detailsPrint"
)

func main() {
	studentDetails, err := controller.StudentDetails()
	if err != nil {
		fmt.Println(err)
		return
	}
	detailsPrint.ShowStudentDetails(studentDetails)
}
