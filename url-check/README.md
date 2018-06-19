# Url Check: minimal tool to test working URLs.

The purpose of this tool is to offer a simple testing interface for http/https 
connection issues in apps written in Go using the net/http package.

It manages http and https connections and accepts insecure tls certificates by default.


Usage:

```
$ url-check https://www.example.com
```

It can be useful to test Docker registries:

```
$ url-check https://myregistry:5000/v2/_catalog
```

Since it uses the net/http package and its transport has support for proxies from env variables
it can be used to test how proxy and proxy exclusions work:

```
$ HTTP_PROXY=http://proxy:8080 HTTPS_PROXY=http://proxy:8080 NO_PROXY=myregistry https://myregistry:5000/v2/_catalog
```

HINT: keep in mind that HTTPS_PROXY has higher priority over HTTP_PROXY.


