package main

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	port = flag.String("port", "8080", "Define what TCP port to bind to")
	root = flag.String("root", "root", "Define the root filesystem path")
)

// IsExist 判断path是否存在
func creatPath(path string) error {

	fi, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(path, 0600)
		}
		return err
	}

	if fi.IsDir() {
		return nil
	} else {
		return fmt.Errorf("root isn't path:" + path)
	}

}

func getPath(name string) string {
	return *root + "/" + name
}

func getName(URLPath string) string {
	name := strings.ReplaceAll(URLPath, "\\", "")
	name = strings.ReplaceAll(name, "/", "")
	return name
}

var (
	name2md5 = make(map[string]string)
	nameM    sync.Mutex
)

type HttpMain struct {
}

func (hm *HttpMain) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	Method := strings.ToUpper(r.Method)
	if Method == "POST" {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		name := getName(r.URL.Path)
		nameM.Lock()
		defer nameM.Unlock()

		f, err := os.OpenFile(getPath(name), os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600) //0644
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		if _, err = f.Write(b); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		md5 := fmt.Sprintf("%x", md5.Sum(b))
		name2md5[name] = md5

		return

	} else if Method == "GET" {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		name := getName(r.URL.Path)
		nameM.Lock()
		defer nameM.Unlock()

		rammd5, ok := name2md5[name]
		if !ok {
			http.Error(w, "not found name("+name+") in map", http.StatusNotFound)
			return
		}

		md5 := strings.ToLower(string(b))
		if md5 == rammd5 {
			http.Error(w, "md5 is same", http.StatusForbidden)
			return
		}

		b, err = ioutil.ReadFile(getPath(name))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(name))
		// w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

		io.Copy(w, bytes.NewReader(b))

		return

	} else if Method == "DELETE" {

		name := getName(r.URL.Path)
		nameM.Lock()
		defer nameM.Unlock()

		err := os.Remove(getPath(name))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		delete(name2md5, name)

		return

	} else {
		http.Error(w, "Not implemented "+r.Method, http.StatusNotImplemented)
		return
	}

}

func main() {
	flag.Parse()

	err := creatPath(*root)
	if err != nil {
		panic(err)
	}

	items, _ := ioutil.ReadDir(*root)
	for _, item := range items {
		if item.IsDir() {
			continue

		}
		name := item.Name()
		b, err := ioutil.ReadFile(getPath(name))
		if err != nil {
			panic(err)
		}
		md5 := fmt.Sprintf("%x", md5.Sum(b))
		println("[" + md5 + "<=" + name + "]")
		name2md5[name] = md5

	}

	hm := &HttpMain{}
	fs := http.NewServeMux()

	fs.Handle("/", hm)

	CORSHeaders := AllowedHeaders([]string{"Authorization", "Content-Type", "User-Agent"})
	CORSOrigins := AllowedOrigins([]string{"*"})
	CORSMethods := AllowedMethods([]string{"GET", "POST", "DELETE"})
	mux2 := CompressHandlerLevel(CORS(CORSHeaders, CORSOrigins, CORSMethods)(fs), gzip.BestCompression)

	log.Println("Starting web server at 0.0.0.0:" + *port)
	panic(http.ListenAndServe(":"+*port, mux2))
}
