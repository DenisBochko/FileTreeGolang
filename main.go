package main

import (
	"fmt"
	"os"
)

func main() {
	out := os.Stdout

	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}

	path := os.Args[1]
	printFiles := false
	printFiles = len(os.Args) == 3 && os.Args[2] == "-f"

	fmt.Fprintln(out, path)
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

// функция dirTree является входной, она вызывает уже рекурсивную функцию
func dirTree(out *os.File, path string, printFiles bool) error {
	if printFiles {

		err := dirTreeRecursion(out, path, "", printFiles)
		if err != nil {
			return err
		}

	} else {

		err := dirTreeRecursion(out, path, "", printFiles)
		if err != nil {
			return err
		}

	}

	return nil
}

func dirTreeRecursion(out *os.File, path string, preffix string, printFiles bool) error {
	// Открывает директорию path
	dir, err := os.Open(path)
	if err != nil {
		return err
	}
	// Закрывам директорию в конце функции
	defer dir.Close()
	// Считываем слайс файлов
	files, err := dir.ReadDir(-1)
	if err != nil {
		return err
	}

	/*
		Проходимся в цикле по всем файла директории,
			чтобы найти последний файл, который будет являться директорией

		Это необходимо для случая "go run main.go ."
	*/
	lastDir := 0
	for index, file := range files {
		if file.IsDir() {
			lastDir = index
		}
	}
	// Проходимся по всем файлам в цикле
	for index, file := range files {
		/*
			Выбираем символ разделителя в зависимости от того,
				является ли элемент file последним в слайсе
		*/
		branchSymbol := ""
		/*
			Первое условие для стандартного случая,
				второе - для случая, когда не нужно выводить файлы
		*/
		if index == (len(files)-1) || (index == lastDir && !printFiles) {
			branchSymbol = "└───"
		} else {
			branchSymbol = "├───"
		}

		// Печатаем строку
		if printFiles {
			fmt.Fprintln(out, preffix+branchSymbol+file.Name())
		} else {
			if file.IsDir() && !printFiles {
				fmt.Fprintln(out, preffix+branchSymbol+file.Name())
			}
		}
		/*
			Если элемент file является директорией,
				то рукурсивно вызываем dirTreeRecursion с новым preffix
		*/
		if file.IsDir() {
			newPreffix := preffix
			/*
				Первое условие для стандартного случая,
					второе - для случая, когда не нужно выводить файлы
			*/

			if index == (len(files)-1) || (index == lastDir && !printFiles) {
				newPreffix += "    "
			} else {
				newPreffix += "│   "
			}
			dirTreeRecursion(out, path+"\\"+file.Name(), newPreffix, printFiles)
		}
	}
	return nil
}
