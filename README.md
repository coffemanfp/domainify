# Domainify

This is a very simple tool, it's NOT recommended to production.

This is a simple valid names provider for domains. You can to parse a string to a valid domain url.
E.j: You can send it thought the standard input `incredible chat$` and it will to return you `www.incrediblechat.com`, omiting both space and `$` chars and adding both `www` hostname and `com` top-level domain name.

You can configure your own top-level domain names, sending to the `tlds` flag the values inside quotes.
E.j: `domainify -tlds="info, edu, io, dev"`. Doing this, you will to receive the result with `.info`, `.edu`, `.io`, `.dev`.

You can configure your own hostnames, sending to the `h` flag the values inside quotes.
E.j: `domainify -h="api,web,www"`
