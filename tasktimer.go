package tasktimer

/*
TaskTimer provide a simple interface to trace your task time.

For example,

	import "github.com/c9s/tasktimer"

	t := tasktimer.Trace("DB Query")
	// do your query stuff here.
	t.End()

Or you can pass a callback.

	tasktimer.TraceFunc("DB Query", func() {
		// do something
	})

*/

import (
	"log"
	"time"
)

type TaskTimer struct {
	TaskName  string
	StartTime time.Time
	EndTime   time.Time
}

func (self *TaskTimer) End() {
	elapsed := time.Since(self.StartTime)
	log.Printf("Trace end: %s, elapsed %f msecs\n", self.TaskName, float32(elapsed.Nanoseconds()/1000/1000))
}

func Trace(s string) *TaskTimer {
	return &TaskTimer{TaskName: s, StartTime: time.Now()}
}

/*
TraceFunc([task name] , [callback])
*/
func TraceFunc(s string, cb func()) {
	var t = Trace(s)
	cb()
	t.End()
}
