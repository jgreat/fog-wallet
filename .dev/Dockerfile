from golang:1.17.3-bullseye

# Share code as volume on /app
WORKDIR /app

RUN apt-get update \
  && apt-get install inotify-tools -y \
  && rm -rf /var/lib/apt/lists/*

RUN addgroup --system --gid 1000 app \
  && adduser --system --ingroup app --uid 1000 app \
  && mkdir -p /go/pkg \
  && chown -R app:app /go/pkg

COPY ./.dev/run.sh /

VOLUME /go/pkg
USER app

CMD [ "/run.sh"]
