package main 

import (
	"fmt"
	"UJS/file"
	"strconv"
	"strings"
	"math"
)

type dictionary struct {
	size 		[]int
	basic 		[]int
	nonBasic 	[]int
	coeff		[]float64
	matrix		[][]float64
	objective	[]float64
}

const fileName string = "D:\\James\\Google Drive\\coding\\linear programming\\assignmentParts\\part5.dict"

var currEnterLeave []int = []int{0,0}

func main() {
	status := 0
	// import data structure
	dict := loadDictionary()
	fmt.Println(dict)
	// pivot operation
	dict, status = pivot(dict)
	// print objective value
	switch status {
	case -1:
		fmt.Println("FINAL")
	case -2:
		fmt.Println("UNBOUNDED")
	default:
		fmt.Println(currEnterLeave[0])
		fmt.Println(currEnterLeave[1])
		fmt.Println(dict.objective[0])
	}
}

func pivot(dict dictionary) (dictionary, int) {
	// choose entering variable
	enter := entering(dict.objective, dict.nonBasic)
	// if no valid entering variable, return
	if enter < 0 {
		return dict, enter
	}
	// print entering variable
	currEnterLeave[0] = dict.nonBasic[enter]
	// choose leaving variable
	leave := leaving(dict.matrix, dict.coeff, dict.basic, enter)
	// if unbounded for enter, return
	if leave < 0 {
		return dict, leave
	}
	// print leaving variable
	currEnterLeave[1] = dict.basic[leave]
	// calculate enter var in terms of leaving var
	dict.matrix[leave], dict.coeff = invert(dict.matrix[leave], dict.coeff, leave, enter)
	// recalculate values of other basic variables
	dict.matrix, dict.coeff = convert(dict.matrix, dict.coeff, enter, leave)
	// recalculate objective
	dict = reObjective(dict, enter, leave)
	// swap enter/leave position
	dict.basic[leave], dict.nonBasic[enter] = dict.nonBasic[enter], dict.basic[leave]
	// return new dictionary
	return dict, 0
}

func reObjective(dict dictionary, enter int, leave int) dictionary {
	dict.objective[0] += dict.coeff[leave] * dict.objective[enter+1]
	for i := 0; i < dict.size[1]; i++ {
		if i != leave {
			dict.objective[i+1] += dict.objective[enter+1] * dict.matrix[leave][i]
		}
	}
	dict.objective[enter+1] = dict.matrix[leave][enter] * dict.objective[enter+1]
	return dict
}

func convert(dict [][]float64, coeff []float64, enter int, leave int) ([][]float64, []float64) {
	for i, v := range dict {
		if i != leave {
			coeff[i] += coeff[leave] * v[enter]
			for j, w := range v {
				if j != enter {
					w += dict[leave][j] * v[enter]
				}
			}
			v[enter] = dict[leave][enter] * v[enter]
		}
	}
	return dict, coeff
}

func invert(mRow []float64, coeff []float64, leave int, enter int) ([]float64, []float64) {
	coeff[leave] = coeff[leave] / -(mRow[enter])
	for i, v := range mRow {
		if i != enter {
			v = v / -(mRow[enter])
		}
	}
	mRow[enter] = 1.0 / -(mRow[enter])
	return mRow, coeff
}

func entering(objective []float64, nonBasic []int) int {
	candidates := make([]int, 0)
	for i := 1; i < len(objective); i++ {
		if objective[i] > 0.0 {
			candidates = append(candidates, i)
		}
	}
	if len(candidates) == 0 {
		return -1
	} else {
		leave := candidates[0] - 1
		for _, v := range candidates {
			if nonBasic[v-1] < nonBasic[leave] {
				leave = v-1
			}
		}
		return leave
	}
}

func leaving(matrix [][]float64, coeff []float64, basic []int, enter int) int {
	leave := -2
	currentConstraint := math.Expm1(64)
	for i, v := range matrix {
		if v[enter] < 0 {
			if q := coeff[i] / -(v[enter]); leave >= 0 && q == currentConstraint {
				if basic[i] < basic[leave] {
					leave = i
				}
			} else if q < currentConstraint {
				currentConstraint = q
				leave = i
			}
		} 
	}
	return leave
}

func loadDictionary() dictionary {
	dict := dictionary{}
	//read file lines
	input := file.FileLines(fileName)
	// break out first line, load its values into size
	dict.size = intSeparate(input[0])
	input = input[1:]
	// break out second line, load it into basic
	dict.basic = intSeparate(input[0])
	input = input[1:]
	// break out third line, load it into nonBasic
	dict.nonBasic = intSeparate(input[0])
	input = input[1:]
	// break out fourth line, load it into coeff
	dict.coeff = floatSeparate(input[0])
	input = input[1:]
	// for next size[0] lines, load into matrix
	dict.matrix = make([][]float64, dict.size[0])
	for i := 0; i < dict.size[0]; i++ {
		dict.matrix[i] = floatSeparate(input[0])
		input = input[1:]
	}
	// load final line into objective
	dict.objective = floatSeparate(input[0])
	return dict
}

func intSeparate(input string) []int {
	newString := (strings.Split(input, "\n"))[0]
	newString = strings.TrimSpace(newString)
	newString = strings.Replace(newString, "  ", " ", -1)
	a := strings.Split(newString, " ")
	intSlice := make([]int, len(a))
	for i, v := range a {
		intSlice[i], _ = strconv.Atoi(v)
	}
	return intSlice
}

func floatSeparate(input string) []float64 {
	newString := (strings.Split(input, "\n"))[0]
	newString = strings.TrimSpace(newString)
	newString = strings.Replace(newString, "  ", " ", -1)
	a := strings.Split(newString, " ")
	floatSlice := make([]float64, len(a))
	for i, v := range a {
		floatSlice[i], _ = strconv.ParseFloat(v, 64)
	}
	return floatSlice
}