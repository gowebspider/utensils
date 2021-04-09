package Head

import (
	"fmt"
	"math/rand"
	"net/http"
)

// SetRandomPcUserAgent is Set pc fake user agent func,
// Receive *http.Request pointer when in use
func SetRandomPcUserAgent(r *http.Request) {
	r.Header.Set(`User-Agent`, constructionPcUserAgent[rand.Intn(len(constructionPcUserAgent))]())
}

// SetRandomMobileUserAgent is Set Mobile fake user agent func,
// Receive *http.Request pointer when in use
func SetRandomMobileUserAgent(r *http.Request) {
	r.Header.Set(`User-Agent`, constructionMobileUserAgent[rand.Intn(len(constructionMobileUserAgent))]())
}

// SetRandomAllUserAgent is Set all fake user agent func,
// Receive *http.Request pointer when in use
func SetRandomAllUserAgent(r *http.Request) {
	r.Header.Set(`User-Agent`, constructionAllUserAgent()[rand.Intn(len(constructionAllUserAgent()))]())
}

// RandomPcUserAgent is Get pc fake user agent func,
// Receive *http.Request pointer when in use
func RandomPcUserAgent() string {
	return constructionPcUserAgent[rand.Intn(len(constructionPcUserAgent))]()
}

// RandomMobileUserAgent is Get Mobile fake user agent func,
// Receive *http.Request pointer when in use
func RandomMobileUserAgent() string {
	return constructionMobileUserAgent[rand.Intn(len(constructionMobileUserAgent))]()
}

// RandomAllUserAgent is Get all fake user agent func,
// Receive *http.Request pointer when in use
func RandomAllUserAgent() string {
	return constructionAllUserAgent()[rand.Intn(len(constructionAllUserAgent()))]()
}

// statement UserAgent Types
var constructionPcUserAgent = []func() string{
	constructionOperaUA,
	constructionChromeUA,
	constructionFirefoxUA,
}
var constructionMobileUserAgent = []func() string{
	constructionMobileUcUA,
	constructionMobileNexus10UA,
}

func constructionAllUserAgent() []func() string {
	return append(constructionPcUserAgent, constructionMobileUserAgent...)
}

// statement PC Mark and browser Versions
var (
	systemMark = []string{
		`Macintosh; Intel Mac OS X 10_10`,
		`Windows NT 10.0`,
		`Windows NT 5.1`,
		`Windows NT 6.1; WOW64`,
		`Windows NT 6.1; Win64; x64`,
		`X11; Linux x86_64`,
	}
	firefoxVersions = []float32{
		35.0,
		40.0,
		41.0,
		44.0,
		45.0,
		48.0,
		48.0,
		49.0,
		50.0,
		52.0,
		52.0,
		53.0,
		54.0,
		56.0,
		57.0,
		57.0,
		58.0,
		58.0,
		59.0,
		6.0,
		60.0,
		61.0,
		63.0,
	}
	chromeVersions = []string{
		`37.0.2062.124`,
		`40.0.2214.93`,
		`41.0.2228.0`,
		`49.0.2623.112`,
		`55.0.2883.87`,
		`56.0.2924.87`,
		`57.0.2987.133`,
		`61.0.3163.100`,
		`63.0.3239.132`,
		`64.0.3282.0`,
		`65.0.3325.146`,
		`68.0.3440.106`,
		`69.0.3497.100`,
		`70.0.3538.102`,
		`74.0.3729.169`,
	}
	operaVersions = []string{
		`2.7.62 Version/11.00`,
		`2.2.15 Version/10.10`,
		`2.9.168 Version/11.50`,
		`2.2.15 Version/10.00`,
		`2.8.131 Version/11.11`,
		`2.5.24 Version/10.54`,
	}
	systemTypes = systemMark[rand.Intn(len(systemMark))]
)

// construction Browser UserAgent for PC
func constructionFirefoxUA() string {
	firefoxVersion := firefoxVersions[rand.Intn(len(firefoxVersions))]
	systemTypes := systemMark[rand.Intn(len(systemMark))]
	return fmt.Sprintf(`Mozilla/5.0 (%s; rv:%.1f) Gecko/20100101 Firefox/%.1f`, systemTypes, firefoxVersion, firefoxVersion)
}
func constructionChromeUA() string {
	chromeVersion := chromeVersions[rand.Intn(len(chromeVersions))]
	return fmt.Sprintf(`Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36`, systemTypes, chromeVersion)
}
func constructionOperaUA() string {
	version := operaVersions[rand.Intn(len(operaVersions))]
	return fmt.Sprintf(`Opera/9.80 (%s; U; en) Presto/%s`, systemMark, version)
}

// statement system mark and browser for Mobile
var (
	ucVersions = []string{
		`10.9.8.1006`,
		`11.0.0.1016`,
		`11.0.6.1040`,
		`11.1.0.1041`,
		`11.1.1.1091`,
		`11.1.2.1113`,
		`11.1.3.1128`,
		`11.2.0.1125`,
		`11.3.0.1130`,
		`11.4.0.1180`,
		`11.4.1.1138`,
		`11.5.2.1188`,
	}
	androidVersions = []string{
		`4.4.2`,
		`4.4.4`,
		`5.0`,
		`5.0.1`,
		`5.0.2`,
		`5.1`,
		`5.1.1`,
		`5.1.2`,
		`6.0`,
		`6.0.1`,
		`7.0`,
		`7.1.1`,
		`7.1.2`,
		`8.0.0`,
		`8.1.0`,
		`9`,
	}
	ucDevices = []string{
		`SM-C111`,
		`SM-J727T1`,
		`SM-J701F`,
		`SM-J330G`,
		`SM-N900`,
		`DLI-TL20`,
		`LG-X230`,
		`AS-5433_Secret`,
		`IdeaTabA1000-G`,
		`GT-S5360`,
		`HTC_Desire_601_dual_sim`,
		`ALCATEL_ONE_TOUCH_7025D`,
		`SM-N910H`,
		`Micromax_Q4101`,
		`SM-G600FY`,
	}
	nexus10Builds = []string{
		`JOP40D`,
		`JOP40F`,
		`JVP15I`,
		`JVP15P`,
		`JWR66Y`,
		`KTU84P`,
		`LMY47D`,
		`LMY47V`,
		`LMY48M`,
		`LMY48T`,
		`LMY48X`,
		`LMY49F`,
		`LMY49H`,
		`LRX21P`,
		`NOF27C`,
	}
	nexus10Safari = []string{
		`534.30`,
		`535.19`,
		`537.22`,
		`537.31`,
		`537.36`,
		`600.1.4`,
	}

	android = androidVersions[rand.Intn(len(androidVersions))]
)

func constructionMobileUcUA() string {
	device := ucDevices[rand.Intn(len(ucDevices))]
	version := ucVersions[rand.Intn(len(ucVersions))]
	return fmt.Sprintf(`UCWEB/2.0 (Java; U; MIDP-2.0; Nokia203/20.37) U2/1.0.0 UCMini/%s (SpeedMode; Proxy; Android %s; %s ) U2/1.0.0 Mobile`, version, android, device)
}

func constructionMobileNexus10UA() string {
	build := nexus10Builds[rand.Intn(len(nexus10Builds))]
	chrome := chromeVersions[rand.Intn(len(chromeVersions))]
	safari := nexus10Safari[rand.Intn(len(nexus10Safari))]
	return fmt.Sprintf(`Mozilla/5.0 (Linux; Android %s; Nexus 10 Build/%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/%s`, android, build, chrome, safari)
}
