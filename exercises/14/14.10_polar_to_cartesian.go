package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type polar struct {
	radius float64
	theta  float64
}

type cartesian struct {
	x float64
	y float64
}

const result = "polar: radius=%.02f angle=%.02f cartesian: x=%.02f y=%.02f\n"

var promt = "%s to quit"

func init() {
	if runtime.GOOS == "windows" {
		promt = fmt.Sprintf(promt, "Ctrl+Z, Enter")
	} else {
		promt = fmt.Sprintf(promt, "Ctrl+D")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-questions
			theta := polarCoord.theta * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(theta)
			y := polarCoord.radius * math.Sin(theta)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(promt)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1]
		if numbers := strings.Fields(line); len(numbers) == 2 {
			polars, err := floatForStrings(numbers)
			if err != nil {
				fmt.Fprintln(os.Stderr, "invalid number")
				continue
			}
			questions <- polar{polars[0], polars[1]}
			coord := <-answers
			fmt.Printf(result, polars[0], polars[1], coord.x, coord.y)
		} else {
			fmt.Fprintln(os.Stderr, "invalid input")
		}
	}
}

func floatForStrings(numbers []string) ([]float64, error) {
	var floats []float64
	for _, number := range numbers {
		if x, err := strconv.ParseFloat(number, 64); err != nil {
			return nil, err
		} else {
			floats = append(floats, x)
		}
	}
	return floats, nil
}
