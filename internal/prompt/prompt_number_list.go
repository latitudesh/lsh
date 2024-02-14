package prompt

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/latitudesh/lsh/internal/utils"
	"github.com/manifoldco/promptui"
)

type InputNumberList struct {
	Name  string
	Label string
}

func NewInputNumberList(name, label string) *InputNumberList {
	return &InputNumberList{
		Name:  name,
		Label: label,
	}
}

func (p *InputNumberList) AssignValue(attributes interface{}) {
	currentValue := utils.GetFieldValue(attributes, p.Name)

	if currentValue.Kind() == reflect.Slice && currentValue.Len() == 0 {
		prompt := promptui.Prompt{
			Label: fmt.Sprintf("%v (separate by comma, for example: option1,option2)", p.Label),
		}

		value, err := prompt.Run()
		if err != nil {
			exitPrompt(err)
		}

		if len(strings.TrimSpace(value)) <= 0 {
			return
		}

		var numbers []int64

		values := strings.Split(value, ",")
		for _, v := range values {
			number, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				exitPrompt(err)
			}

			numbers = append(numbers, int64(number))
		}

		utils.AssignValue(attributes, p.Name, numbers)
	}
}

func exitPrompt(err error) {
	fmt.Printf("Failed to read input: %v\n", err)
	os.Exit(1)
}
