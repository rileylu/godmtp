package channel

import (
	"github.com/jlaffaye/ftp"
)

type FtpChannel struct {
	Channel
}

func (ftpChannel *FtpChannel) GetFileList() (fileList []*File, err error) {
	ftpClient, err := ftp.Connect(ftpChannel.host + ":" + ftpChannel.port)
	defer ftpClient.Quit()
	if err != nil {
		return nil, err
	}
	err = ftpClient.Login(ftpChannel.user, ftpChannel.pwd)
	if err != nil {
		return nil, err
	}
	var fl []*ftp.Entry
	fl, err = ftpClient.List(ftpChannel.dir)
	if err != nil {
		return nil, err
	}
	var t *ftp.Entry
	for _, t = range fl {
		if t.Type == ftp.EntryTypeFile {
			newFile := &File{
				Name: t.Name,
				Size: t.Size,
				Time: t.Time,
			}
			fileList = append(fileList, newFile)
		}
	}
	return
}
