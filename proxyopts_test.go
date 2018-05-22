package main

import (
	"os"
	"testing"
)

var testCases = []struct {
	HTTP_PROXY  string
	HTTPS_PROXY string
	NO_PROXY    string
	Output      string
	Error       bool
}{
	{
		HTTP_PROXY: "http://foo.bar:1234",
		Output:     `-Dhttp.proxyHost="foo.bar" -Dhttps.proxyHost="foo.bar" -Dhttp.proxyPort="1234" -Dhttps.proxyPort="1234"`,
	},
	{
		HTTPS_PROXY: "http://foo.bar:1234",
		Output:      ``,
	},
	{
		HTTP_PROXY:  "http://foo.bar:1234",
		HTTPS_PROXY: "http://baz:12345",
		Output:      `-Dhttp.proxyHost="foo.bar" -Dhttps.proxyHost="foo.bar" -Dhttp.proxyPort="1234" -Dhttps.proxyPort="1234"`,
	},
	{
		NO_PROXY: "internalhost",
		Output:   `-Dhttp.nonProxyHosts="internalhost"`,
	},
	{
		NO_PROXY: "host1,host2,.wildcard.local,.local,foo",
		Output:   `-Dhttp.nonProxyHosts="host1|host2|*.wildcard.local|*.local|foo"`,
	},
	{
		HTTP_PROXY: "http://foo.bar:1234",
		NO_PROXY:   "host1,host2,.wildcard.local,.local,foo",
		Output:     `-Dhttp.proxyHost="foo.bar" -Dhttps.proxyHost="foo.bar" -Dhttp.proxyPort="1234" -Dhttps.proxyPort="1234" -Dhttp.nonProxyHosts="host1|host2|*.wildcard.local|*.local|foo"`,
	},
	{
		HTTP_PROXY: "foo.bar:1234",
		Error:      true,
	},
	{
		HTTP_PROXY: "http://proxy",
		Error:      true,
	},
}

func TestSuccess(t *testing.T) {
	for i, test := range testCases {
		os.Setenv("HTTP_PROXY", test.HTTP_PROXY)
		os.Setenv("HTTPS_PROXY", test.HTTPS_PROXY)
		os.Setenv("NO_PROXY", test.NO_PROXY)

		output, err := ProxyOptions()

		if test.Error {
			if err == nil {
				t.Fatalf("Test #%d: expected error, got success instead", i)
			}
		} else {
			if err != nil {
				t.Fatalf("Test #%d: expected success, got error: %s", i, err)
			}
			if output != test.Output {
				t.Fatalf("Test #%d: expected output \"%s\", got \"%s\"", i, test.Output, output)
			}
		}
	}
}
