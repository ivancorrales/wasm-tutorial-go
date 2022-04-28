FROM golang as build
ENV PROJECT github.com/ivancorrales/wasm-pos-tagging-service
WORKDIR /build
ADD . .
RUN GOOS=js GOARCH=wasm go build -o bin/pos-tagging-service.sample1.wasm $PROJECT/wasm/sample1
RUN GOOS=js GOARCH=wasm go build -o bin/pos-tagging-service.sample2.wasm $PROJECT/wasm/sample2
RUN GOOS=js GOARCH=wasm go build -o bin/pos-tagging-service.sample3.wasm $PROJECT/wasm/sample3
RUN GOOS=js GOARCH=wasm go build -o bin/pos-tagging-service.sample4.wasm $PROJECT/wasm/sample4
RUN GOOS=js GOARCH=wasm go build -o bin/pos-tagging-service.sample5.wasm $PROJECT/wasm/sample5
RUN GOOS=js GOARCH=wasm go build -o bin/pos-tagging-service.sample6.wasm $PROJECT/wasm/sample6
RUN GOOS=js GOARCH=wasm go build -o bin/pos-tagging-service.sample7.wasm $PROJECT/wasm/sample7

RUN CGO_ENABLED=0 GOOS=linux go build -o pos-tagging-service main.go

RUN cat /etc/passwd | grep nobody > passwd.nobody


### Run
FROM scratch
WORKDIR /var/www

COPY --from=build /build/passwd.nobody /etc/passwd
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/local/go/misc/wasm/wasm_exec.js .

COPY --from=build /build/bin/* .

ADD *.html .
ADD style.css .
ADD load_wasm.js .

COPY --from=build /build/pos-tagging-service .

USER nobody
EXPOSE 3000
CMD ["./pos-tagging-service","/var/www"]