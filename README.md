![image](logo.png)

# About

- SHS is short for Simple HTTP Service
- SHS can be used as an HTTP request debug tool or http service placeholder such as Redfish event subscription target.

# Usage

Help:

```shell

shs -h
Usage of shs:
  -detail
        Show detail info
  -path string
        Service path (default "/")
  -port int
        Port for listen. (default 80)
```

Example:

```shell
shs -path "/test" -port 8080 -detail false
```