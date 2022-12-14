package task

import (
	"fmt"
	"time"
)

type PortStats struct {
	Name      string
	Frequency string
}

func (ps *PortStats) setName(name string) {
	ps.Name = name
}

func (ps *PortStats) getFrequency() string {
	return ps.Frequency
}

func (ps *PortStats) getName() string {
	return ps.Name
}

func (ps PortStats) Execute() (suceess bool, error bool) {
	fmt.Println("Port stats - Started")
	time.Sleep(2 * time.Second)

	return true, false
}

func (ps *PortStats) onComplete() {
	fmt.Println("Port Stats Task Completed")
}

func (ps *PortStats) onError() {
	fmt.Println("Error in Port Stats Task")
}
