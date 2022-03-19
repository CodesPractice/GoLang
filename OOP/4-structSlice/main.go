//------------- Embedding slice of structs----------

package main

import "fmt"

type website struct {
	blogPost []blogPost
}
type author struct {
	firstName string
	lastName  string
	bio       string
}

type blogPost struct {
	title   string
	content string
	author  // embeded struct type author -anonymous field
}

func (w website) content() {
	for _, v := range w.blogPost {
		v.details()
		fmt.Println()
	}
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

	author1 := author{
		firstName: "Sandy",
		lastName:  "wilson",
		bio:       "Golang Enthusiast",
	}

	author2 := author{
		firstName: "Bob",
		lastName:  "Joe",
		bio:       "Golang Enthusiast",
	}

	blog1 := blogPost{
		title:   "Inheritance in Go",
		content: "Go supports composition instead of inheritance",
		author:  author1,
	}

	blog2 := blogPost{
		title:   "Struct instead of Classes in Go",
		content: "Go does not support classes but methods can be added to structs",
		author:  author1,
	}

	blog3 := blogPost{
		title:   "Concurrency",
		content: "Go is a concurrent language and not a parallel one",
		author:  author2,
	}

	w := website{
		blogPost: []blogPost{blog1, blog2, blog3},
	}
	w.content()

	/*
		Title :          Inheritance in Go
		Content :        Go supports composition instead of inheritance
		Author :         Sandy wilson
		Bio :            Golang Enthusiast

		Title :          Struct instead of Classes in Go
		Content :        Go does not support classes but methods can be added to structs
		Author :         Sandy wilson
		Bio :            Golang Enthusiast

		Title :          Concurrency
		Content :        Go is a concurrent language and not a parallel one
		Author :         Bob Joe
		Bio :            Golang Enthusias
	*/
}
