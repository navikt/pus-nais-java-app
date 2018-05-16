#!/bin/sh -e
#
# Convert Linux proxy parameters to their Java equivalents.
#
# For HTTP_PROXY and HTTPS_PROXY, Linux uses the "http://host:port" notation,
# but for Java we have to split them into host and port parts.
#
# For NO_PROXY, Linux uses "host,.wildcard.host,..." whereas Java requires the
# form "host|*.wildcard.host|...".
#
# We assume the proxy settings for HTTP and HTTPS will never differ, so for
# convenience we use only the HTTP_PROXY setting.
#
# Thanks to pjz@stackoverflow for the sh URL parser.
# https://stackoverflow.com/questions/6174220/parse-url-in-shell-script

if [ "$HTTP_PROXY" != "" ]; then
	# extract the protocol
	proto="`echo $HTTP_PROXY | grep '://' | sed -e's,^\(.*://\).*,\1,g'`"
	# remove the protocol
	url=`echo $HTTP_PROXY | sed -e s,$proto,,g`

	# extract the user and password (if any)
	userpass="`echo $url | grep @ | cut -d@ -f1`"
	pass=`echo $userpass | grep : | cut -d: -f2`
	if [ -n "$pass" ]; then
		user=`echo $userpass | grep : | cut -d: -f1`
	else
		user=$userpass
	fi

	# extract the host -- updated
	hostport=`echo $url | sed -e s,$userpass@,,g | cut -d/ -f1`
	port=`echo $hostport | grep : | cut -d: -f2`
	if [ -n "$port" ]; then
		host=`echo $hostport | grep : | cut -d: -f1`
	else
		host=$hostport
	fi

	# extract the path (if any)
	path="`echo $url | grep / | cut -d/ -f2-`"

	# set up the configuration
	non_proxy_hosts=$(echo ${NO_PROXY} | sed 's/,\./|*./g' | sed 's/,/|/g')
	java_proxy_options="
		-Dhttp.proxyHost \"${host}\"
		-Dhttp.proxyPort \"${port}\"
		-Dhttp.nonProxyHosts \"${non_proxy_hosts}\"
		-Dhttps.proxyHost \"${host}\"
		-Dhttps.proxyPort \"${port}\"
	"
fi
