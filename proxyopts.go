package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func httpOpts(flags []string) ([]string, error) {
	v, found := os.LookupEnv("HTTP_PROXY")
	if !found {
		return flags, nil
	}

	u, err := url.Parse(v)
	if err != nil {
		return flags, err
	}

	if len(u.Hostname()) == 0 || len(u.Port()) == 0 {
		return flags, fmt.Errorf("if specifying a proxy URL, both hostname and port is required")
	}

	flags = append(flags, fmt.Sprintf("-Dhttp.proxyHost=\"%s\"", u.Hostname()))
	flags = append(flags, fmt.Sprintf("-Dhttp.proxyPort=\"%s\"", u.Port()))
	flags = append(flags, fmt.Sprintf("-Dhttps.proxyHost=\"%s\"", u.Hostname()))
	flags = append(flags, fmt.Sprintf("-Dhttps.proxyPort=\"%s\"", u.Port()))

	return flags, nil
}

// mangleWildcard takes a list of hostnames and prepends '*' if the hostname
// starts with '.', then returns a new slice with the modified hostnames.
func mangleWildcard(hosts []string) []string {
	mangled := make([]string, len(hosts))
	for i, host := range hosts {
		if len(host) > 0 && host[0] == '.' {
			host = "*" + host
		}
		mangled[i] = host
	}
	return mangled
}

func noProxyOpts(flags []string) ([]string, error) {
	v, found := os.LookupEnv("NO_PROXY")
	if !found {
		return flags, nil
	}

	hosts := mangleWildcard(strings.Split(v, ","))

	if len(hosts) > 0 {
		flags = append(flags, fmt.Sprintf("-Dhttp.nonProxyHosts=\"%s\"", strings.Join(hosts, "|")))
	}

	return flags, nil
}

func main() {
	var err error
	flags := make([]string, 0)

	flags, err = httpOpts(flags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in parsing proxy URL: %s\n", err)
		os.Exit(1)
	}

	flags, _ = noProxyOpts(flags)

	fmt.Fprintf(os.Stdout, strings.Join(flags, " "))
}
