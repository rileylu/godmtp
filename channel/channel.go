package channel

import "time"

type Channel struct {
	host string
	port string
	user string
	pwd  string
	dir string
}

type File struct {
	Name string
	Size uint64
	Time time.Time
}

type ChannelI interface {
	GetFileList() ([]*File, error)
}
