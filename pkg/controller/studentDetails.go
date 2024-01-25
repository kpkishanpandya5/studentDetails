package controller

import (
	"fmt"
	"time"
)

const DateLayout = "02-01-2006"

type StudentData struct {
	Name       string
	Age        int
	Dob        time.Time
	Married    bool
	Phone      map[string]int
	Email      string
	SpouseName string
	Subject    []string
	CourseFee  func(subject []string) int
}

func StudentDetails() (map[string]StudentData, error) {
	var err error
	studentDetails := make(map[string]StudentData)
	kishanDetails := StudentData{Name: "Kishan",
		Phone: map[string]int{"mobile": 7597402772,
			"home":      8866305865,
			"emergency": 9414349046},
		Email:      "kpkishanpandya5@gmail.com",
		Married:    true,
		SpouseName: "Deepthi",
		Subject:    []string{"golang", "react", "mysql", "kubernetes/docker"},
		CourseFee: func(subject []string) int {
			return getFee(subject)
		}}
	kishanDetails.Dob, kishanDetails.Age, err = getAge("17-10-1992")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	studentDetails["Kishan Details"] = kishanDetails

	divyDetails := StudentData{Name: "Divy",
		Phone: map[string]int{"mobile": 8765432109,
			"home":      9876543210,
			"emergency": 7654321098},
		Email:      "divy5@gmail.com",
		Married:    false,
		SpouseName: "",
		Subject:    []string{"golang", "react", "mysql", "kubernetes/docker"},
		CourseFee: func(subject []string) int {
			return getFee(subject)
		}}
	divyDetails.Dob, divyDetails.Age, err = getAge("17-08-2001")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	studentDetails["Divy Details"] = divyDetails
	return studentDetails, nil
}

func getAge(dobString string) (time.Time, int, error) {
	dob, err := time.Parse(DateLayout, dobString)
	age := int(time.Now().Sub(dob).Hours()) / 24 / 365

	return dob, age, err
}

func getFee(courses []string) int {
	if len(courses) > 3 {
		return 15000
	} else if len(courses) < 1 {
		return 0
	} else {
		fee := 0

		for _, course := range courses {
			switch course {
			case "react":
				fee = fee + 5000
			case "golang":
				fee = fee + 5000
			default:
				fee = fee + 2500

			}
		}
		return fee
	}
}
