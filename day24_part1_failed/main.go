package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type ALU struct {
	w, x, y, z int
}

func (a *ALU) execute(instruction string, valueA, valueB int) {
	switch instruction {
	case "add":
		a.add(valueA, valueB)
	case "mul":
		a.mul(valueA, valueB)
	case "div":
		a.div(valueA, valueB)
	case "mod":
		a.mod(valueA, valueB)
	case "eql":
		a.eql(valueA, valueB)
	}
}

func (a *ALU) add(aVar, bVar int) {
	a.setVar(aVar, a.getVar(aVar)+a.getVar(bVar))
}

func (a *ALU) mul(aVar, bVar int) {
	a.setVar(aVar, a.getVar(aVar)*a.getVar(bVar))
}

func (a *ALU) div(aVar, bVar int) {
	a.setVar(aVar, a.getVar(aVar)/a.getVar(bVar))
}

func (a *ALU) mod(aVar, bVar int) {
	a.setVar(aVar, a.getVar(aVar)%a.getVar(bVar))
}

func (a *ALU) eql(aVar, bVar int) {
	if a.getVar(aVar) == a.getVar(bVar) {
		a.setVar(aVar, 1)
	} else {
		a.setVar(aVar, 0)
	}
}

func (a *ALU) setVar(variable int, value int) {
	switch variable {
	case 0:
		a.w = value
	case 1:
		a.x = value
	case 2:
		a.y = value
	case 3:
		a.z = value
	}
}

func (a *ALU) getVar(variable int) int {
	switch variable {
	case 0:
		return a.w
	case 1:
		return a.x
	case 2:
		return a.y
	case 3:
		return a.z
	default:
		return 0
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getVarIdx(varName string) int {
	switch varName {
	case "w":
		return 0
	case "x":
		return 1
	case "y":
		return 2
	case "z":
		return 3
	default:
		return -1
	}
}

func checkNumber(number int, remainingDigits int, program []string, maxModel *int) {
	if remainingDigits == 0 {
		alu := ALU{0, 0, 0, 0}
		for _, line := range program {
			parts := strings.Split(line, " ")
			switch parts[0] {
			case "inp":
				alu.setVar(getVarIdx(parts[1]), number%10)
				number /= 10
			default:
				valueA := getVarIdx(parts[1])
				valueB, _ := strconv.Atoi(parts[2])
				if valueB < 0 || valueB > 3 {
					alu.execute(parts[0], valueA, valueB)
				} else {
					alu.execute(parts[0], valueA, getVarIdx(parts[2]))
				}
			}
		}

		if alu.z == 0 {
			*maxModel = max(*maxModel, number)
		}
		return
	}

	for i := 1; i <= 9; i++ {
		checkNumber(number*10+i, remainingDigits-1, program, maxModel)
	}
}

func main() {
	data, err := ioutil.ReadFile("day24_part1/test.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	maxModel := 0
	checkNumber(0, 14, lines, &maxModel)
	fmt.Println("The largest valid model number accepted by MONAD:", maxModel)
}
