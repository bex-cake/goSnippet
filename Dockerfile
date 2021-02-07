FROM golang
RUN mkdir -p /usr/src/goSnippetBox/
ADD . /usr/src/goSnippetBox/
WORKDIR /usr/src/goSnippetBox/
ENTRYPOINT ./cmd/web*
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /usr/src/goSnippetBox/cmd/web* .
CMD ["/usr/src/goSnippetBox/cmd/web*"]