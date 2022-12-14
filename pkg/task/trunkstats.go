package task

import (
	"fmt"
	"time"
)

type TrunkStats struct {
	Name      string
	Frequency string
}

func (ps *TrunkStats) setName(name string) {
	ps.Name = name
}

func (ps *TrunkStats) getFrequency() string {
	return ps.Frequency
}

func (ps *TrunkStats) getName() string {
	return ps.Name
}

func (ps TrunkStats) Execute() (suceess bool, error bool) {
	fmt.Println("Trunk stats - Started")
	time.Sleep(2 * time.Second)

	return true, false
}

func (ps *TrunkStats) onComplete() {
	fmt.Println("Trunk Stats Task Completed")
}

func (ps *TrunkStats) onError() {
	fmt.Println("Error in Trunk Stats Task")
}
