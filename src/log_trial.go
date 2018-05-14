package main

import "fmt"

type Logger struct {
	orderType        string
	debugEnabled     *bool
}

func MakeLogger() *Logger {
	debug := false
	return &Logger{
		orderType: "type",
		debugEnabled: &debug,
	}
}

func (l *Logger) UpdateDebug(debugFlag bool)  {
	*l.debugEnabled = debugFlag
}

func (l *Logger) MakeSubLogger() *Logger {
	return &Logger{
		orderType: l.orderType,
		debugEnabled: l.debugEnabled,
	}
}

func main() {
	log := MakeLogger()
	log.UpdateDebug(true)
	fmt.Println(log)
	log2 := log.MakeSubLogger()
	log.UpdateDebug(false)
	fmt.Println(*log.debugEnabled, *log2.debugEnabled)
	log2.UpdateDebug(true)
	fmt.Println(*log.debugEnabled, *log2.debugEnabled)
	a := 10
	b := &a
	c := b
	fmt.Println(a, b, c)
	fmt.Println(a, *b, *c)
	a = 20
	fmt.Println(a, *b, *c)
}
