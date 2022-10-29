package loggo

import (
	"encoding/json"
	"fmt"
	"strings"
)

type OutConsole struct {
	Data JSONLogFormat
}

// OutConsole - Console function for write data
func (cl *OutConsole) OutData() {
	head := fmt.Sprintf("%s%s%s", BOLDText, strings.ToUpper(cl.Data.Level), COLORReset)
	switch cl.Data.Level {
	case "info":
		{
			head = fmt.Sprintf("%s%s", COLORWhite, head)
			break
		}
	case "trace", "debug":
		{
			head = fmt.Sprintf("%s%s", COLORCyan, head)
			break
		}
	case "warning":
		{
			head = fmt.Sprintf("%s%s", COLORYellow, head)
			break
		}
	case "fatal", "error", "panic":
		{
			head = fmt.Sprintf("%s%s", COLORRed, head)
			break
		}
	default:
		{
			head = fmt.Sprintf("%s%s", COLORPurple, head)
		}
	}
	fmt.Printf("%s\t[%s] \t%s\n", head, cl.Data.Time, cl.Data.Mgs)
}

func (ou *OutConsole) Write(p []byte) (n int, err error) {
	err = json.Unmarshal(p, &ou.Data)
	if err != nil {
		return 0, err
	}
	ou.OutData()
	return len(p), nil
}
