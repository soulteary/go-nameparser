# Base Images
FROM golang:1.20.4-alpine3.18 AS go-builder
FROM python:3.7-alpine3.18 AS builder
# Base Builder Env
COPY --from=go-builder /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
RUN python -m pip install --upgrade pip
RUN apk add build-base pkgconfig
ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig/
ENV CGO_ENABLED=1
# Copy source code
# Install deps
RUN pip install nameparser
COPY ./go.* /app/
WORKDIR /app
RUN go mod tidy
COPY . ./
# Build the binary
RUN go build -o HumanName

# Run Image
FROM python:3.7-alpine3.18
# Copy Python Deps
COPY --from=builder /usr/local/lib/python3.7/site-packages/nameparser /usr/local/lib/python3.7/site-packages/nameparser
COPY --from=builder /app/convert /convert
# Copy binary
COPY --from=builder /app/HumanName /HumanName
CMD ["/HumanName"]