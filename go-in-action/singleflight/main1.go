package main

// import (
// 	"errors"
// 	"flag"
// 	"log"
// 	"sync"
// )

// var errorNotExist = errors.New("not exist")

// var n int

// func init() {
// 	flag.IntVar(&n, "n", 20, "模拟的并发数，默认 5")
// }

// func main() {
// 	flag.Parse()

// 	var wg sync.WaitGroup
// 	wg.Add(n)

// 	// 模拟并发访问
// 	for i := 0; i < n; i++ {
// 		go func() {
// 			defer wg.Done()
// 			// 假设都获取 id = 1 这篇文章
// 			article := fetchArticle(1)
// 			log.Println(article)
// 		}()
// 	}
// 	wg.Wait()
// }

// type Article struct {
// 	ID      int
// 	Content string
// }

// func fetchArticle(id int) *Article {
// 	article := findArticleFromCache(id)

// 	if article != nil && article.ID > 0 {
// 		return article
// 	}

// 	return findArticleFromDB(id)
// }

// var (
// 	cache   = make(map[int]*Article)
// 	rwmutex sync.RWMutex
// )

// // 模拟从缓存获取数据
// func findArticleFromCache(id int) *Article {
// 	rwmutex.RLock()
// 	defer rwmutex.RUnlock()
// 	return cache[id]
// }

// // 模拟从数据库中获取数据
// func findArticleFromDB(id int) *Article {
// 	log.Printf("SELECT * FROM article WHERE id=%d", id)
// 	article := &Article{ID: id, Content: "polarisxu"}
// 	rwmutex.Lock()
// 	defer rwmutex.Unlock()
// 	cache[id] = article
// 	return article
// }
