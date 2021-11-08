package main

// https://mp.weixin.qq.com/s/lhuv27BuaX-J0gcKyJC0Bw
import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//r := mux.NewRouter()
	//r.HandleFunc("/", BooksHandler)
	//r.HandleFunc("/books/{isbn}", BookHandler).Host("bookstore.com").Methods("GET").Schemes("HTTP")
	//r.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
	//	return r.ProtoMajor == 1 && r.ProtoMinor == 1
	//})
	//http.Handle("/", r)
	//log.Fatal(http.ListenAndServe(":8080", nil))


	//r := mux.NewRouter()
	//bs := r.PathPrefix("/books").Subrouter()
	//bs.HandleFunc("/", BooksHandler)
	//bs.HandleFunc("/books/{isbn}", BookHandler)
	//
	//ms := r.PathPrefix("/movies").Subrouter()
	//ms.HandleFunc("/", BooksHandler)
	//ms.HandleFunc("/{imdb}", BookHandler)
	//http.Handle("/", r)
	//log.Fatal(http.ListenAndServe(":8080", nil))

	r := mux.NewRouter()
	r.Host("github.io")
	r.Host("{subdomain:[a-zA-Z0-9]+}.github.io")
	r.PathPrefix("/books/") // 只处理路径前缀为`/books/`的请求
	r.Methods("GET", "POST") // 只处理 GET/POST 请求
	r.Schemes("https") // 只处理 https 的请求
	// 只处理首部 X-Requested-With 的值为 XMLHTTPRequest 的请求
	r.Headers("X-Requested-With", "XMLHTTPRequest")
	// 只处理查询参数包含key=value的请求
	r.Queries("key", "value")
	// 自定义匹配器
	r.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		return r.ProtoMajor == 1 && r.ProtoMinor == 1
	})

	// 为一个路由起一个名字
	r.HandleFunc("/books/{isbn}", BookHandler).Name("book")

	r.Use(loggingMiddleware)
	// /books/978-7-111-55842-2 <nil>
	fmt.Println(r.Get("book").URL("isbn", "978-7-111-55842-2"))

	// 创建子路由
	InitBooksRouter(r)
	InitMoviesRouter(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
