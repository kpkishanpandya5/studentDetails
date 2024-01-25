package detailsPrint

import (
	"fmt"
	"pkg/main/controller"
)

func ShowStudentDetails(studentDetails map[string]controller.StudentData) {

	for _, details := range studentDetails {
		fmt.Println("Hello", details.Name)
		fmt.Println("Your age is", details.Age)
		fmt.Println("You were born on", details.Dob.Format(controller.DateLayout))
		fmt.Println("Married!!! {true or false}-", details.Married)
		for key, value := range details.Phone {
			fmt.Println("Phone Number", key, value)
		}
		fmt.Println("Email Id", details.Email)
		if details.Married == true {
			fmt.Println("Spouse Name", details.SpouseName)
		} else if details.Age >= 21 {
			fmt.Println("You are eligible to marry, why don't you marry")
		}
		fmt.Println("subjects you are going to read", details.Subject)
		fmt.Println("Fee you need to pay", details.CourseFee(details.Subject))

		fmt.Println()
		fmt.Println()
	}

}
