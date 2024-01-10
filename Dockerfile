# TODO Refactor to use: ARG TEMPORAL_VERSION=1.16.2

FROM golang:1.18 AS base

WORKDIR /go/src/

COPY go.mod go.sum ./

RUN go mod download

FROM base AS build

#COPY activities ./activities
#COPY api ./api
#COPY cli ./cli
#COPY temporal ./temporal
#COPY utils ./utils
#COPY ui ./ui
COPY cmd ./cmd
COPY internal ./internal
#COPY workflows ./workflows

#RUN go install -v ./cmd/tfsecret-backend
#RUN go install -v ./cli/tfsecret-company
#RUN go install -v ./cli/tfsecret-candidate
#RUN go install -v ./cli/tfsecret-researcher
##RUN go install -v ./temporal/dataconverter-plugin

FROM golang:1.18 AS app

#ENV TEMPORAL_CLI_PLUGIN_DATA_CONVERTER=dataconverter-plugin

COPY --from=temporalio/admin-tools:1.16.2 /usr/local/bin/tctl /usr/local/bin/tctl
#COPY --from=build /go/bin/dataconverter-plugin /usr/local/bin/dataconverter-plugin

#COPY --from=build /go/bin/tfsecret-backend /usr/local/bin/tfsecret-backend
#COPY --from=build /go/bin/tfsecret-company /usr/local/bin/tfsecret-company
#COPY --from=build /go/bin/tfsecret-candidate /usr/local/bin/tfsecret-candidate
#COPY --from=build /go/bin/tfsecret-researcher /usr/local/bin/tfsecret-researcher
