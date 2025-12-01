package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var input = "2024/day17/input.txt"

type Computer struct {
    regA, regB, regC int
    ip               int
    program          []int
    output           []int
}

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %s", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func run() (string, int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1, part2 := "", 0

    computer := &Computer{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in := scanner.Text()
		if in != "" {
			if strings.HasPrefix(in, "Register") {
                if strings.Contains(in, "A") {
                    computer.regA, _ = strconv.Atoi(strings.Split(in, " ")[2])
                }
                if strings.Contains(in, "B") {
                    computer.regB, _ = strconv.Atoi(strings.Split(in, " ")[2])
                }
                if strings.Contains(in, "C") {
                    computer.regC, _ = strconv.Atoi(strings.Split(in, " ")[2])
                }
            } else if strings.HasPrefix(in, "Program") {
                p := strings.Split(in, " ")
                tasks := strings.Split(p[1], ",")
                for _, task := range tasks {
                    num, _ := strconv.Atoi(task)
                    computer.program = append(computer.program, num)
                }
            }
		}
	}
    
    // computer.execute()
    
    // result := make([]string, len(computer.output))
    // for i, v := range computer.output {
    //     result[i] = fmt.Sprint(v)
    // }

    
    // part1 = strings.Join(result, ",")
    
    part2 = computer.findCopy()

	return part1, part2
}

func (c *Computer) findCopy() int {
    // a := 35190494000000
    a := 57768335000000
    programStr := []string{}
    for _, v := range c.program {
        programStr = append(programStr, strconv.Itoa(v))
    }
    check := strings.Join(programStr, ",")
    for {
        c.regA = a
        c.regB = 0
        c.regC = 0
        c.ip = 0
        c.output = []int{}

        c.execute()

        result := []string{}
        for _, v := range c.output {
            result = append(result, fmt.Sprint(v))
        }
        output := strings.Join(result, ",")

        if output == check {
            return a
        }

        if len(c.output) < len(check) {
            a += int(math.Pow(2, float64(len(check)-len(c.output))))
            fmt.Println("incrementing", len(c.output), len(check), len(c.output) < len(check))
            fmt.Println(a)
            fmt.Println(output)
            fmt.Println(check)
        }
        if a % 1000000 == 0 {
            fmt.Println(a)
            fmt.Println(output)
            fmt.Println(check)
            fmt.Println()
        }
        a++
    }
}

func (c *Computer) getComboValue(operand int) int {
    switch operand {
    case 0, 1, 2, 3:
        return operand
    case 4:
        return c.regA
    case 5:
        return c.regB
    case 6:
        return c.regC
    default:
        return 0
    }
}

func (c *Computer) execute() {
    for c.ip < len(c.program)-1 {
        opcode := c.program[c.ip]
        operand := c.program[c.ip+1]
        combo := c.getComboValue(operand)
        
        switch opcode {
        case 0: // adv
            c.regA = c.regA / int(math.Pow(2, float64(combo)))
        case 1: // bxl
            c.regB = c.regB ^ operand
        case 2: // bst
            c.regB = combo % 8
        case 3: // jnz
            if c.regA != 0 {
                c.ip = operand
                continue
            }
        case 4: // bxc
            c.regB = c.regB ^ c.regC
        case 5: // out
            c.output = append(c.output, combo%8)
        case 6: // bdv
            c.regB = c.regA / int(math.Pow(2, float64(combo)))
        case 7: // cdv
            c.regC = c.regA / int(math.Pow(2, float64(combo)))
        }
        
        c.ip += 2
    }
}
