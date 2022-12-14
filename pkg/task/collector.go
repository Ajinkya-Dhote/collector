package task

import (
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

var (
	wg        sync.WaitGroup
	scheduler = gocron.NewScheduler(time.UTC)
)

type Collector struct {
	taskList []Task
	Name     string
}

func NewCollector(name string) *Collector {
	return &Collector{
		Name: name,
	}
}

func (c *Collector) Start() {
	c.start()
}

func (c *Collector) Register(t Task) {
	c.taskList = append(c.taskList, t)
}

func (c *Collector) Deregister(o Task) {
	c.taskList = removeFromslice(c.taskList, o)
}

func (c *Collector) start() {
	wg.Add(1)
	for _, task := range c.taskList {
		scheduleTask(task)
	}
	scheduler.StartAsync()
	wg.Wait()
}

func scheduleTask(task Task) {

	scheduler.Every(task.getFrequency()).Do(func() {
		s, e := task.Execute()
		if s {
			task.onComplete()
		}
		if e {
			task.onError()
		}
	})

}

func removeFromslice(taskList []Task, taskToRemove Task) []Task {
	taskListLength := len(taskList)
	for i, task := range taskList {
		if taskToRemove.getName() == task.getName() {
			taskList[taskListLength-1], taskList[i] = taskList[i], taskList[taskListLength-1]
			return taskList[:taskListLength-1]
		}
	}
	return taskList
}
