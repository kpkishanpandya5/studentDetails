package main

import "fmt"

func showStudentDetails(studentDetails map[string]studentData) {

	for _, details := range studentDetails {
		fmt.Println("Hello", details.name)
		fmt.Println("Your age is", details.age)
		fmt.Println("You were born on", details.dob.Format(dateLayout))
		fmt.Println("Married!!! {true or false}-", details.married)
		for key, value := range details.phone {
			fmt.Println("Phone Number", key, value)
		}
		fmt.Println("Email Id", details.email)
		if details.married == true {
			fmt.Println("Spouse Name", details.spouseName)
		} else if details.age >= 21 {
			fmt.Println("You are eligible to marry, why don't you marry")
		}
		fmt.Println("subjects you are going to read", details.subject)
		fmt.Println("Fee you need to pay", details.courseFee(details.subject))

		fmt.Println()
		fmt.Println()
	}

}
