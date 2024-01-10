# Terraform

Using temporal to provision a secret and setting a secret value in AWS.

## Usage

### Execute Locally

```
# Build
make worker
make client

# Execute
./tfsecret-backend &
./tfsecret-client
```

## Prerequisites

- Temporal

## Architecture

### Functional Requirements

#### APIs

- Start secret creation
- Provide secret value
- Cancel secret creation
- List secret creations
- Get report for a specific secret creation

##### UIs

- The requester needs a dedicated UI to start, cancel, or list requested multiple secret creations (only own ones, owned by the team).
- The platform team needs to approve the creation. (just for fun to train usage)
- The requester needs a UI to put the secret value.
- The platform team needs a ui to see all requested secret creations.

##### CLIs

- have a CLI for each user type that uses the same APIs that the UIs do

##### How to handle secret value

- the application must have a way to encrypt and decrypt data
- Data must be consistently encrypted while it is in the Temporal Platform but be decryptable when needed.

#### Wait on humans

- Wait for platform team to approve the request
- Wait on requester to provide the secret value

#### Concurrent

- secret creation can happen in parallel

### Design

#### Workflows

- Secret creation request
- Platform team acceptance
- Infrastructure creation
- Value injection

- Rollback
- Removal

#### Activities

- Sending email
- Calling third-party APIs

#### Data Injection

- We use Signals to get data into a running Workflow Execution.

### Temporal Platform

#### Cluster Setup

Vagrant

- Temporal Server
- Database (Postgresql)
- Elasticsearch

#### Data Encryption

#### REST

- To ensure data is encrypted while in the Temporal Platform, we use a customized Data Converter.

#### Transit

- ???

## References

- [Example: Background Check project](https://docs.temporal.io/docs/learning-paths/background-checks/)
- [Localstack: Secret Manager](https://stackoverflow.com/questions/57154039/how-to-set-up-local-aws-secrets-manager-docker-container-for-local-testing-purpo)
- [Project Structure](https://github.com/golang-standards/project-layout)
- [Sophisticated Temporal Environment Docker Compose Example](https://github.com/temporalio/docker-compose)
- [Demo Temporal AWS VPC, Subnet](https://github.com/dynajoe/temporal-terraform-demo)
- [How to develop a Worker Program in Go](https://docs.temporal.io/go/how-to-develop-a-worker-program-in-go)
- [Test: Workflow](https://docs.temporal.io/go/how-to-test-workflow-definitions-in-go)
