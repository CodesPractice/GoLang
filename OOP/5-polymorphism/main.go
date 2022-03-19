//-----------Polymorphism using an interface-------------

/*
	Any type which provides definition for all the methods of an interface
	is said to implicitly implement that interface.
*/

package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TmeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

type Advertisement struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (f FixedBilling) calculate() int {
	return f.biddedAmount
}

func (f FixedBilling) source() string {
	return f.projectName
}

func (tm TmeAndMaterial) calculate() int {
	return tm.hourlyRate * tm.noOfHours
}

func (tm TmeAndMaterial) source() string {
	return tm.projectName
}

func (a Advertisement) calculate() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {
	return a.adName
}

func calculateNetIncome(ic []Income) {
	var netIncome int = 0
	for _, income := range ic {
		fmt.Printf("From %s  Amount %d\n", income.source(), income.calculate())
		netIncome += income.calculate()
	}
	fmt.Printf("\nNet income of the organization %d", netIncome)
}

func main() {
	fb1 := FixedBilling{"Project 1", 5000}
	fb2 := FixedBilling{"Project 2", 8000}
	tm1 := TmeAndMaterial{"Project3", 160, 24}
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}

	inc := []Income{fb1, fb2, tm1, bannerAd, popupAd}
	calculateNetIncome(inc)
}

/*
From Project 1  Amount 5000
From Project 2  Amount 8000
From Project3  Amount 4000
From Banner Ad  Amount 1000
From Popup Ad  Amount 3750

Net income of the organization 17000
*/
