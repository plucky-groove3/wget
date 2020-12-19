/*
Copyright © 2020 Lucas Hsueh <lucas_hsueh@hotmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package wget

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	bufferSize = 1024 * 8
)

func getResp(url string) *http.Response {
	cli := &http.Client{Transport: new(http.Transport)}
	resp, err := cli.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}

func writeFile(filename string, resp *http.Response) bool {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err = io.Copy(bufio.NewWriterSize(file, bufferSize), resp.Body); err != nil {
		panic(err)
	}
	return true
}

func Download(url, filename string, output string) bool {
	if url == "" {
		fmt.Println("param 'url' should not be empty.")
		return false
	}
	if filename == "" {
		filename = filepath.Base(url)
	}
	writeFile(filepath.Join(output, filename), getResp(url))
	return true
}
