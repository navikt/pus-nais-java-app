package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type JavaOption struct {
	Key   string
	Value string
}

type JavaOptions []JavaOption

func (o JavaOption) Format() string {
	return fmt.Sprintf("-D%s=%s", o.Key, strconv.Quote(o.Value))
}

func NewJavaOption(key string, value string) JavaOption {
	return JavaOption{
		Key:   key,
		Value: value,
	}
}

func (o JavaOptions) Format() string {
	s := make([]string, len(o))
	for i, opt := range o {
		s[i] = opt.Format()
	}
	return strings.Join(s, " ")
}

func httpOpts(flags JavaOptions) (JavaOptions, error) {
	v, found := os.LookupEnv("HTTP_PROXY")
	if !found || len(v) == 0 {
		return flags, nil
	}

	u, err := url.Parse(v)
	if err != nil {
		return flags, err
	}

	if len(u.Hostname()) == 0 || len(u.Port()) == 0 {
		return flags, fmt.Errorf("if specifying a proxy URL, both hostname and port is required")
	}

	flags = append(flags, NewJavaOption("http.proxyHost", u.Hostname()))
	flags = append(flags, NewJavaOption("https.proxyHost", u.Hostname()))
	flags = append(flags, NewJavaOption("http.proxyPort", u.Port()))
	flags = append(flags, NewJavaOption("https.proxyPort", u.Port()))

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

func noProxyOpts(flags JavaOptions) (JavaOptions, error) {
	v, found := os.LookupEnv("NO_PROXY")
	if !found || len(v) == 0 {
		return flags, nil
	}

	hosts := mangleWildcard(strings.Split(v, ","))

	if len(hosts) > 0 {
		flags = append(flags, NewJavaOption("http.nonProxyHosts", strings.Join(hosts, "|")))
	}

	return flags, nil
}

func ProxyOptions() (s string, err error) {
	flags := make(JavaOptions, 0)

	flags, err = httpOpts(flags)
	if err != nil {
		err = fmt.Errorf("error in parsing proxy URL: %s", err)
		return
	}

	flags, _ = noProxyOpts(flags)

	s = flags.Format()

	return
}

func main() {
	s, err := ProxyOptions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, s)
}
