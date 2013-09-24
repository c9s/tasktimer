TaskTimer for Go
================

TaskTimer provide a simple interface to trace your task time.

For example,

```go
import "github.com/c9s/tasktimer"

t := tasktimer.Trace("DB Query")
// do your query stuff here.
t.End()
```

Or you can pass a callback.

```go
tasktimer.TraceFunc("DB Query", func() {
    // do something
})
```

