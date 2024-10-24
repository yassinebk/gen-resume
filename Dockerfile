FROM surnet/alpine-wkhtmltopdf:3.20.2-0.12.6-full as wkhtmltopdf

FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o genresume .

FROM alpine:latest

RUN apk update && apk add --no-cache \
    libstdc++ \
    libx11 \
    libxrender \
    libxext \
    libssl3 \
    ca-certificates \
    fontconfig \
    freetype \
    ttf-droid \
    ttf-freefont \
    ttf-liberation 
# wkhtmltopdf copy bins from ext image
COPY --from=wkhtmltopdf /bin/wkhtmltopdf /bin/wkhtmltopdf
COPY --from=wkhtmltopdf /bin/wkhtmltoimage /bin/wkhtmltoimage
COPY --from=wkhtmltopdf /lib/libwkhtmltox* /lib/

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/genresume .
COPY --from=builder /app/templates/ ./templates

# Set executable permissions
RUN chmod +x genresume

# Define the default command to run the binary
ENTRYPOINT ["./genresume"]