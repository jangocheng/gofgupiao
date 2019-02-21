package controller

import (
	"github.com/itzujun/gofgupiao/basic"
	"github.com/itzujun/gofgupiao/downloader"
	"sync"
)

var wg sync.WaitGroup

var logger basic.ConsoleLogger

type Controller struct {
	Downloader downloader.Downloader // 下载器
	Analyzer
}