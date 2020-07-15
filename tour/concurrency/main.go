package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int, d time.Duration) {
	time.Sleep(d)
	fmt.Println("sum called", s)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c

}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// ここではこれをしないとデッドロックを検出してエラーになる
	// とはいえ、チャネルはファイルと違って必ずcloseしなければならないわけではない
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("waiting...")
		select {
		case c <- x:
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("sent %v to c\n", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {

		v1, ok1 := <-c1
		v2, ok2 := <-c2
		//fmt.Printf("v1: %v, ok1: %v, v2: %v, ok2: %v\n", v1, ok1, v2, ok2)
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if ok1 == false && ok2 == false {
			return true
		}
	}
}

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	// Inc() が goroutine として実行されるという前提があり、goroutine として一つだけが実行されている状態を作るためにやる
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func main() {
	if false {
		{
			// goroutine(Goのランタイムに管理される軽量なスレッド)
			go say("world")
			say("hello")
		}
		{
			s := []int{7, 2, 8, -9, 4, 0}
			// Channel 型の変数を作成(intの送受信用)
			c := make(chan int)
			go sum(s[:len(s)/2], c, 1000*time.Millisecond)
			// 当然↑よりも↓の方が、500ms早く実行される
			go sum(s[len(s)/2:], c, 500*time.Millisecond)
			time.Sleep(2000 * time.Millisecond)
			fmt.Println("start receiving")
			// 500msの方が早いのでそちらが最初に受信されてxに入る
			x, y := <-c, <-c // receive from c
			fmt.Println(x, y, x+y)
			//x := <-c
			//fmt.Println(x)
			//y := <-c
			//fmt.Println(y)
		}
		{
			ch := make(chan int, 2)
			ch <- 13
			ch <- 29
			// ↓バッファの長さを超えるのでエラー
			// ch <- 31
			fmt.Println(<-ch)
			ch <- 31
			fmt.Println(<-ch)
			v, ok := <-ch
			fmt.Println(v, ok)
			close(ch)
			v, ok = <-ch
			fmt.Println(v, ok)

		}
		{
			c := make(chan int, 10)
			go fibonacci(cap(c), c)
			// チャネルが閉じられるまで値を受信し続ける
			for i := range c {
				fmt.Println(i)
			}
		}
		{
			c := make(chan int)
			quit := make(chan int)
			go func() {
				for i := 0; i < 10; i++ {
					fmt.Println(<-c)
				}
				quit <- 0
			}()
			fibonacci2(c, quit)
		}

		{
			tick := time.Tick(100 * time.Millisecond)
			boom := time.After(500 * time.Millisecond)
			// このSleepを入れると、selectで複数のcaseに当てはまる時にランダムで選ばれることを実証できる
			// time.Sleep(1000 * time.Millisecond)
		L:
			for {
				select {
				case <-tick:
					fmt.Println("tick.")
				case <-boom:
					fmt.Println("BOOM!")
					// Labeled Break
					break L
				default:
					fmt.Println("   .")
					time.Sleep(50 * time.Millisecond)
				}
			}

		}
		{
			// Exercise: Equivalent Binary Trees
			fmt.Println(Same(tree.New(1), tree.New(1)))
			fmt.Println(Same(tree.New(4), tree.New(6)))
		}
		c := SafeCounter{v: make(map[string]int)}
		for i := 0; i < 1000; i++ {
			go c.Inc("hoge")
		}
		time.Sleep(time.Second)
		fmt.Println(c.v["hoge"])
	} // false end

	{
		// map の挙動チェック
		m := make(map[string]int)
		v, ok := m["hoge"]
		m["foo"] = 0
		v, ok = m["foo"]
		fmt.Println(v, ok)
		// Exercise: Web Crawler
		Crawl("https://golang.org/", 4, fetcher)
	}
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type CrawlResult struct {
	ok   bool
	body string
}

type SafeCrawler struct {
	v   map[string]CrawlResult
	mux sync.Mutex
	wg  sync.WaitGroup
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	c := SafeCrawler{v: make(map[string]CrawlResult)}
	c.crawl(url, depth, fetcher)
	c.wg.Wait()
	for url, res := range c.v {
		if res.ok {
			fmt.Printf("<found> url: %s, body: %s\n", url, res.body)
		} else {
			fmt.Printf("<not found> url: %s\n", url)
		}
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *SafeCrawler) crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		c.v[url] = CrawlResult{ok: false}
		return
	}
	//fmt.Printf("found: %s %q\n", url, body)
	for _, url := range urls {
		if _, ok := c.v[url]; ok {
			continue
		}
		c.mux.Lock()
		c.v[url] = CrawlResult{true, body}
		c.mux.Unlock()
		c.wg.Add(1)
		go c.crawl(url, depth-1, fetcher)
	}
	c.wg.Done()
	return
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		//fmt.Println(res.body, res.urls)
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
