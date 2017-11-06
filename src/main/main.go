package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
//	"runtime"
	"time"
)

func main() {
//	runtime.GOMAXPROCS(4)

	start := time.Now()

	count := 25

	for i := 0; i < count; i++ {
//		go func() {

			resp, err := http.Get("https://api.chucknorris.io/jokes/random")

			if err == nil {
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)

				joke := new(ChuckNorrisJoke)
				json.Unmarshal(body, &joke)

				fmt.Printf("id: %s  joke : %s\n", joke.Id, joke.Value)
			} else {
				fmt.Println(err)
			}
//			count--
//		}()
	}

//	for count > 0 {
//		time.Sleep(10 * time.Microsecond)
//	}

	elapsed := time.Since(start)

	fmt.Printf("Execution Time: %s", elapsed)
}

type ChuckNorrisJoke struct {
	Id       string
	Category string
	Icon_url string
	Url      string
	Value    string
}
