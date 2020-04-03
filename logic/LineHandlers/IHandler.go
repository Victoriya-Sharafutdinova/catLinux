package LineHandlers

type IHandler interface {
	Handle(line string, printLine bool) (string, bool)
}
