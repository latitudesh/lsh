package prompt

type PromptInput interface {
	AssignValue(attributes interface{})
}

type Prompt []PromptInput

func RunPrompt(attributes interface{}, prompts Prompt) {
	for _, prompt := range prompts {
		prompt.AssignValue(attributes)
	}
}
