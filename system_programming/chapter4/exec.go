package chapter4

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func Run() {
	{
		//// Q4.1
		//signals := make(chan os.Signal, 1)
		//signal.Notify(signals, syscall.SIGINT)
		//scanner := bufio.NewScanner(os.Stdin)
		//scanner.Buffer(make([]byte, 2), 1)
		//fmt.Println("Please enter a number. (n < 10)")
		//go func() {
		//	for scanner.Scan() {
		//		n, err := strconv.Atoi(scanner.Text())
		//		if err != nil {
		//			fmt.Println("Please enter a number. (n < 10)")
		//			continue
		//		}
		//		<-time.After(time.Second * time.Duration(n))
		//		fmt.Println(n, "seconds passed.")
		//	}
		//	if err := scanner.Err(); err != nil {
		//		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		//		signals <- syscall.SIGINT
		//	}
		//}()
		//<-signals
		//fmt.Println("SIGINT arrived!")

		//// Q4.1 with no use of scanner
		r := bufio.NewReader(os.Stdin)
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT)
		go func() {
			for {
				fmt.Println("Please enter a number. (n < 10)")
				l, err := r.ReadString('\n')
				if err != nil {
					panic(err)
				}
				n, err := strconv.Atoi(l[:len(l)-1])
				if err != nil || n >= 10 {
					continue
				}
				fmt.Println("waiting", n, "seconds...")
				<-time.After(time.Second * time.Duration(n))
				fmt.Println(n, "seconds passed.")
			}
		}()
		<-signals
		fmt.Println("SIGINT catched.")
	}

}
func DryRun() {

}
