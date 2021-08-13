# Docker

* [Commands](docker/commands.md)
* [Dockerfile](docker/docker-file.md)
* [Docker Compose](docker/docker-compose.md)

## Network

Default to bridge, but you can use host
This file is used to configure the Cidr: `/etc/docker/daemon.json`

```bash
cat /etc/docker/daemon.json

ip route show | grep docker0

ifconfig docker0
```

## Files

## Reference

* <https://docs.docker.com/network/>
* <https://docs.docker.com/network/host/>
