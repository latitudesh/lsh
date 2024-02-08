package prompt

type PromptInput interface {
	AssignValue(attributes interface{})
}

type Prompt []PromptInput

func New(inputs ...PromptInput) Prompt {
	var prompt Prompt

	for _, input := range inputs {
		prompt = append(prompt, input)
	}

	return prompt
}

func (p *Prompt) Run(attributes interface{}) {
	for _, input := range *p {
		input.AssignValue(attributes)
	}
}
