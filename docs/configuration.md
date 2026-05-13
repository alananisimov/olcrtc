# Configuration

olcrtc accepts the same settings via CLI flags or a YAML file. Use whichever
fits your deployment:

```bash
# CLI flags (existing behaviour)
olcrtc -mode srv -auth wbstream -id room123 -key $(openssl rand -hex 32) ...

# YAML file
olcrtc -config /etc/olcrtc/server.yaml

# YAML file plus CLI overrides — any flag wins over the corresponding YAML field
olcrtc -config /etc/olcrtc/server.yaml -id room999
```

Examples:

- [`server.example.yaml`](./server.example.yaml)
- [`client.example.yaml`](./client.example.yaml)

## Schema

| YAML path                  | CLI flag             | Notes                                         |
|----------------------------|----------------------|-----------------------------------------------|
| `mode`                     | `-mode`              | `srv`, `cnc`, or `gen`                        |
| `link`                     | `-link`              | `direct`                                      |
| `auth.provider`            | `-auth`              | `telemost`, `jazz`, `wbstream`, `none`        |
| `room.id`                  | `-id`                | conference room id                            |
| `room.client_id`           | `-client-id`         | deprecated, will be removed                   |
| `crypto.key`               | `-key`               | 64-char hex (32 bytes)                        |
| `net.transport`            | `-transport`         | `datachannel`, `videochannel`, `seichannel`, `vp8channel` |
| `net.dns`                  | `-dns`               | resolver `host:port`                          |
| `socks.host` / `.port`     | `-socks-host` / `-socks-port` | client-side listener                  |
| `socks.user` / `.pass`     | `-socks-user` / `-socks-pass` | optional client-side auth             |
| `socks.proxy_addr` / `.proxy_port` | `-socks-proxy` / `-socks-proxy-port` | server-side egress proxy   |
| `engine.name` / `.url` / `.token` | `-engine` / `-url` / `-token` | only when `auth.provider: none` |
| `video.*`                  | `-video-*`           | videochannel tuning                           |
| `vp8.*`                    | `-vp8-*`             | vp8channel tuning                             |
| `sei.fps` / `.batch_size` / `.fragment_size` / `.ack_timeout_ms` | `-fps` / `-batch` / `-frag` / `-ack-ms` | seichannel tuning |
| `gen.amount`               | `-amount`            | gen mode: number of rooms to create           |
| `data`                     | `-data`              | path to data directory                        |
| `debug`                    | `-debug`             | verbose logging                               |
| `ffmpeg`                   | `-ffmpeg`            | path to ffmpeg binary                         |

## Precedence

`CLI flag (non-zero) > YAML value > zero value`.

A CLI flag with its zero value (e.g. `-socks-port 0`) does NOT override a YAML
value — pass an explicit non-zero value to override.
