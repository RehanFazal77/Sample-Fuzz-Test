FROM gcr.io/oss-fuzz-base/base-builder-go
WORKDIR ~/Sample-Fuzz-Test
COPY . .
RUN go get github.com/AdamKorcz/go-118-fuzz-build/testing
