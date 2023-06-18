# Cgroup Stats CLI Tool
This repository contains a command line tool written in Go that uses the containerd/cgroups library to print out cgroup (control group) statistics for a given cgroup. You can specify the type of stats to print out using command line flags.

### Dependencies
This tool requires the containerd/cgroups Go library. You can install it using the following command:
````
go get github.com/containerd/cgroups
````

### Usage
You can run the tool using the go run command. You must provide the cgroup path using the -path flag. You can also specify what kind of stats to print out using the -cpu, -mem, -pids, -blkio, and -hugetlb flags.
```
go run main.go -path=<cgroup_path> [-cpu] [-mem] [-pids] [-blkio] [-hugetlb]
```
For example, the following command will print out the CPU, Memory, and PIDs stats of the specified cgroup:

```
go run main.go -path=/sys/fs/cgroup/unified/system.slice/docker.service -cpu -mem -pids
```

### Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.