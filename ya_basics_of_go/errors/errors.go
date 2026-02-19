package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type SliceOfErrors []error

func (soe *SliceOfErrors) Error() string {
	var result strings.Builder
	for _, v := range *soe {
		formattedError := fmt.Errorf("%w;\n", v)
		result.WriteString(formattedError.Error())
	}

	return result.String()
}

func MyCheck(str string) error {
	sliceOfErrors := SliceOfErrors{}
	switch {
	case strings.ContainsAny(str, "1234567890"):
		sliceOfErrors = append(sliceOfErrors, errors.New("found numbers"))
		fallthrough
	case strings.Contains(str, "  "):
		sliceOfErrors = append(sliceOfErrors, errors.New("no two spaces"))
		fallthrough
	case len(str) > 20:
		sliceOfErrors = append(sliceOfErrors, errors.New("line is too long"))
	default:
		return nil
	}
	return &sliceOfErrors

}

func main() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}
