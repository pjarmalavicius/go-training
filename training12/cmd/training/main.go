package main

import (
	"flag"
	"log"
	"log/slog"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cpuCores := flag.Int("cpu", 1, "number of CPU cores to use")
	duration := flag.Duration("duration", 60*time.Second, "duration to run the load generator")

	flag.Parse()

	logger.Info("starting CPU and memory load generator",
		slog.Int("cpu_cores", *cpuCores),
		slog.String("duration", duration.String()),
	)

	runtime.GOMAXPROCS(*cpuCores)

	go func() {
		log.Fatal(http.ListenAndServe(":6060", nil))
	}()

	// generate CPU and memory activity
	for i := 0; i < *cpuCores; i++ {
		go cpuBurn()
	}
	go memChurn()

	time.Sleep(*duration)
}

var keep [][]byte

func cpuBurn() {
	for {
		x := 0.0
		for i := 1; i < 5_000_000; i++ {
			x += math.Sqrt(float64(i))
		}
		if x == -1 {
			log.Println(x) // prevent optimization
		}
	}
}

func memChurn() {
	const holdMB = 256 // keep ~256MB live
	keep = make([][]byte, 0, holdMB)
	i := 0
	for {
		b := make([]byte, 1<<20) // ~1MB
		for j := 0; j < len(b); j += 4096 {
			b[j] = 1
		}
		if len(keep) < holdMB {
			keep = append(keep, b)
		} else {
			keep[i%holdMB] = b
			i++
		}
		time.Sleep(150 * time.Millisecond)
	}
}
