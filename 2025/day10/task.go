package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var input = "2025/day10/input.txt"

func main() {
	start := time.Now()

	part1, part2 := run()

	log.Println(fmt.Sprintf("Part 1: %d", part1))
	log.Println(fmt.Sprintf("Part 2: %d", part2))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

type Machine struct {
	lights []bool
	buttons [][]int
	joltage []int
}

func run() (int, int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	machines := []Machine{}

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        in := scanner.Text()
        if in != "" {
			buttonStart := strings.Index(in, "(")
			buttonEnd := strings.LastIndex(in, ")")
			joltageStart := strings.Index(in, "{")

			lightsStr := in[:buttonStart]
			buttonsStr := in[buttonStart:buttonEnd+1]
			joltageStr := in[joltageStart:]

			lights := []bool{}
			for _, l := range lightsStr {
				if l == '#' {
					lights = append(lights, true)
				} else if l == '.' {
					lights = append(lights, false)
				}
			}

			buttonsStr = strings.Trim(buttonsStr, "()")
			buttonParts := strings.Split(buttonsStr, ") (")
			buttons := [][]int{}
			for _, bp := range buttonParts {
				button := []int{}
				for _, numStr := range strings.Split(bp, ",") {
					num, _ := strconv.Atoi(numStr)
					button = append(button, num)
				}
				buttons = append(buttons, button)
			}

			joltageStr = strings.Trim(joltageStr, "{}")
			joltage := []int{}
			for _, jStr := range strings.Split(joltageStr, ",") {
				num, _ := strconv.Atoi(jStr)
				joltage = append(joltage, num)
			}
			machines = append(machines, Machine{
				lights:   lights,
				buttons:  buttons,
				joltage:  joltage,
			})
        }
    }  

	part1, part2 := 0, 0

	for _, machine := range machines {
		part1 += optimalPressesLights(machine)
		part2 += optimalPressesJoltage(machine)
	}

	return part1, part2
}

func optimalPressesLights(machine Machine) int {
    numLights := len(machine.lights)
    numButtons := len(machine.buttons)
    
    A := make([][]int, numLights)
    for i := range A {
        A[i] = make([]int, numButtons)
    }
    
    for buttonIdx, button := range machine.buttons {
        for _, lightIdx := range button {
            if lightIdx < numLights {
                A[lightIdx][buttonIdx] = 1
            }
        }
    }
    
    b := make([]int, numLights)
    for i, light := range machine.lights {
        if light {
            b[i] = 1
        }
    }
    
    solution := solveGaussianElimination(A, b, numButtons)
    
    if solution == nil {
        return 0
    }
    
    presses := 0
    pressedButtons := []int{}
    for i, press := range solution {
        if press == 1 {
            pressedButtons = append(pressedButtons, i)
            presses++
        }
    }
        
    verify := make([]int, numLights)
    for i, press := range solution {
        if press == 1 {
            for _, lightIdx := range machine.buttons[i] {
                if lightIdx < numLights {
                    verify[lightIdx] ^= 1
                }
            }
        }
    }    
    return presses
}

func solveGaussianElimination(A [][]int, b []int, numButtons int) []int {
    numLights := len(A)

    augmented := make([][]int, numLights)
    for i := range augmented {
        augmented[i] = make([]int, numButtons+1)
        copy(augmented[i], A[i])
        augmented[i][numButtons] = b[i]
    }
    
    pivotCols := []int{}
    pivotRows := make([]int, numButtons)
    for i := range pivotRows {
        pivotRows[i] = -1
    }
    currentRow := 0
    
    for col := 0; col < numButtons && currentRow < numLights; col++ {
        pivotRow := -1
        for row := currentRow; row < numLights; row++ {
            if augmented[row][col] == 1 {
                pivotRow = row
                break
            }
        }
        
        if pivotRow == -1 {
            continue
        }
        
        augmented[currentRow], augmented[pivotRow] = augmented[pivotRow], augmented[currentRow]
        pivotCols = append(pivotCols, col)
        pivotRows[col] = currentRow
        
        for row := 0; row < numLights; row++ {
            if row != currentRow && augmented[row][col] == 1 {
                for c := 0; c <= numButtons; c++ {
                    augmented[row][c] ^= augmented[currentRow][c]
                }
            }
        }
        
        currentRow++
    }
    
    for row := currentRow; row < numLights; row++ {
        if augmented[row][numButtons] == 1 {
            return nil
        }
    }
    
    freeVars := []int{}
    for col := 0; col < numButtons; col++ {
        if pivotRows[col] == -1 {
            freeVars = append(freeVars, col)
        }
    }
    
    var bestSolution []int
    minPresses := -1
    numCombinations := 1 << len(freeVars)
    
    for combo := 0; combo < numCombinations; combo++ {
        solution := make([]int, numButtons)
        
        for i, freeVar := range freeVars {
            if (combo >> i) & 1 == 1 {
                solution[freeVar] = 1
            }
        }
        
        for i := len(pivotCols) - 1; i >= 0; i-- {
            col := pivotCols[i]
            row := pivotRows[col]
            
            val := augmented[row][numButtons]
            for j := col + 1; j < numButtons; j++ {
                val ^= (augmented[row][j] * solution[j])
            }
            solution[col] = val
        }
        
        presses := 0
        for _, press := range solution {
            presses += press
        }
        
        if minPresses == -1 || presses < minPresses {
            minPresses = presses
            bestSolution = solution
        }
    }
    
    return bestSolution
}

func optimalPressesJoltage(machine Machine) int {
    numJoltages := len(machine.joltage)
    numButtons := len(machine.buttons)
    
    A := make([][]int, numJoltages)
    for i := range A {
        A[i] = make([]int, numButtons)
    }
    
    for buttonIdx, button := range machine.buttons {
        for _, joltageIdx := range button {
            if joltageIdx < numJoltages {
                A[joltageIdx][buttonIdx] = 1
            }
        }
    }
    
    b := make([]int, numJoltages)
    copy(b, machine.joltage)
    
    solution := solveIntegerLinearSystem(A, b, numButtons)
    
    if solution == nil {
        return 0
    }
    
    presses := 0
    for _, press := range solution {
        presses += press
    }
    
    return presses
}

func solveIntegerLinearSystem(A [][]int, b []int, numButtons int) []int {
    numEquations := len(A)
    
    augmented := make([][]int, numEquations)
    for i := range augmented {
        augmented[i] = make([]int, numButtons+1)
        copy(augmented[i], A[i])
        augmented[i][numButtons] = b[i]
    }
    
    pivotCols := []int{}
    pivotRows := make(map[int]int)
    currentRow := 0
    
    for col := 0; col < numButtons && currentRow < numEquations; col++ {
        pivotRow := -1
        for row := currentRow; row < numEquations; row++ {
            if augmented[row][col] != 0 {
                pivotRow = row
                break
            }
        }
        
        if pivotRow == -1 {
            continue
        }
        
        augmented[currentRow], augmented[pivotRow] = augmented[pivotRow], augmented[currentRow]
        pivotCols = append(pivotCols, col)
        pivotRows[col] = currentRow
        
        for row := 0; row < numEquations; row++ {
            if row != currentRow && augmented[row][col] != 0 {
                factor := augmented[row][col]
                pivot := augmented[currentRow][col]
                
                for c := 0; c <= numButtons; c++ {
                    augmented[row][c] = augmented[row][c]*pivot - augmented[currentRow][c]*factor
                }
            }
        }
        
        currentRow++
    }
    
    for row := currentRow; row < numEquations; row++ {
        if augmented[row][numButtons] != 0 {
            return nil
        }
    }
    
    freeVars := []int{}
    for col := 0; col < numButtons; col++ {
        if _, exists := pivotRows[col]; !exists {
            freeVars = append(freeVars, col)
        }
    }
    
    var bestSolution []int
    minPresses := -1
    
    maxVal := 500
    if len(freeVars) == 0 {
        maxVal = 1
    } else if len(freeVars) == 1 {
        maxVal = 500
    } else if len(freeVars) == 2 {
        maxVal = 250
    } else if len(freeVars) == 3 {
        maxVal = 50
    } else {
        maxVal = 20
    }
    
    generateCombinations(freeVars, maxVal, 0, make([]int, numButtons), &bestSolution, &minPresses, augmented, pivotCols, pivotRows, numButtons)
    
    return bestSolution
}

func generateCombinations(freeVars []int, maxVal, idx int, current []int, bestSolution *[]int, minPresses *int, augmented [][]int, pivotCols []int, pivotRows map[int]int, numButtons int) {
    if idx == len(freeVars) {
        solution := make([]int, numButtons)
        copy(solution, current)
        
        valid := true
        for i := len(pivotCols) - 1; i >= 0; i-- {
            col := pivotCols[i]
            row := pivotRows[col]
            
            val := augmented[row][numButtons]
            for j := col + 1; j < numButtons; j++ {
                val -= augmented[row][j] * solution[j]
            }
            
            if augmented[row][col] == 0 || val%augmented[row][col] != 0 {
                valid = false
                break
            }
            
            solution[col] = val / augmented[row][col]
            
            if solution[col] < 0 {
                valid = false
                break
            }
        }
        
        if !valid {
            return
        }
        
        presses := 0
        for _, press := range solution {
            presses += press
        }
        
        if *minPresses == -1 || presses < *minPresses {
            *minPresses = presses
            *bestSolution = make([]int, numButtons)
            copy(*bestSolution, solution)
        }
        return
    }
    
    for val := 0; val < maxVal; val++ {
        current[freeVars[idx]] = val
        generateCombinations(freeVars, maxVal, idx+1, current, bestSolution, minPresses, augmented, pivotCols, pivotRows, numButtons)
    }
}