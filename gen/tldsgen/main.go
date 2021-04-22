package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"sync"
)
// 分别从iana和pubsuffix两处拉取顶级域名名单---> tlds []string
// 涉及到去重和对拉取的内容进行过滤(两处资源中包含注释等其他信息) -->  map[string]bool  reg
const (
	path = "tlds.go"
	source_iana = "https://data.iana.org/TLD/tlds-alpha-by-domain.txt"
	source_pubsuffix = "https://publicsuffix.org/list/effective_tld_names.dat"
	iana_pat = `^[^#]+$`
	pubsuffix_pat = `^[^/.]+$`
)

func handleTld(tld string) string {
	tld = strings.ToLower(tld)
	if strings.HasPrefix(tld,"xn--") {
		return ""
	}
	return tld
}

func fetchFromURL(url,pat string,tldSet map[string]bool) {
	resp, err := http.Get(url)
	if err == nil && resp.StatusCode >= 400{
		err = errors.New(resp.Status)
	}
	if err != nil {
		panic(fmt.Errorf("%s,%s",url,err))
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	re := regexp.MustCompile(pat)
	for scanner.Scan() {
		line:= scanner.Text()
		tld := re.FindString(line)
		handleTld(tld)		//对tld进行小写,并过滤xn--开头的域名
		if tld == "" {
			continue
		}
		tldSet[tld] = true // add set
	}
	if err := scanner.Err();err != nil {
		panic(fmt.Errorf("%s,%s", url, err))
	}
}
func getTlds()(urls []string,tlds []string){
	tldSet := make(map[string]bool) // set
	var wg sync.WaitGroup
	fetchWorker := func(url,pat string) {
		urls = append(urls,url)
		wg.Add(1)
		go fetchFromURL(url,pat,tldSet)
	}
	fetchWorker(source_iana,iana_pat)
	fetchWorker(source_pubsuffix,pubsuffix_pat)
	wg.Wait() // wait worker finish and do operate

	tlds = make([]string,0,len(tldSet))
	for tld := range tldSet {
		tlds = append(tlds, tld)
	}
	sort.Strings(tlds)
	return urls, tlds
}
func main() {
	tldSet:= make(map[string]bool)
	fetchFromURL(source_iana,iana_pat,tldSet)
	for tld := range tldSet {
		fmt.Println(tld)
	}
}