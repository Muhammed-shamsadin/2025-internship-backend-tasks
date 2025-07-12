package main

import (
	"fmt"
)

func main() {
	var name string
	var total_subjects int
	var subject_name string
	var subject_mark int

	subjects_and_grades := make(map[string]string)

	grade := map[string]float64{
		"A+": 4.00,
		"A": 4.00,
		"A-": 3.75,
		"B+": 3.5,
		"B": 3.0,
		"B-": 2.75,
		"C+": 2.5,
		"C": 2.0,
		"C-": 1.75,
		"D": 1.0,
		"F": 0.0,
	}

	var total_grade []float64

	var average_grade float64

	fmt.Println("Hello, Welcome to Student Grade Calculator!")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - ")
	fmt.Println("Please enter your name:")
	fmt.Scan(&name)

	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - ")
	fmt.Println("How many subjects are you taking?")
	fmt.Scan(&total_subjects)

	for i:= 1; i <= total_subjects; i++ {
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - ")
		fmt.Println("Please enter the name of subject", i, ":")
		fmt.Scan(&subject_name)
		fmt.Println("Please enter the marks obtained in out of 100%", subject_name, ":")
		fmt.Scan(&subject_mark)

		if subject_mark >= 90 {
			subjects_and_grades[subject_name] = "A+"
			total_grade = append(total_grade, grade["A+"])
		} else if subject_mark >= 85 && subject_mark < 90 {
			subjects_and_grades[subject_name] = "A"
			total_grade = append(total_grade, grade["A"])
		} else if subject_mark >= 80 && subject_mark < 85 {
			subjects_and_grades[subject_name] = "A-"
			total_grade = append(total_grade, grade["A-"])
		} else if subject_mark >= 75 && subject_mark < 80 {
			subjects_and_grades[subject_name] = "B+"
			total_grade = append(total_grade, grade["B+"])
		} else if subject_mark >= 70 && subject_mark < 75 {
			subjects_and_grades[subject_name] = "B"
			total_grade = append(total_grade, grade["B"])
		} else if subject_mark >= 65 && subject_mark < 70 {
			subjects_and_grades[subject_name] = "B-"
			total_grade = append(total_grade, grade["B-"])
		} else if subject_mark >= 60 && subject_mark < 65 {
			subjects_and_grades[subject_name] = "C+"
			total_grade = append(total_grade, grade["C+"])
		} else if subject_mark >= 50 && subject_mark < 60 {
			subjects_and_grades[subject_name] = "C"
			total_grade = append(total_grade, grade["C"])
		} else if subject_mark >= 45 && subject_mark < 50 {
			subjects_and_grades[subject_name] = "C-"
			total_grade = append(total_grade, grade["C-"])
		} else if subject_mark >= 40 && subject_mark < 45 {
			subjects_and_grades[subject_name] = "D"
			total_grade = append(total_grade, grade["D"])
		} else {
			subjects_and_grades[subject_name] = "F"
			total_grade = append(total_grade, grade["F"])
		}
	}

	for _, num := range total_grade {
		average_grade += num
	}

	average_grade /= float64(len(total_grade))

	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - ")
	fmt.Println("Summary for ->", name)
	fmt.Println("Total Subjects:", total_subjects)
	fmt.Println("Subjects and Grades:")
	for subject, grade := range subjects_and_grades {
		fmt.Printf("  %s: %s\n", subject, grade)
	}
	//can u fix this line
	fmt.Printf("Average Grade: %.2f GPA\n", average_grade)



	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - ")
	var continueChoice string
	fmt.Println("Do you want to continue? (yes/no)")
	fmt.Scan(&continueChoice)

	if continueChoice == "yes" {
		main()
	} else {
		fmt.Println("Thank you for using the Student Grade Calculator!")
	}
}