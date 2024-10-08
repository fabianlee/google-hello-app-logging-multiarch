# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# [START gke_quickstarts_hello_app_dockerfile]
FROM golang:1.23.1-alpine3.20 AS builder
# trivy will find HIGH vuln in Go 1.21
#FROM golang:1.21.13-alpine3.19 AS builder
WORKDIR /app
RUN go mod init hello-app
COPY *.go ./
ARG MY_VERSION=0.1
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Version=${MY_VERSION}" -o /hello-app

FROM gcr.io/distroless/base-debian12
# trivy will find HIGH vuln in older Debian dist
#FROM debian:bullseye-20240812-slim
WORKDIR /
COPY --from=builder /hello-app /hello-app
ENV PORT=8080
USER nonroot:nonroot
CMD ["/hello-app"]
# [END gke_quickstarts_hello_app_dockerfile]
