package main

import (
	"fmt"
	"time"
)

func studentDetails() (map[string]studentData, error) {
	var err error
	studentDetails := make(map[string]studentData)
	kishanDetails := studentData{name: "Kishan",
		phone: map[string]int{"mobile": 7597402772,
			"home":      8866305865,
			"emergency": 9414349046},
		email:      "kpkishanpandya5@gmail.com",
		married:    true,
		spouseName: "Deepthi",
		subject:    []string{"golang", "react", "mysql", "kubernetes/docker"},
		courseFee: func(subject []string) int {
			return getFee(subject)
		}}
	kishanDetails.dob, kishanDetails.age, err = getAge("17-10-1992")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	studentDetails["Kishan Details"] = kishanDetails

	divyDetails := studentData{name: "Divy",
		phone: map[string]int{"mobile": 8765432109,
			"home":      9876543210,
			"emergency": 7654321098},
		email:      "divy5@gmail.com",
		married:    false,
		spouseName: "",
		subject:    []string{"golang", "react", "mysql", "kubernetes/docker"},
		courseFee: func(subject []string) int {
			return getFee(subject)
		}}
	divyDetails.dob, divyDetails.age, err = getAge("17-08-2001")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	studentDetails["Divy Details"] = divyDetails
	return studentDetails, nil
}

func getAge(dobString string) (time.Time, int, error) {
	dob, err := time.Parse(dateLayout, dobString)
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
