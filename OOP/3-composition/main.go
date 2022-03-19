//----------------Composition Instead of Inheritance-------------

// Composition by embedding structs


package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

type blogPost struct {
	title   string
	content string
	author  // embeded struct type author
}

func (a author) fullName() string {
	//return fmt.Sprintf("%s %s", a.firstName, a.lastName)
	return a.firstName + " " + a.lastName
}

func (b blogPost) details() {
	fmt.Println("Title :\t\t", b.title)
	fmt.Println("Content :\t", b.content)
	fmt.Println("Author :\t", b.fullName())
	fmt.Println("Bio :\t\t", b.bio)
}

func main() {

	blog := blogPost{
		title:   "Inheritance in Go",
		content: "Go supports composition instead of inheritance",
		author: author{
			firstName: "Sandy",
			lastName:  "wilson",
			bio:       "Golang Enthusiast",
		},
	}

	blog.details()
}

/*
	Title :          Inheritance in Go
	Content :        Go supports composition instead of inheritance
	Author :         Sandy wilson
	Bio :            Golang Enthusiast

*/
