# PHCK
PHCK is process health checker web server.

## Usage

#### listen web server
```
$ ./phck phck.conf
```

#### check process & result success
HTTP status code 200

```
$ curl http://127.0.0.1:8001/ | jq
{
  "status": true,
  "process": [
    {
      "label": "nginx",
      "pidfile": "/var/run/nginx.pid",
      "running": true,
      "detail": {
        "name": "nginx",
        "pid": 10709,
        "ppid": 1,
        "command": "nginx: master process /usr/sbin/nginx -c /etc/nginx/nginx.conf",
        "stat": "S",
        "thread": 1,
        "use_memory": 0.10912581
      }
    },
    {
      "label": "td-agent",
      "pidfile": "/var/run/td-agent/td-agent.pid",
      "running": true,
      "detail": {
        "name": "ruby",
        "pid": 22499,
        "ppid": 1,
        "command": "/opt/td-agent/embedded/bin/ruby /usr/sbin/td-agent --log /var/log/td-agent/td-agent.log --use-v1-config --group td-agent --daemon /var/run/td-agent/td-agent.pid",
        "stat": "S",
        "thread": 2,
        "use_memory": 1.3995278
      }
    }
  ]
}
```

#### check process & result error
HTTP status code 500

```
$ curl http://127.0.0.1:8001/ | jq
{
  "status": false,
  "process": [
    {
      "label": "nginx",
      "pidfile": "/var/run/nginx.pid",
      "running": true,
      "detail": {
        "name": "nginx",
        "pid": 10709,
        "ppid": 1,
        "command": "nginx: master process /usr/sbin/nginx -c /etc/nginx/nginx.conf",
        "stat": "S",
        "thread": 1,
        "use_memory": 0.10912581
      }
    },
    {
      "label": "td-agent",
      "pidfile": "/var/run/td-agent/td-agent.pid",
      "running": false,
      "message": "PID file not opened",
      "detail": {}
    }
  ]
}
```

#### CLI mode
```
$ ./phck --cli phck.conf | jq
```

## Help
```
Usage:
  PHCK [OPTIONS] CONFIG_FILE

Application Options:
  -c, --cli      CLI mode
  -h, --help     Show this help message
  -v, --version  Show this build version

```