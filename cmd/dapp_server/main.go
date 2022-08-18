package main

import (
	"github.com/civet148/log"
)


func main() {
	//strWCURI := "wc:38d80e38-1c9a-470f-aed1-27b8ced3338d@1?bridge=https://6.bridge.walletconnect.org&key=d8089b579aeef1ba4692d72bd3a0bb870f8bbceafaab325f23739635e2e54f05"
	//qrcode.QrencodeTerminal(strWCURI, qrencode.ECLevelL)
	ch := make(chan interface{})
	<-ch
	log.Infof("program exit...")
}
