## TODO
- [x] websocket support
- [ ] ~~bring back DHT, currently it cheats by using etcd discovery for finding peer address and connect with libp2p (DHT was flaky cz new address would take time to propagate 🤷 )~~ remove etcd discovery, currently its is used to pass lowercase(hash), save that info in DHT or libp2p gossip
- [ ] webextension someday 🤞 (browser extension would route <hash>.lpweb traffic through js impl of libp2p hence no need to run seperate proxy)


## Run in seperate host

```bash


docker run -it golang:1.21

apt update && apt install -y tmux && git clone https://github.com/blue-monads/lpweb && cd lpweb && git checkout mj-packet-splitter

tmux # Ctrl+b " -> to split screen

python3 -m http.server # in one tmux shell

mkdir -p /root/.config/lpweb/

/usr/local/go/bin/go run main.go http-tunnel --port=8000 # in another tmux shell