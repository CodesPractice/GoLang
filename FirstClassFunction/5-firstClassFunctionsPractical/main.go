package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}

	s3 := student{"Ann", "Marry", "B", "India"}

	s := []student{s1, s2, s3}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})

	c := filter(s, func(s student) bool {
		if s.country == "India" {
			return true
		}
		return false
	})

	fmt.Println("Grade B Students")
	fmt.Println(f)
	fmt.Println("Indian Students")
	fmt.Println(c)
}
