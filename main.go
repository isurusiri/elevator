package main

import (
	"bufio"
	"elevator"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxFloorNumber = 21
const minFloorNumber = 0

type controlState int

const (
	Status controlState = iota
	Pickup
	Step
	Exit
	Error
)

func StringToInt(value string) int {
	result, _ := strconv.ParseInt(value, 10, 64)
	return int(result)
}

func StringToIntSlice(in []string) []int {
	out := make([]int, len(in))

	for i, val := range in {
		out[i] = StringToInt(val)
	}

	return out
}

func readFromStdin() string {
	fmt.Print("$ ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = line[0 : len(line)-1]
	return line
}

func formatCmd(line string) (controlState, []int) {
	if line == "status" {
		return Status, nil
	} else if line == "step" {
		return Step, nil
	} else if line == "exit" {
		return Exit, nil
	} else if strings.HasPrefix(line, "pickup") {
		line = strings.Trim(line, "pickup ")
		args := strings.Split(line, " ")
		params := StringToIntSlice(args)

		if len(params) == 2 &&
			(params[0] <= maxFloorNumber && params[0] >= minFloorNumber) &&
			(params[1] == 1 || params[1] == -1) {
			return Pickup, params
		}
	}

	return Error, nil
}

var flagNumElev = flag.Int("n", 0, "intflag")

func main() {
	flag.Parse()
	if *flagNumElev <= 0 {
		fmt.Printf("Usage: %s -n NumberOfElevators\n", os.Args[0])
		os.Exit(1)
	}

	ecs := elevator.NewElevatorControlSystem(*flagNumElev)

	for {
		line := readFromStdin()
		cmd, args := formatCmd(line)

		switch cmd {
		case Status:
			stat := ecs.Status()
			for _, e := range stat {
				fmt.Println(e)
			}
		case Pickup:
			ecs.Pickup(args[0], args[1])
		case Step:
			ecs.Step()
		case Exit:
			return
		case Error:
			fmt.Println("invalid cmd")
		}
	}
}
