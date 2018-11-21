package channel

import (
	"fmt"
	"testing"
)

func TestFtpChannel_GetFileList(t *testing.T) {
	ftpChannel := &FtpChannel{}
	ftpChannel.host = "wmsappqa"
	ftpChannel.port = "21"
	ftpChannel.user = "dmtpftpq"
	ftpChannel.pwd = "vmM6sp47"
	ftpChannel.dir = "/"
	list, err := ftpChannel.GetFileList()
	if err != nil {
		return
	}
	fmt.Println(list)
}
