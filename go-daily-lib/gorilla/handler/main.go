package main

// https://mp.weixin.qq.com/s/0gWmwOf2hhA-N3FJWCrQ7A
import (
	"bytes"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Hello World")
}

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "login")
}

func doLogin(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "login")
}

type greeting string

func (g greeting)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome, %s", g)
}

func Logging(handler http.Handler) http.Handler {
	//return handlers.CombinedLoggingHandler(os.Stdout, handler)
	return handlers.CustomLoggingHandler(os.Stdout, handler, myLogFormatter)
}

func myLogFormatter(writer io.Writer, params handlers.LogFormatterParams) {
	var buf bytes.Buffer
	buf.WriteString(time.Now().Format("2006-01-02 15:04:05 -0700"))
	buf.WriteString(fmt.Sprintf(` "%s %s %s" `, params.Request.Method, params.URL.Path, params.Request.Proto))
	buf.WriteString(strconv.Itoa(params.StatusCode))
	buf.WriteByte('\n')

	writer.Write(buf.Bytes())
}

func main() {
	r := mux.NewRouter()
	//r.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(index)))
	r.Use(Logging)
	r.Use(handlers.CompressHandler)
	//r.HandleFunc("/", index)
	// 压缩
	//r.Handle("/greeting", handlers.CombinedLoggingHandler(os.Stdout, greeting("dj")))
	//r.Methods("GET").Path("/login").HandlerFunc(login)
	// Content-Type处理
	//r.Methods("POST").Path("/login").Handler(handlers.ContentTypeHandler(http.HandlerFunc(doLogin), "application/x-www-form-urlencoded"))

	//方法分发器
	r.Handle("/login", handlers.MethodHandler{
		"GET":  http.HandlerFunc(login),
		"POST": http.HandlerFunc(doLogin),
	})

	// 重定向
	r.Use(handlers.CanonicalHost("http://www.gorillatoolkit.org", 302))

	// Recovery
	r.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}