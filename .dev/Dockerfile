#FROM mcr.microsoft.com/devcontainers/go:1.22-bookworm
FROM cimg/go:1.22

ARG APP=ginnavsvr

# Set CGO_CFLAGS to enable large file support
ENV CGO_CFLAGS "-D_LARGEFILE64_SOURCE"

RUN go install github.com/google/wire/cmd/wire@latest \
    && go install github.com/swaggo/swag/cmd/swag@latest \
    && go install github.com/gin-admin/gin-admin-cli/v10@latest \
    


