# Domainify

This is a very simple tool, it's NOT recommended to production.

This is a simple valid names provider for domains. You can to parse a string to a valid top-level domain name.
E.j: You can send it `incredible chat$` and it will to return you `incrediblechat.com`, omiting both space and `$` chars.

You can configure your own top-level domain names, sending to the `tlds` flag the values inside quotes.
E.j: `domainify -tlds="info, edu, io, dev"`. Doing this, you will to receive the result with `.info`, `.edu`, `.io`, `.dev`.
