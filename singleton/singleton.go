package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

// 懒汉式
var sLazy *singleton
var mu sync.Mutex

func getLazyInstance() *singleton {
	if sLazy == nil {
		mu.Lock()
		defer mu.Unlock()
		if sLazy == nil {
			sLazy = &singleton{}
		}
	}
	return sLazy
}

// 饿汉式
var sHungry *singleton = &singleton{}

func getHungryInstance() *singleton {
	return sHungry
}

// sync.Once
var sSync *singleton
var once sync.Once

func getSyncInstance() *singleton {
	once.Do(func() {
		sSync = &singleton{}
	})
	return sSync
}

func main() {
	s1 := getLazyInstance()
	s2 := getLazyInstance()
	s3 := getHungryInstance()
	s4 := getHungryInstance()
	s5 := getHungryInstance()
	s6 := getHungryInstance()
	fmt.Printf("%T, %T, %T, %T, %T, %T\n", s1, s2, s3, s4, s5, s6)
	fmt.Println(s1 == s2) // true
	fmt.Println(s3 == s4) // true
	fmt.Println(s5 == s6) // true
	fmt.Println(s1 == s3) // false
	fmt.Println(s1 == s5) // false
}
