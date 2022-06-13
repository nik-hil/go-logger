# Go Logger tester

Simple app to log into the sysout.

## Problem statement
While testing the observability of my cluster, I faced the diffiutly with logging and response code.
Many times we test the log alerts by making our app to log certain value at certain level. 
Similarly we test the HTTP alerts by making our app to return certain response code (5xx).

The difficult part is to get right log level, log text or right response code which is acceptable for our alerting mechanism.

In such cases dumb approach is to tweak our app to acheive results.
This app solve this pain point.
Just add this app as a side car container and test your observaibility.

## How to use

```bash
curl http://localhost:8080/log \
--header "Content-Type: application/json" \
--request "POST" \
--data  '{"data":"logging single line","statusCode":201, "level":"INFO",}'
```


## Resources

1. https://github.com/chunghha/docker-go-gin/blob/master/Dockerfile
1. https://github.com/gin-gonic/gin
1. https://github.com/IBM/template-go-gin/blob/master/Dockerfile