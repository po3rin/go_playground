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
https://cloud.google.com/community/tutorials/docker-gcplogs-driver
https://qiita.com/akm/items/cc6c07181750b728d56a