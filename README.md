# HTTP Server/Client examples written in Golang

|Project|Notes|
|---|---|
|01.simple-server|A web server that listens on specific port. When request is received on "/", templated response will be sent. The response contains hostname, request URI and timestamp. The server panics when a service can not be started on given port. It comes with a simple Dockerfile.|
|02.server-static-contents|A web server serving static contents.|
|03.gorilla-mux|There are advanced routers and dispatcher that support request matching based on url, path, header query value etc.|
|04.client|HTTP Client|
