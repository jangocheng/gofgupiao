package main

// 2019.02.22
import (
	"fmt"
	"github.com/itzujun/gofgupiao/controller"
)

func main() {
	url := "http://quote.eastmoney.com/stocklist.html"
	fmt.Println("url:", url)
	ctrl := controller.NewController(url)
	fmt.Print(ctrl)
	ctrl.Go()
}
