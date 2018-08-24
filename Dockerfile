from golang:1.10

WORKDIR /go/src/daisy

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
COPY Gopkg.* ./
RUN dep ensure -vendor-only

COPY . .
RUN dep ensure
RUN go install
ENTRYPOINT ["daisy"]
