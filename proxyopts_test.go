package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
}

func TestSuccess(t *testing.T) {
	for _, test := range testCases {
		os.Setenv("HTTP_PROXY", test.HTTP_PROXY)
		os.Setenv("HTTPS_PROXY", test.HTTPS_PROXY)
		os.Setenv("NO_PROXY", test.NO_PROXY)

		opts, err := ProxyOptions()

		if test.Error {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, opts, test.Output)
		}
	}
}
