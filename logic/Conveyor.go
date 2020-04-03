package logic

import (
	"bufio"
	"cat/logic/LineHandlers"
	"fmt"
)

type Conveyor struct {
	handlers []LineHandlers.IHandler
}

func (self *Conveyor) Run(scanner bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		printLine := true

		for _, h := range self.handlers {
			line, printLine = h.Handle(line, printLine)
		}

		if !printLine {
			continue
		}
		fmt.Println(line)
	}
}

func NewConveyor(handlers []LineHandlers.IHandler) *Conveyor {
	c := new(Conveyor)
	c.handlers = handlers
	return c
}
