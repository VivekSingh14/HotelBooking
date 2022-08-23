FROM golang:alpine as builder


WORKDIR /cmd

COPY ./ /cmd
COPY go.mod go.sum ./
RUN go mod download

RUN go build -o bin/main app/*  


FROM alpine:3.14
# RUN addgroup -S cmd && \
#     adduser -S -G cmd cmd
# RUN mkdir /cmd
# WORKDIR /cmd
#RUN chown -R cmd:cmd /cmd
COPY config/config-env.yaml config/
#COPY --chown=cmd:cmd --chmod=0740 --from=builder /cmd/bin/main .
COPY --from=builder /cmd/bin/main .
#USER cmd
CMD ["./main"]