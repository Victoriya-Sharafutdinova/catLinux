package main

import (
	"bufio"
	"cat/logic"
	"cat/logic/LineHandlers"
	"flag"
	"log"
	"os"
)

func main() {

	numberOnlyNonEmptyLines := flag.Bool("b", false, "нумеровать только непустые строки")
	showEOL := flag.Bool("E", false, "показывать символ $ в конце каждой строки")
	numberAllLines := flag.Bool("n", false, "нумеровать все строки")
	deleteEmptyDuplicateLines := flag.Bool("s", false, "удалять пустые повторяющиеся строки")
	displayTabs := flag.Bool("T", false, "отображать табуляции в виде ^I")

	flag.Parse()

	if len(os.Args) <= 1 {
		log.Fatalf("Ошибка. Слишком мало аргументов!")
	}

	fileName := os.Args[len(os.Args)-1]
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Ошибка открытия файла: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var handlers []LineHandlers.IHandler

	if *numberAllLines && *numberOnlyNonEmptyLines {
		*numberAllLines = false
	}

	if *deleteEmptyDuplicateLines {
		handlers = append(handlers, LineHandlers.NewDuplicateHandler())
	}
	if *numberOnlyNonEmptyLines {
		handlers = append(handlers, LineHandlers.NewNumberHandler(false))
	}
	if *numberAllLines {
		handlers = append(handlers, LineHandlers.NewNumberHandler(true))
	}
	if *displayTabs {
		handlers = append(handlers, LineHandlers.NewTabHandler())
	}
	if *showEOL {
		handlers = append(handlers, LineHandlers.NewEolHandler())
	}

	conv := logic.NewConveyor(handlers)
	conv.Run(*scanner)
}
