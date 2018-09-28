# try logrus output stackdriver logging

## Quick Start

When set default logging driver to docker daemon.

```bash
$ docker daemon --log-driver=gcplogs
```

When set logging driver to docker container.

```bash
$ docker run --log-driver=gcplogs ...
```

## reference to study
http://docs.docker.jp/engine/admin/logging/gcplogs.html
