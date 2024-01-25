package main

import (
	"fmt"
	"time"
)

const dateLayout = "02-01-2006"

type studentData struct {
	name       string
	age        int
	dob        time.Time
	married    bool
	phone      map[string]int
	email      string
	spouseName string
	subject    []string
	courseFee  func(subject []string) int
}

func main() {
	studentDetails, err := studentDetails()
	if err != nil {
		fmt.Println(err)
		return
	}
	showStudentDetails(studentDetails)
}
