package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
)

func startPublicServer(addr string) {
	http.Handle("/", http.FileServer(http.Dir("./public")))

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Fatal(new(http.Server).Serve(l))
	}()
}

func startHugoServer(addr string) {
	panic("not implemented yet")
}

func checkURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("status code %d", resp.StatusCode)
	}
	return nil
}

func main() {
	log.SetFlags(0)

	addrF := flag.String("addr", "127.0.0.1:1313", "Listen address")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: go run check-urls.go [ public | hugo | golangshow.com ]\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	baseURL := &url.URL{
		Scheme: "http",
		Host:   *addrF,
		Path:   "/",
	}
	switch arg := flag.Arg(0); arg {
	case "public":
		startPublicServer(*addrF)
	case "hugo":
		startHugoServer(*addrF)
	case "golangshow.com":
		baseURL.Host = "golangshow.com"
	default:
		log.Printf("Unexpected argument %q.", arg)
	}
	log.Printf("Serving on %s ...", baseURL)

	for _, path := range []string{
		"/episode/2015/06-11-001",
		"/episode/2015/06-11-001/",
		"/page/13",
		"/page/13/",
		"/categories/гости", // guests
		"/categories/гости/",
		"/categories/интервью", // interview
		"/categories/интервью/",
		"/categories/итоги", // results
		"/categories/итоги/",
		"/categories/камин", // fireplace
		"/categories/камин/",
		"/categories/новости", // news
		"/categories/новости/",
	} {
		baseURL.Path = path
		u := baseURL.String()
		log.Printf("Checking %s (%s) ...", path, u)
		if err := checkURL(u); err != nil {
			log.Fatalf("%s (%s): %s", path, u, err)
		}
	}
}
