package LineHandlers

import "strings"

type TabHandler struct{}

func (self *TabHandler) Handle(line string, printLine bool) (string, bool) {
	if !printLine {
		return line, printLine
	}

	return strings.ReplaceAll(line, "	", "^I"), printLine
}

func NewTabHandler() *TabHandler {
	return new(TabHandler)
}
