# Quickstart

1. To build plugin run 
>  docker run -it -v "$PWD:/app" -w /app devopsfaith/krakend-plugin-builder:2.1.2 go build -buildmode=plugin -o plugins/graphql-plugin.so .
2. To Run KrakenD with the plugin
> docker run -p 8080:8080 -it -v $PWD:/etc/krakend/ devopsfaith/krakend run --config /etc/krakend/krakend.json