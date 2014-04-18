package memory

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

// by default, send a memory report to stdout, or to whatever io.Writer you pass in
func Report( w... io.Writer ) {
	if w == nil {
		w = append(make([]io.Writer,0),os.Stderr)
	}
   m := &runtime.MemStats{}
   runtime.ReadMemStats(m)
	for _,out := range w {
		fmt.Fprintf(out, "Memory Report: Alloc(%v), TotalAlloc(%v)\n", m.Alloc, m.TotalAlloc)
	}
}
