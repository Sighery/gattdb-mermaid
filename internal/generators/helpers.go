package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/goccy/go-yaml"
)

type envelope[T any] struct {
	Etag string
	Data []T
}

func fileEtag(filepath string) (etag string) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening file: %w\n", err)
		return etag
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	prefix := "// ETag: "

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, prefix) {
			continue
		}

		etag, _ = strings.CutPrefix(line, prefix)
		etag = strings.TrimSpace(etag)
		return etag
	}

	return etag
}

func fetchData[T any](etag, url, yamlPath string) (envelope envelope[T], changed bool) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating new request: %w", err)
	}

	// For some reason the manufacturers page has no ETag when using gzip encoding?
	req.Header.Set("Accept-Encoding", "identity")
	req.Header.Add("If-None-Match", etag)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error fetching data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return envelope, false
	}

	path, err := yaml.PathString(yamlPath)
	if err != nil {
		log.Fatalf("Error creating YAML path: %w", err)
	}

	if err := path.Read(resp.Body, &envelope.Data); err != nil {
		log.Fatalf("Error parsing YAML: %w", err)
	}

	envelope.Etag = resp.Header.Get("ETag")
	return envelope, true
}

func writeTemplate[T any](templateStr string, data envelope[T], wr io.Writer) {
	funcMap := template.FuncMap{
		// No native Go string escaping function. Only real worry is a sub double quote
		"escape": func(s string) string { return strings.ReplaceAll(s, `"`, `\"`) },
	}

	tmpl := template.Must(template.New("t").Funcs(funcMap).Parse(templateStr))

	if err := tmpl.Execute(wr, data); err != nil {
		log.Fatalf("Error writing template: %v", err)
	}
}
