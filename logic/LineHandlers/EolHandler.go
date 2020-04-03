package LineHandlers

type EolHandler struct{}

func (self *EolHandler) Handle(line string, printLine bool) (string, bool) {
	if !printLine {
		return line, printLine
	}

	return line + "$", printLine
}

func NewEolHandler() *EolHandler {
	return new(EolHandler)
}
