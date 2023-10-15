# ------------- BUILD --------------- #
FROM golang:1.21 as build

RUN mkdir -p /src/build
WORKDIR /src/build

COPY . .

RUN make build

# -------------- RUN ---------------- #
FROM scratch

COPY --from=build /src/build/dist/whoamip ./

EXPOSE 8080 9100

ENTRYPOINT ["./whoamip"]
