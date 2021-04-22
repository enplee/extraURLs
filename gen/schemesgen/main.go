package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

// 从iana拉取最新的url访问策略,生成对应的go-> []string
// https://www.iana.org/assignments/uri-schemes/uri-schemes-1.csv

const (
	path = "schemes.go" // gen go flie
	source = "https://www.iana.org/assignments/uri-schemes/uri-schemes-1.csv" // schemes source
)

var schemesTmpl = template.Must(template.New("schemes").Parse(`// gen by gen/schemesgen

package extraURLs

// schemes是由IANA认证的所有访问策略的有序列表
// IANA资源地址：https://www.iana.org/assignments/uri-schemes/uri-schemes-1.csv

var Schemes = []string{
{{range $scheme := .Schemes}}` + "\t`" + `{{$scheme}}` + "`" + `,
{{end}}}`))

func getSchemes() []string {
	resp, err := http.Get(source)
	if err != nil {
		log.Fatal(err) // pull fail -> fatal
	}
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	schemes := make([]string,0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		schemes = append(schemes, record[0])
	}
	return schemes
}

func genSchemesFile(schemes []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return schemesTmpl.Execute(file, struct {
		Schemes []string
	}{
		schemes,
	})
}

func main() {
	schemes := getSchemes()
	fmt.Printf("Generating %s...", path)
	if err := genSchemesFile(schemes); err != nil {
		log.Fatalf("write path: %v fail",err)
	}
	fmt.Println("gene sucss")
}