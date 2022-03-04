package FakeHeader

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestRandomUserAgent(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest(`GET`, `https://www.baidu.com`, nil)
	if err != nil {
		log.Printf("NewRequest Error %#v", err)
		panic(err)
	}
	SetRandomPcUserAgent(req)
	res, err := client.Do(req)
	if err != nil {
		log.Printf("client Error %#v", err)
		panic(err)
	}
	defer res.Body.Close()
	//all, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(all))
	fmt.Println(res.Request.Header)

}

func BenchmarkRandomUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//client := &http.Client{}
		//req, err := http.NewRequest(`GET`, `https://www.baidu.com`, nil)
		//if err != nil {
		//	log.Printf("NewRequest Error %#v", err)
		//	panic(err)
		//}
		//RandomAllUserAgent(req)
		//res, err := client.Do(req)
		//if err != nil {
		//	log.Printf("client Error %#v", err)
		//	panic(err)
		//}
		//defer res.Body.Close()
		////all, err := ioutil.ReadAll(res.Body)
		////fmt.Println(string(all))
		//fmt.Println(res.Request.Header)
		fmt.Println(RandomPcUserAgent())
		fmt.Println(RandomAllUserAgent())
		fmt.Println(RandomMobileUserAgent())

	}
}

func benchmarkRandomUserAgent(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < n; i++ {
			fmt.Println(RandomPcUserAgent())
			fmt.Println(RandomAllUserAgent())
			fmt.Println(RandomMobileUserAgent())
		}
	}
}

func BenchmarkGroup1(b *testing.B)   { benchmarkRandomUserAgent(b, 1) }
func BenchmarkGroup5(b *testing.B)   { benchmarkRandomUserAgent(b, 5) }
func BenchmarkGroup10(b *testing.B)  { benchmarkRandomUserAgent(b, 10) }
func BenchmarkGroup20(b *testing.B)  { benchmarkRandomUserAgent(b, 20) }
func BenchmarkGroup50(b *testing.B)  { benchmarkRandomUserAgent(b, 50) }
func BenchmarkGroup100(b *testing.B) { benchmarkRandomUserAgent(b, 100) }

//https://github.com/gowebspider/utensils.git
