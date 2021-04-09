package Requester

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSimpleFetch(t *testing.T) {
	type args struct {
		method string
		url    string
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
			},
			want:    []byte{0012},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SimplePcFetch(tt.args.method, tt.args.url)
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
		HTML, _ := SimplePcFetch(`GET`, `https://ssr1.scrape.center`)
		fmt.Println(string(HTML))
	}
}

func benchmarkSimplePcFetch(b *testing.B, page int) {
	for i := 1; i < b.N; i++ {
		for i := 0; i <= page; i++ {
			HTML, _ := SimplePcFetch(`GET`, `https://ssr1.scrape.center/page/`+strconv.Itoa(i))
			fmt.Println(string(HTML))
		}
	}
}

func BenchmarkSimplePcFetch1(b *testing.B) { benchmarkSimplePcFetch(b, 1) }

func BenchmarkSimplePcFetch2(b *testing.B) { benchmarkSimplePcFetch(b, 2) }

func BenchmarkSimplePcFetch5(b *testing.B) { benchmarkSimplePcFetch(b, 5) }

func BenchmarkSimplePcFetch10(b *testing.B) { benchmarkSimplePcFetch(b, 10) }
