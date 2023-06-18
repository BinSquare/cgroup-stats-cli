package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/containerd/cgroups"
	v1 "github.com/containerd/cgroups/stats/v1"
)

type cliOptions struct {
	cgroupPath  *string
	onlyCPU     *bool
	onlyMem     *bool
	onlyPids    *bool
	onlyBlkio   *bool
	onlyHugetlb *bool
}

func parseFlags(flags *cliOptions) *cliOptions {
	flags.cgroupPath = flag.String("path", "", "The cgroup path to read stats. Path should not include `/sys/fs/cgroup/` prefix, it should start with your own cgroups name")
	flags.onlyCPU = flag.Bool("cpu", false, "show cpu stats only")
	flags.onlyMem = flag.Bool("mem", false, "show mem stats only")
	flags.onlyPids = flag.Bool("pids", false, "show pids stats only")
	flags.onlyBlkio = flag.Bool("blkio", false, "show blkio stats only")
	flags.onlyHugetlb = flag.Bool("HugeTLB", false, "show hugetlb stats only")

	flag.Parse()
	return flags
}

func getMetrics(cg cgroups.Cgroup) *v1.Metrics {
	metrics, err := cg.Stat(cgroups.IgnoreNotExist)
	if err != nil {
		fmt.Println("err retrieving cgroup stats:", err)
		os.Exit(1)
	}

	return metrics
}

func printMetrics(metrics *v1.Metrics, flags *cliOptions) {
	if *flags.onlyCPU {
		fmt.Printf("CPU Metrics: %+v\n", metrics.CPU)
	}
	if *flags.onlyMem {
		fmt.Printf("Memory Metrics: %+v\n", metrics.Memory)
	}
	if *flags.onlyPids {
		fmt.Printf("PIDs Metrics: %+v\n", metrics.Pids)
	}
	if *flags.onlyBlkio {
		fmt.Printf("BlkIO Metrics: %+v\n", metrics.Blkio)
	}
	if *flags.onlyHugetlb {
		fmt.Printf("HugeTLB Metrics: %+v\n", metrics.Hugetlb)
	}

	if !*flags.onlyCPU && !*flags.onlyMem && !*flags.onlyPids && !*flags.onlyBlkio && !*flags.onlyHugetlb {
		fmt.Printf("Cgroup metrics: %+v\n", metrics)
	}
}

func main() {
	flags := parseFlags(new(cliOptions))

	if *flags.cgroupPath == "" {
		fmt.Println("Missing cgroup path")
		os.Exit(1)
	}

	cg, err := cgroups.Load(cgroups.V1, cgroups.StaticPath(*flags.cgroupPath))
	if err != nil {
		fmt.Println("loading cgroup error:", err)
		os.Exit(1)
	}

	metrics := getMetrics(cg)

	printMetrics(metrics, flags)
}
