package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func cFunc(data []string, new_data []string) []string {
	answer := []string{}
	if len(new_data) == 0 {
		return answer
	}
	last, c := 0, 1
	for i := 1; i < len(new_data); i++ {
		if new_data[i] != new_data[last] {
			answer = append(answer, string(c)+data[last])
			c = 1
			last = i
		} else {
			c += 1
		}
	}
	//для последней строки
	answer = append(answer, string(c)+data[last])
	return answer
}

func dFunc(data []string, new_data []string) []string {
	answer := []string{}
	if len(new_data) == 0 {
		return answer
	}
	last, c := 0, 1
	for i := 1; i < len(new_data); i++ {
		if new_data[i] != new_data[last] {
			if c > 1 {
				answer = append(answer, data[last])
			}
			c = 1
			last = i
		} else {
			c += 1
		}
	}
	//для последней строки
	if c > 1 {
		answer = append(answer, data[last])
	}
	return answer
}

func uFunc(data []string, new_data []string) []string {
	answer := []string{}
	if len(new_data) == 0 {
		return answer
	}
	last, c := 0, 1
	for i := 1; i < len(new_data); i++ {
		if new_data[i] != new_data[last] {
			if c == 1 {
				answer = append(answer, data[last])
			}
			c = 1
			last = i
		} else {
			c += 1
		}
	}
	//для последней строки
	if c == 1 {
		answer = append(answer, data[last])
	}
	return answer
}

//шоб строчки переделать, как хотят
func fAndSfunc(data []string, num_fields int, num_chars int, i_flag bool) []string {
	new_data := []string{} //тут будут переделанные строки
	for i := 0; i < len(data); i++ {
		line := data[i]
		if i_flag { //если на регистр все равно
			line = strings.ToLower(line)
		}
		words := strings.Split(line, " ")
		words = words[num_fields:] //убрали первые n_fiels полей из строки
		new_line := strings.Join(words, " ")
		new_line = new_line[min(len(new_line), num_chars):] //убрали первые n_chars символов
		new_data = append(new_data, new_line)
	}
	return new_data
}

//если флагов ваще нет
func w_outFlags(data []string, new_data []string) []string {
	answer := []string{}
	if len(new_data) == 0 {
		return answer
	}
	last := 0
	for i := 1; i < len(new_data); i++ {
		if new_data[i] != new_data[last] {
			answer = append(answer, data[last])
			last = i
		}
	}
	//для последней строки
	answer = append(answer, data[last])
	return answer
}

//вывод
func outputFunc(writer io.Writer, output []string) {
	for _, line := range output {
		fmt.Fprintln(writer, line)
	}
}

func main() {

	//работаем с флажками
	var cFlag = flag.Bool("c", false, "подсчитать количество встречаний строки во входных данных")
	var dFlag = flag.Bool("d", false, "вывести только те строки, которые повторились во входных данных")
	var uFlag = flag.Bool("u", false, "вывести только те строки, которые не повторились во входных данных")
	var fFlag = flag.Int("f", 0, "не учитывать первые num_fields полей в строке")
	var sFlag = flag.Int("s", 0, "не учитывать первые num_chars символов в строке")
	var iFlag = flag.Bool("i", false, "не учитывать регистр букв")
	flag.Parse()

	file := os.Stdin //file - стандартный поток
	if flag.NArg() > 0 { //если после флага есть какой-то аргумент (нужен файл)
		var err error
		//открыли файлик
		file, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Printf("Ошибка при открытии файла: %v\n", err)
			return
		}
		defer file.Close() //закрыли файлик
	}

	//читаем файлик
	lines := []string{} //тут будут строки файла
	reader := bufio.NewReader(file)
	for { //считываем построчно
		line, err := reader.ReadString('\n')
		//удаляем \n
		line = strings.TrimRight(line, "\n")
		lines = append(lines, line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
	}

	if (*cFlag && *dFlag) || (*cFlag && *uFlag) || (*dFlag && *uFlag) {
		fmt.Println("Нельзя одновременно использовать флаги -c, -d и -u. Использование: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
		return
	}

	new_lines := fAndSfunc(lines, *fFlag, *sFlag, *iFlag) //переделанные строки

	//answer - итоговый ответ
	var answer []string
	if *cFlag {
		answer = cFunc(lines, new_lines)
	} else if *dFlag {
		answer = dFunc(lines, new_lines)
	} else if *uFlag {
		answer = uFunc(lines, new_lines)
	} else {
		answer = w_outFlags(lines, new_lines)
	}

	//тут мы вывод делаем крч
	if flag.NArg() > 1 {
		outputFile, err := os.Create(flag.Arg(1)) //второй аргумет - файл вывода
		if err != nil {
			fmt.Printf("Ошибка при записи в файл: %v\n", err)
			return
		}
		defer outputFile.Close()
		outputFunc(outputFile, answer)
	} else {
		outputFunc(os.Stdout, answer)
	}

}