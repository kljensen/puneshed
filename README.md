# Puneshed

A go web program that helps me punish my children
by blocking web traffic using
[unbound](https://nlnetlabs.nl/projects/unbound/about/)
and 
[pf](https://www.openbsd.org/faq/pf/)
on 
[OpenBSD](https://www.openbsd.org/).


We use OpenBSD on our 
We use OpenBSD on our router at home, with a setup akin
to that described in this [example](https://www.openbsd.org/faq/pf/example1.html). That router also hands out DHCP leases
to clients on the local network, telling the clients to use
DNS server on the router, which is unbound. Our unbound setup
has a bunch of rules for ad-blocking, malware blocking, etc.
Same with pf.

I often wish to limit my kids' internet access. It is inconvenient to ssh into our router to do so. This is a web application that runs the necessary shell commands to add unbound and pf rules that enable or disable internet access.

## License

See the `LICENSE` file (the unlicense).