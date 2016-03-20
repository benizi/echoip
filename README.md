# echoip

Super simple Golang server that just echos the client's IP address + port.

(Just something I wanted for debugging.)

# Simpler

This is basically equivalent to:

```sh
socat \
  tcp-listen:${port?},reuseaddr,pktinfo,fork \
  system:'echo HTTP/1.1 200 OK; echo; echo $SOCAT_PEERADDR\:$SOCAT_PEERPORT'
```
