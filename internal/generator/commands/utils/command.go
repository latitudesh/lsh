package commands

type Command struct {
	Name       string
	Short      string
	Long       string
	Method     string
	Root       string
	Parameters []CmdParameter
}

type CmdParameter struct {
	Name        string
	Description string
	Required    bool
}
