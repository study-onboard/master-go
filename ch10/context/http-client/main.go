package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}

	url := os.Args[1]

	response, err := request(url)
	if err != nil {
		fmt.Printf("request url %s failed: %s\n", url, err.Error())
		os.Exit(1)
	}

	fmt.Println(response)
}

// request url
//
// return:
//
// ----------------------------------------------
//
// result - text content of response
//
// err - error
func request(url string) (result string, err error) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(30)*time.Second)

	// when request function return, call cancel function to cancel request,
	// if request is success, cancel function will do nothing for this operation.
	defer cancel()

	data := make(chan *requestResponse)

	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}
	go func() {
		req, _ := http.NewRequest("GET", url, nil)
		response, err := httpClient.Do(req)
		if err != nil {
			data <- &requestResponse{
				result: "",
				err:    err,
			}
		} else {
			b, _ := ioutil.ReadAll(response.Body)
			data <- &requestResponse{
				result: string(b),
				err:    nil,
			}
		}
	}()

	select {
	// when request timeout, ctx will done, this channel will close
	case <-ctx.Done():
		result = ""
		err = ctx.Err()
		return

	// request data return
	case response := <-data:
		result = response.result
		err = response.err
		return
	}
}

type requestResponse struct {
	result string
	err    error
}
