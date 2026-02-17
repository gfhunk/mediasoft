package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UniqueMatrix struct {
	Rows    int
	Cols    int
	Matrix  [][]int
	Numbers map[int]bool
}

func NewUniqueMatrix(rows, cols int) *UniqueMatrix {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	return &UniqueMatrix{
		Rows:    rows,
		Cols:    cols,
		Matrix:  matrix,
		Numbers: make(map[int]bool),
	}
}

func (um *UniqueMatrix) GenerateRandomUnique(min, max int) error {
	totalCells := um.Rows * um.Cols
	availableNumbers := max - min + 1

	if availableNumbers < totalCells {
		return errors.New("недостаточно уникальных чисел в диапазоне")
	}

	rand.Seed(time.Now().UnixNano())

	allNumbers := make([]int, availableNumbers)
	for i := 0; i < availableNumbers; i++ {
		allNumbers[i] = min + i
	}

	for i := len(allNumbers) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		allNumbers[i], allNumbers[j] = allNumbers[j], allNumbers[i]
	}

	index := 0
	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			um.Matrix[i][j] = allNumbers[index]
			um.Numbers[allNumbers[index]] = true
			index++
		}
	}

	return nil
}

func (um *UniqueMatrix) GenerateRandomUniqueV2(min, max int) error {
	totalCells := um.Rows * um.Cols
	availableNumbers := max - min + 1

	if availableNumbers < totalCells {
		return errors.New("недостаточно уникальных чисел в диапазоне")
	}

	rand.Seed(time.Now().UnixNano())
	um.Numbers = make(map[int]bool)

	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			for {
				num := min + rand.Intn(availableNumbers)
				if !um.Numbers[num] {
					um.Matrix[i][j] = num
					um.Numbers[num] = true
					break
				}
			}
		}
	}

	return nil
}

func (um *UniqueMatrix) Get(row, col int) (int, error) {
	if row < 0 || row >= um.Rows || col < 0 || col >= um.Cols {
		return 0, errors.New("индекс вне диапазона")
	}
	return um.Matrix[row][col], nil
}

func (um *UniqueMatrix) Set(row, col, value int) error {
	if row < 0 || row >= um.Rows || col < 0 || col >= um.Cols {
		return errors.New("индекс вне диапазона")
	}

	if um.Numbers[value] {
		return errors.New("число уже существует в матрице")
	}

	oldValue := um.Matrix[row][col]
	delete(um.Numbers, oldValue)

	um.Matrix[row][col] = value
	um.Numbers[value] = true

	return nil
}

func (um *UniqueMatrix) PrintMatrix() {
	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			fmt.Printf("%4d ", um.Matrix[i][j])
		}
		fmt.Println()
	}
}

func (um *UniqueMatrix) GetRow(row int) ([]int, error) {
	if row < 0 || row >= um.Rows {
		return nil, errors.New("строка вне диапазона")
	}

	result := make([]int, um.Cols)
	copy(result, um.Matrix[row])
	return result, nil
}

func (um *UniqueMatrix) GetColumn(col int) ([]int, error) {
	if col < 0 || col >= um.Cols {
		return nil, errors.New("колонка вне диапазона")
	}

	result := make([]int, um.Rows)
	for i := 0; i < um.Rows; i++ {
		result[i] = um.Matrix[i][col]
	}
	return result, nil
}

func (um *UniqueMatrix) Find(value int) (row, col int, found bool) {
	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			if um.Matrix[i][j] == value {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func (um *UniqueMatrix) Sum() int {
	sum := 0
	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			sum += um.Matrix[i][j]
		}
	}
	return sum
}

func (um *UniqueMatrix) Min() int {
	min := um.Matrix[0][0]
	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			if um.Matrix[i][j] < min {
				min = um.Matrix[i][j]
			}
		}
	}
	return min
}

func (um *UniqueMatrix) Max() int {
	max := um.Matrix[0][0]
	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			if um.Matrix[i][j] > max {
				max = um.Matrix[i][j]
			}
		}
	}
	return max
}
