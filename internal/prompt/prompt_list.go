package prompt

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/latitudesh/lsh/internal/utils"
	"github.com/manifoldco/promptui"
)

type InputList struct {
	Name  string
	Label string
}

func NewInputList(name, label string) *InputList {
	return &InputList{
		Name:  name,
		Label: label,
	}
}

func (p *InputList) AssignValue(attributes interface{}) {
	currentValue := utils.GetFieldValue(attributes, p.Name)

	if currentValue.Kind() == reflect.Slice && currentValue.Len() == 0 {
		prompt := promptui.Prompt{
			Label: fmt.Sprintf("%v (separate by comma, for example: option1,option2)", p.Label),
		}

		value, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}

		values := strings.Split(value, ",")

		for i, v := range values {
			values[i] = strings.TrimSpace(v)
		}

		utils.AssignValue(attributes, p.Name, values)
	}
}
