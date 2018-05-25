package archiver

import (
	"fmt"

	"github.com/mholt/archiver"
)

func Run() {
	err := archiver.Unarchive("/Users/liuwei/Downloads/1.zip", "_test_archiver")
	fmt.Println(err)
}
