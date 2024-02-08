package operation

type Operation interface {
	RegisterFlags()
	AssignID(bodyData interface{}) error
	AssignBodyAttributes(attributes interface{}) error
	PromptID(bodyData interface{})
	PromptAttributes(attributes interface{})
}
