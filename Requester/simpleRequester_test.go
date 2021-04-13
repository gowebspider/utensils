package Requester

import (
	"fmt"
	"io"
	"strconv"
	"testing"
)

func TestSimpleFetch(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "simpleExample",
			args: args{
				`GET`,
				`https://www.baidu.com`,
				nil,
			},
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SimplePcFetch(tt.args.method, tt.args.url, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(string(got))
		})
	}
}

func BenchmarkSimplePcFetch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimplePcFetch(`GET`, `https://ssr1.scrape.center`, nil)
		//fmt.Println(string(HTML))
	}
}

func benchmarkSimplePcFetch(b *testing.B, page int) {
	for i := 1; i < b.N; i++ {
		for i := 0; i <= page; i++ {
			SimplePcFetch(`GET`, `https://ssr1.scrape.center/page/`+strconv.Itoa(i), nil)
		}
	}
}

func BenchmarkSimplePcFetch1(b *testing.B) { benchmarkSimplePcFetch(b, 1) }

func BenchmarkSimplePcFetch2(b *testing.B) { benchmarkSimplePcFetch(b, 2) }

func BenchmarkSimplePcFetch5(b *testing.B) { benchmarkSimplePcFetch(b, 5) }

func BenchmarkSimplePcFetch10(b *testing.B) { benchmarkSimplePcFetch(b, 10) }

func BenchmarkSimplePcFetchParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 100; i++ {
				Body, _ := SimplePcFetch(`GET`, `https://www.baidu.com`, nil)
				fmt.Println(string(Body))
			}
		}
	})
}
