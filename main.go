package main

import (
	"fmt"

	"github.con/Ajinkya-Dhote/collector/pkg/task"
)

func main() {
	fmt.Println("Collector Started")

	colector := task.NewCollector("PortStasts")

	pst := &task.PortStats{Name: "PortStats", Frequency: "5s"}
	colector.Register(pst)

	tst := &task.TrunkStats{Name: "TrunkStats", Frequency: "5s"}
	colector.Register(tst)

	colector.Start()
}
