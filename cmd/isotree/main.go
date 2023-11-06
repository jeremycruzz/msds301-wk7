package main

import (
	"fmt"

	"github.com/jeremycruzz/msds301-wk7/internal/mnist"
)

func main() {
	mnist := mnist.NewMnist()
	fmt.Println("loading data..")
	mnist.LoadData()
	fmt.Println("doing isolation forest..")
	mnist.IsolationForest()
	fmt.Println("writing to file..")
	mnist.WriteCsv()
	fmt.Println("done.")
}
