package main

import "fmt"

func main() {
	r := Constructor()
	fmt.Println(r.Ping(1))
	fmt.Println(r.Ping(100))
	fmt.Println(r.Ping(3001))
	fmt.Println(r.Ping(3002))
}

type RecentCounter struct {
	Request []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		Request: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.Request = append(this.Request, t)
	for this.Request[0] < t-3000 {
		this.Request = this.Request[1:]
	}
	return len(this.Request)
}
