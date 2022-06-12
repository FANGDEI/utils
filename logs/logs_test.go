package logs

import "testing"

func TestLogs(t *testing.T) {
	M.PrintDebug("DEBUG INFORMATION")
	M.PrintError("ERROR INFORMATION")
	M.PrintInfo("INFO INFORMATION")
	M.PrintTrace("TRACE INFORMATION")
	M.PrintWarn("WARN INFORMATION")
}
