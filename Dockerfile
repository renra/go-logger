FROM golang:1.10.7-alpine

RUN apk update && apk add make dep

ENV DIR ${GOPATH}/src/app

RUN mkdir -p ${DIR}
WORKDIR ${DIR}

COPY ./ ${DIR}/

RUN make dep

CMD ["make", "dev"]

