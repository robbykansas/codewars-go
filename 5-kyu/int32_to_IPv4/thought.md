- Use net.IP from package [net](https://pkg.go.dev/net) 4 means IPv4
- n >> 24 means Shifts the bits of n 24 positions to the right, effectively isolating the most significant byte (the first octet of the IP address).
- n >> 16 means isolates the second octet.
- n >> 8 means isolates the third octet.
- use fmt for result

best solve
- instead using net, just mod it with 256, since the maximum of octet is 255