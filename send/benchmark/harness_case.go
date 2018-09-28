package send

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type caseDefinition struct {
	name               string
	bench              benchCase
	count              int
	numOps             int
	size               int
	requiredIterations int
	runtime            time.Duration

	cumulativeRuntime time.Duration
	elapsed           time.Duration
	startAt           time.Time
	isRunning         bool
}

// TimerManager is a subset of the testing.B tool, used to manage
// setup code.
type TimerManager interface {
	ResetTimer()
	StartTimer()
	StopTimer()
}

func (c *caseDefinition) ResetTimer() {
	c.startAt = time.Now()
	c.elapsed = 0
	c.isRunning = true
}

func (c *caseDefinition) StartTimer() {
	c.startAt = time.Now()
	c.isRunning = true
}

func (c *caseDefinition) StopTimer() {
	if !c.isRunning {
		return
	}
	c.elapsed += time.Since(c.startAt)
	c.isRunning = false
}

func (c *caseDefinition) roundedRuntime() time.Duration {
	return roundDurationMS(c.runtime)
}

func (c *caseDefinition) Run(ctx context.Context) *benchResult {
	out := &benchResult{
		dataSize:   c.size * c.numOps * c.count,
		name:       c.name,
		operations: c.numOps * c.count,
	}
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, 2*executionTimeout)
	defer cancel()

	fmt.Println("=== RUN", out.name)
	if c.requiredIterations == 0 {
		c.requiredIterations = minIterations
	}

benchRepeat:
	for {
		if ctx.Err() != nil {
			break
		}
		if c.cumulativeRuntime >= c.runtime {
			if out.trials >= c.requiredIterations {
				break
			} else if c.cumulativeRuntime >= executionTimeout {
				break
			}
		} else if out.trials >= maxIterations {
			break
		}

		res := result{
			iterations: c.count,
		}

		start := time.Now()
		res.err = c.bench(ctx, c, c.count, c.size, c.numOps)
		realTestLength := time.Since(start)

		res.duration = c.elapsed
		c.cumulativeRuntime += realTestLength

		switch res.err {
		case context.DeadlineExceeded:
			break benchRepeat
		case context.Canceled:
			break benchRepeat
		case nil:
			out.trials++
			c.elapsed = 0
			out.raw = append(out.raw, res)
		default:
			continue
		}

	}

	out.duration = out.totalDuration()
	fmt.Printf("    --- REPORT: count=%d trials=%d requiredtrials=%d runtime=%s\n",
		c.count, out.trials, c.requiredIterations, c.runtime)
	if out.hasErrors() {
		fmt.Printf("    --- ERRORS: %s\n", strings.Join(out.errReport(), "\n       "))
		fmt.Printf("--- FAIL: %s (%s)\n", out.name, out.roundedRuntime())
	} else {
		fmt.Printf("--- PASS: %s (%s)\n", out.name, out.roundedRuntime())
	}

	return out

}

func (c *caseDefinition) String() string {
	return fmt.Sprintf("name=%s, count=%d, runtime=%s timeout=%s, maxiters=%d",
		c.name, c.count, c.runtime, executionTimeout, maxIterations)
}

func getProjectRoot() string { return filepath.Dir(filepath.Dir(getDirectoryOfFile())) }

func getDirectoryOfFile() string {
	_, file, _, _ := runtime.Caller(1)

	return filepath.Dir(file)
}
