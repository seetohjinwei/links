# Links

Go app that generates a static page of links, and serves re-directs.

https://link.jinwei.dev/

- https://link.jinwei.dev/me
- https://link.jinwei.dev/blog
- `https://link.jinwei.dev/...`

To update the links, just update `links.yaml` and re-build the application :)

```sh
go build -o bin/links && ./links

go test ./...
```
