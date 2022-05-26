# Source for https://practicalgobook.net

Hugo + go:embed magic - checkout the [GopherCon 2021 lightning talk](https://www.youtube.com/watch?v=y6VnTM1f2cc)

## Infrastructure

AWS Route 53 - DNS Record -> Linode instance -> Caddy (443) ->  Go server (8080)

