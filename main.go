package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/containerd/cgroups"
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

	flags.cgroupPath = flag.String("path", "", "The cgroup path to read stats.")
	flags.onlyCPU = flag.Bool("cpu", false, "show cpu stats only")
	flags.onlyMem = flag.Bool("mem", false, "show mem stats only")
	flags.onlyPids = flag.Bool("pids", false, "show pids stats only")
	flags.onlyBlkio = flag.Bool("blkio", false, "show blkio stats only")
	flags.onlyHugetlb = flag.Bool("HugeTLB", false, "show hugetlb stats only")

	flag.Parse()
	return flags
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

	metrics, err := cg.Stat(cgroups.IgnoreNotExist)
	if err != nil {
		fmt.Println("err retrieving memory stats:", err)
		os.Exit(1)
	}

	fmt.Printf("Cgroup metrics: %+v\n", metrics)
}
