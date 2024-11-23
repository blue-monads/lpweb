# lpweb

lpweb is a way to access http service though libp2p.
This allows to expose your http service even if you only have private ip like services running in your local machine or your home rapPi.
It gives you domain name using your libp2p public key.

<libp2p_key_hash>.lpweb

browser => ( proxy_or_extension_with_libp2p) => (another_libp2p_node_with_http_service)

> But you need access to both ends, run tunnel on one side and use proxy on other 

```bash

# target machine
lpweb http-tunnel --port=4000 # where 4000 is port you want to tunnel, your websevice/ dev port / python3 -m http.server

# go run main.go http-tunnel --port=4000 # same as above

# another machine
lpweb web-proxy --port=8080 # 8080 is a proxy port, use in browser/networking

```

