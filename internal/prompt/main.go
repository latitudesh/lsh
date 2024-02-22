package prompt

import "fmt"

type PromptInput interface {
	AssignValue(attributes interface{})
}

type Prompt struct {
	Inputs      []PromptInput
	Description string
	ParamsType  string
}

func (p *Prompt) Run(params interface{}) {
	if p.Description != "" {
		fmt.Printf("\n%v\n\n", p.Description)
	}

	for _, v := range p.Inputs {
		v.AssignValue(params)
	}
}
