package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"sync"

	"golang.org/x/sync/singleflight"
)

var errorNotExist = errors.New("not exist")

var n int

func init() {
	flag.IntVar(&n, "n", 20, "模拟的并发数，默认 5")
}

func main() {
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(n)

	// 模拟并发访问
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			// 假设都获取 id = 1 这篇文章
			article := fetchArticle(1)
			log.Println(article)
		}()
	}
	wg.Wait()
}

type Article struct {
	ID      int
	Content string
}

var g singleflight.Group

func fetchArticle(id int) *Article {
	article := findArticleFromCache(id)

	if article != nil && article.ID > 0 {
		return article
	}

	// Do 执行并返回给定函数的结果，确保一次只对给定键执行一次。
	// 如果传入了一个副本，则该副本调用者等待原始调用完成并接收相同的结果。
	// 返回值 shared 指示是否将 v 赋给了多个调用者。
	v, err, shared := g.Do(strconv.Itoa(id), func() (interface{}, error) {
		return findArticleFromDB(id), nil
	})

	// 打印 shared，看看都什么值
	fmt.Println("shared===", shared)

	if err != nil {
		log.Println("singleflight do error:", err)
		return nil
	}

	return v.(*Article)
}

var (
	cache   = make(map[int]*Article)
	rwmutex sync.RWMutex
)

// 模拟从缓存获取数据
func findArticleFromCache(id int) *Article {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	return cache[id]
}

// 模拟从数据库中获取数据
func findArticleFromDB(id int) *Article {
	log.Printf("SELECT * FROM article WHERE id=%d", id)
	article := &Article{ID: id, Content: "Hello World"}
	rwmutex.Lock()
	defer rwmutex.Unlock()
	cache[id] = article
	return article
}
