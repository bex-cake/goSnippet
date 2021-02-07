FROM golang
RUN mkdir -p /snippet
WORKDIR /snippet
COPY . /snippet
CMD ["go", "run", "./cmd/web"]