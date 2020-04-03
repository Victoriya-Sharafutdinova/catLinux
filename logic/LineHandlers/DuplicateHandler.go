package LineHandlers

type DuplicateHandler struct {
	lastLineIsEmpty bool
}

func (self *DuplicateHandler) Handle(line string, printLine bool) (string, bool) {
	if !printLine {
		return line, printLine
	}

	printLine = !(self.lastLineIsEmpty && line == "")
	self.lastLineIsEmpty = line == ""

	return line, printLine
}

func NewDuplicateHandler() *DuplicateHandler {
	h := new(DuplicateHandler)
	h.lastLineIsEmpty = false
	return h
}
