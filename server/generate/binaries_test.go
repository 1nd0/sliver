package generate

import (
	"fmt"
	clientpb "github.com/bishopfox/sliver/protobuf/client"
	"github.com/bishopfox/sliver/server/log"
	"testing"
)

var (
	buildTestLog = log.NamedLogger("generate", "testbuild")
)

func TestSliverExecutableWindows(t *testing.T) {

	// mTLS C2
	mtlsExe(t, "windows", "amd64", false)
	mtlsExe(t, "windows", "386", false)
	mtlsExe(t, "windows", "amd64", true)
	mtlsExe(t, "windows", "386", true)

	// DNS C2
	dnsExe(t, "windows", "amd64", false)
	dnsExe(t, "windows", "386", false)
	dnsExe(t, "windows", "amd64", true)
	dnsExe(t, "windows", "386", true)

	// HTTP C2
	httpExe(t, "windows", "amd64", false)
	httpExe(t, "windows", "386", false)
	httpExe(t, "windows", "amd64", true)
	httpExe(t, "windows", "386", true)

	// Multiple C2s
	multiExe(t, "windows", "amd64", true)
	multiExe(t, "windows", "amd64", false)
}

func TestSliverSharedLibWindows(t *testing.T) {
	multiLibrary(t, "windows", "amd64", true)
	multiLibrary(t, "windows", "amd64", false)
	multiLibrary(t, "windows", "386", true)
	multiLibrary(t, "windows", "386", false)
}

func TestSliverExecutableLinux(t *testing.T) {
	multiExe(t, "linux", "amd64", true)
	multiExe(t, "linux", "amd64", false)
}

func TestSliverExecutableDarwin(t *testing.T) {
	multiExe(t, "darwin", "amd64", true)
	multiExe(t, "darwin", "amd64", false)
}

func mtlsExe(t *testing.T, goos string, goarch string, debug bool) {
	t.Logf("[mtls] EXE %s/%s - debug: %v", goos, goarch, debug)
	config := &SliverConfig{
		GOOS:   goos,
		GOARCH: goarch,
		C2: []SliverC2{
			SliverC2{URL: "mtls://1.example.com"},
		},
		MTLSc2Enabled: true,
		Debug:         debug,
	}
	_, err := SliverExecutable(config)
	if err != nil {
		t.Errorf(fmt.Sprintf("%v", err))
	}
}

func dnsExe(t *testing.T, goos string, goarch string, debug bool) {
	t.Logf("[dns] EXE %s/%s - debug: %v", goos, goarch, debug)
	config := &SliverConfig{
		GOOS:   goos,
		GOARCH: goarch,
		C2: []SliverC2{
			SliverC2{URL: "dns://3.example.com"},
		},
		DNSc2Enabled: true,
		Debug:        debug,
	}
	_, err := SliverExecutable(config)
	if err != nil {
		t.Errorf(fmt.Sprintf("%v", err))
	}
}

func httpExe(t *testing.T, goos string, goarch string, debug bool) {
	t.Logf("[http] EXE %s/%s - debug: %v", goos, goarch, debug)
	config := &SliverConfig{
		GOOS:   goos,
		GOARCH: goarch,
		C2: []SliverC2{
			SliverC2{
				Priority: 1,
				URL:      "http://4.example.com",
				Options:  "asdf",
			},
		},
		HTTPc2Enabled: true,
		Debug:         debug,
	}
	_, err := SliverExecutable(config)
	if err != nil {
		t.Errorf(fmt.Sprintf("%v", err))
	}
}

func multiExe(t *testing.T, goos string, goarch string, debug bool) {
	t.Logf("[multi] %s/%s - debug: %v", goos, goarch, debug)
	config := &SliverConfig{
		GOOS:   goos,
		GOARCH: goarch,

		C2: []SliverC2{
			SliverC2{URL: "mtls://1.example.com"},
			SliverC2{URL: "mtls://2.example.com", Options: "asdf"},
			SliverC2{URL: "https://3.example.com"},
			SliverC2{Priority: 3, URL: "dns://4.example.com"},
		},
		MTLSc2Enabled: true,
		HTTPc2Enabled: true,
		DNSc2Enabled:  true,
		Debug:         debug,
	}
	_, err := SliverExecutable(config)
	if err != nil {
		t.Errorf(fmt.Sprintf("%v", err))
	}
}

func multiLibrary(t *testing.T, goos string, goarch string, debug bool) {
	t.Logf("[multi] LIB %s/%s - debug: %v", goos, goarch, debug)
	config := &SliverConfig{
		GOOS:   goos,
		GOARCH: goarch,

		C2: []SliverC2{
			SliverC2{URL: "mtls://1.example.com"},
			SliverC2{Priority: 2, URL: "mtls://2.example.com"},
			SliverC2{URL: "https://3.example.com"},
			SliverC2{URL: "dns://4.example.com", Options: "asdf"},
		},

		Debug:  debug,
		Format: clientpb.SliverConfig_SHARED_LIB,
	}
	_, err := SliverSharedLibrary(config)
	if err != nil {
		t.Errorf(fmt.Sprintf("%v", err))
	}
}
