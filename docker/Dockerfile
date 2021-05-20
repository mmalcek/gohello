FROM golang:latest
RUN go get -v golang.org/x/tools/gopls  && \
    go get -v golang.org/x/tools/cmd/goimports  && \
    go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs  && \
    go get -v github.com/ramya-rao-a/go-outline && \
    go get -v github.com/go-delve/delve/cmd/dlv && \
    mkdir /workspaces
COPY init.sh /home
WORKDIR "/workspaces"
ENTRYPOINT ["/home/init.sh"]