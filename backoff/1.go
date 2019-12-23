package main

import (
	"fmt"
	"time"

	"errors"

	"github.com/cenkalti/backoff"
)

func Test1() {
	// An operation that may fail.
	operation := func() error {
		return errors.New("haha") // or an error
	}

	b := backoff.NewExponentialBackOff()
	b.InitialInterval = time.Second

	err := backoff.RetryNotify(operation, b, func(err error, duration time.Duration) {
		fmt.Println("aaaaaaaaaaa", err, duration)
	})
	if err != nil {
		// Handle error.
		return
	}

	fmt.Println("222222222222")
	// Operation is successful.
}
func main() {

	Test1()

}

/*
aaaaaaaaaaa haha 1.104660288s
aaaaaaaaaaa haha 2.160763633s
aaaaaaaaaaa haha 2.62026012s
aaaaaaaaaaa haha 3.164785382s
aaaaaaaaaaa haha 4.680977329s
aaaaaaaaaaa haha 9.01243771s
aaaaaaaaaaa haha 6.442959172s
aaaaaaaaaaa haha 11.217246954s
aaaaaaaaaaa haha 15.299675834s
aaaaaaaaaaa haha 30.789742484s
aaaaaaaaaaa haha 58.542275879s
aaaaaaaaaaa haha 1m18.81839766s

*/
