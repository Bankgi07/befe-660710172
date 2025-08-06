package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

// เช็คว่าเกียรตินิยมมั้ย
func (s *Student) IsHonor() bool {
	return s.GPA >= 3.50
}

// ตรวจสอบความถูกต้องของข้อมูล
func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Year < 1 || s.Year > 4 {
		return errors.New("year must be between 1-4")
	}
	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("gpa must be between 0-4")
	}
	return nil
}

func main() {
	// สร้าง slice ของ student
	students := []Student{
		{ID: "1", Name: "Thammatorn", Email: "BANKTHAMMATORN2547@GMAIL.COM", Year: 4, GPA: 3.75},
		{ID: "2", Name: "Alice", Email: "alice@gmail.com", Year: 4, GPA: 2.75},
	}

	// เพิ่มคนใหม่เข้า slice
	newStudent := Student{ID: "3", Name: "Trudy", Email: "trudy@gmail.com", Year: 4, GPA: 3.50}
	students = append(students, newStudent)

	// Loop แสดงผล
	for i, student := range students {
		fmt.Printf("%d. Honor: %v\n", i, student.IsHonor())
		fmt.Printf("%d. Validation: %v\n", i, student.Validate())
	}
}
