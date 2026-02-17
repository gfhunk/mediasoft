package main

import "fmt"

func main() {
	fmt.Println("ТЕСТИРОВАНИЕ")

	testStack()
	fmt.Println()

	testQueue()
	fmt.Println()

	testLinkedList()
	fmt.Println()

	testRomanConverter()
	fmt.Println()

	testUniqueMatrix()
}

func testStack() {
	fmt.Println("СТЕК")
	stack := NewStack()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println("Добавили: 10, 20, 30")

	val, _ := stack.Peek()
	fmt.Printf("Вершина стека: %v\n", val)

	val, _ = stack.Pop()
	fmt.Printf("Извлекли: %v\n", val)
	val, _ = stack.Pop()
	fmt.Printf("Извлекли: %v\n", val)

	fmt.Printf("Стек пуст? %v\n", stack.IsEmpty())
}

func testQueue() {
	fmt.Println("ОЧЕРЕДЬ")
	queue := NewQueue()

	queue.Enqueue("первый")
	queue.Enqueue("второй")
	queue.Enqueue("третий")
	fmt.Println("Добавили: первый, второй, третий")

	val, _ := queue.Peek()
	fmt.Printf("Первый в очереди: %v\n", val)

	val, _ = queue.Dequeue()
	fmt.Printf("Извлекли: %v\n", val)
	val, _ = queue.Dequeue()
	fmt.Printf("Извлекли: %v\n", val)

	fmt.Printf("Очередь пуста? %v\n", queue.IsEmpty())
}

func testLinkedList() {
	fmt.Println("ОДНОСВЯЗНЫЙ СПИСОК")
	list := NewLinkedList()

	list.Add(100)
	list.Add(200)
	list.Add(300)
	fmt.Printf("Список: %v\n", list.Values())
	fmt.Printf("Размер: %d\n", list.Size)

	val, _ := list.Get(1)
	fmt.Printf("Элемент с индексом 1: %v\n", val)

	list.Remove(1)
	fmt.Printf("После удаления индекса 1: %v\n", list.Values())

	list.Add(400)
	fmt.Printf("После добавления 400: %v\n", list.Values())
}

func testRomanConverter() {
	fmt.Println("КОНВЕРТЕР РИМСКИХ ЦИФР")
	rc := NewRomanConverter()

	examples := rc.GetRomanExamples()
	fmt.Println("Римские → Арабские:")
	for roman, arabic := range examples {
		result, err := rc.RomanToArabic(roman)
		if err != nil {
			fmt.Printf("  Ошибка: %v\n", err)
		} else {
			fmt.Printf("  %s = %d (ожидалось %d) %v\n",
				roman, result, arabic, result == arabic)
		}
	}

	fmt.Println("\nАрабские → Римские:")
	testNumbers := []int{1, 4, 9, 40, 90, 400, 900, 2024, 1950, 3999}
	for _, num := range testNumbers {
		roman, err := rc.ArabicToRoman(num)
		if err != nil {
			fmt.Printf("  Ошибка: %v\n", err)
		} else {
			fmt.Printf("  %d = %s\n", num, roman)
		}
	}

	fmt.Println("\nОбработка ошибок:")
	_, err := rc.RomanToArabic("IIII")
	fmt.Printf("  'IIII' → ошибка: %v\n", err)

	_, err = rc.ArabicToRoman(4000)
	fmt.Printf("  4000 → ошибка: %v\n", err)
}

func testUniqueMatrix() {
	fmt.Println("МАТРИЦА С УНИКАЛЬНЫМИ ЧИСЛАМИ")

	um := NewUniqueMatrix(3, 3)

	err := um.GenerateRandomUnique(1, 20)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Println("Сгенерированная матрица 3x3:")
	um.PrintMatrix()

	fmt.Printf("\nСтатистика:\n")
	fmt.Printf("  Сумма всех элементов: %d\n", um.Sum())
	fmt.Printf("  Минимальное: %d\n", um.Min())
	fmt.Printf("  Максимальное: %d\n", um.Max())

	num := um.Matrix[1][1]
	row, col, found := um.Find(num)
	fmt.Printf("\nПоиск числа %d: строка %d, колонка %d (найдено: %v)\n",
		num, row, col, found)

	rowData, _ := um.GetRow(1)
	fmt.Printf("Вторая строка: %v\n", rowData)

	colData, _ := um.GetColumn(1)
	fmt.Printf("Вторая колонка: %v\n", colData)

	fmt.Println("\nПроверка уникальности:")
	numbers := make(map[int]bool)
	allUnique := true
	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			if numbers[um.Matrix[i][j]] {
				allUnique = false
				fmt.Printf("  Найдено дублирование: %d\n", um.Matrix[i][j])
			}
			numbers[um.Matrix[i][j]] = true
		}
	}
	if allUnique {
		fmt.Println("  Все числа уникальны! ✓")
	}
}
