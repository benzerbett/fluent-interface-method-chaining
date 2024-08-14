package main

import "fmt"

// enum for job group
type JobGroup int

const (
	Engineer JobGroup = iota
	Manager  JobGroup = iota
	Founder  JobGroup = iota
	Designer JobGroup = iota
)

type Employee struct {
	name     string
	age      int
	jobgroup JobGroup
}

func (e *Employee) SetName(name string) *Employee {
	e.name = name
	return e
}

func (e *Employee) SetAge(age int) *Employee {
	e.age = age
	return e
}

func (e *Employee) SetJobGroup(jobgroup JobGroup) *Employee {
	e.jobgroup = jobgroup
	return e
}

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) GetAge() int {
	return e.age
}

func (e *Employee) GetJobGroup() JobGroup {
	return e.jobgroup
}

func main() {
	e := new(Employee)
	e.SetName("Amanda").SetAge(28).SetJobGroup(Founder)

	fmt.Printf("Name: %s\n", e.GetName())
	fmt.Printf("Age: %d\n", e.GetAge())
	fmt.Printf("JobGroup: %d\n", e.jobgroup)
}
