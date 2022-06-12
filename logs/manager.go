package logs

import (
	"log"
	"os"
	"runtime"
)

const (
	errPrefix   = "[ERROR]"
	infoPrefix  = "[INFO]"
	warnPrefix  = "[WARN]"
	debugPrefix = "[DEBUG]"
	tracePrefix = "[TRACE]"
)

var M *Manager

func init() {
	M = New()
}

type Manager struct {
	logger *log.Logger
}

func New() *Manager {
	return &Manager{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (m *Manager) PrintInfo(v ...any) {
	m.logger.Println(Sprint(Green, append([]any{m.runFuncName(), infoPrefix}, v...)...))
}

func (m *Manager) PrintError(v ...any) {
	m.logger.Println(Sprint(Red, append([]any{m.runFuncName(), errPrefix}, v...)...))
}

func (m *Manager) FatallnError(v ...any) {
	m.logger.Fatalln(Sprint(Red, append([]any{m.runFuncName(), errPrefix}, v...)...))
}

func (m *Manager) PrintWarn(v ...any) {
	m.logger.Println(Sprint(Yellow, append([]any{m.runFuncName(), warnPrefix}, v...)...))
}

func (m *Manager) PrintDebug(v ...any) {
	m.logger.Println(Sprint(Cyan, append([]any{m.runFuncName(), debugPrefix}, v...)...))
}

func (m *Manager) PrintTrace(v ...any) {
	m.logger.Println(Sprint(White, append([]any{m.runFuncName(), tracePrefix}, v...)...))
}

func (m *Manager) runFuncName() string {
	pc := make([]uintptr, 2)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[1])
	return f.Name()
}
