from golang:1.17.3-bullseye as builder

WORKDIR /app
COPY . /app

ENV CGO_ENABLED 0

RUN  go mod vendor \
     go generate -v api/api.go \
  && go build -v -o fog-wallet

FROM debian:bullseye-slim

RUN  addgroup --system --gid 1000 app \
  && adduser --system --ingroup app --uid 1000 app \
  && apt-get update \
  && apt-get upgrade -y \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/fog-wallet /

USER app

EXPOSE 8080

CMD [ "/fog-wallet" ]
