package URL

import "fmt"

var all []string = []string{
	"urlparse", "urlunparse", "urljoin", "urldefrag",
	"urlsplit", "urlunsplit", "urlencode", "parse_qs",
	"parse_qsl", "quote", "quote_plus", "quote_from_bytes",
	"unquote", "unquote_plus", "unquote_to_bytes",
	"DefragResult", "ParseResult", "SplitResult",
	"DefragResultBytes", "ParseResultBytes", "SplitResultBytes",
}

func example() {
	fmt.Println(all)
}
