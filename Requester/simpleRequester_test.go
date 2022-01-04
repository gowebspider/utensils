package Requester

import (
	"net/http"
	"reflect"
	"testing"
)

//func TestSimpleFetch(t *testing.T) {
//	type args struct {
//		method string
//		url    string
//		body   io.Reader
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    []byte
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//		{
//			name: "simpleExample",
//			args: args{
//				`GET`,
//				`https://www.baidu.com`,
//				nil,
//			},
//			want:    []byte{},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := SimpleScrape(tt.args.method, tt.args.url, tt.args.body)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			fmt.Println(string(got))
//		})
//	}
//}
//
//func BenchmarkSimpleScrape(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		SimpleScrape(`GET`, `https://ssr1.scrape.center`, nil)
//		//fmt.Println(string(HTML))
//	}
//}
//
//func benchmarkSimpleScrape(b *testing.B, page int) {
//	for i := 1; i < b.N; i++ {
//		for i := 0; i <= page; i++ {
//			SimpleScrape(`GET`, `https://ssr1.scrape.center/page/`+strconv.Itoa(i), nil)
//		}
//	}
//}
//
//func BenchmarkSimpleScrape1(b *testing.B) { benchmarkSimpleScrape(b, 1) }
//
//func BenchmarkSimpleScrape2(b *testing.B) { benchmarkSimpleScrape(b, 2) }
//
//func BenchmarkSimpleScrape5(b *testing.B) { benchmarkSimpleScrape(b, 5) }
//
//func BenchmarkSimpleScrape10(b *testing.B) { benchmarkSimpleScrape(b, 10) }
//
//func BenchmarkSimpleScrapeParallel(b *testing.B) {
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			for i := 0; i < 100; i++ {
//				Body, _ := SimpleScrape(`GET`, `https://www.baidu.com`, nil)
//				fmt.Println(string(Body))
//			}
//		}
//	})
//}

func TestReqArgs_SimpleScrape(t *testing.T) {
	type fields struct {
		method string
		url    string
		header map[string]string
		data   map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: `TestReqArgs_SimpleScrape`,
			fields: struct {
				method string
				url    string
				header map[string]string
				data   map[string]string
			}{
				method: http.MethodGet,
				url:    `https://httpbin.org/get`,
				header: nil,
				data:   nil,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ReqArgs{
				method: tt.fields.method,
				url:    tt.fields.url,
				header: tt.fields.header,
				data:   tt.fields.data,
			}
			got, err := a.SimpleScrape()
			if (err != nil) != tt.wantErr {
				t.Errorf("SimpleScrape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleScrape() got = %v, want %v", string(got), tt.want)
			}
		})
	}
}
