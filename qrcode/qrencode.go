package qrcode

import (
	"bytes"
	"github.com/qpliu/qrencode-go/qrencode"
	"os"
)

func QrencodeTerminal(strText string, level qrencode.ECLevel) {
	var buf bytes.Buffer
	buf.WriteString(strText)
	grid, err := qrencode.Encode(buf.String(), level)
	if err != nil {
		panic(err)
	}
	grid.Image(0)
	grid.TerminalOutput(os.Stdout)
}
