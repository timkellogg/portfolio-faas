# Portfolio FAAS

This application handles saving contacts from my portfolio. It is designed to run on AWS Lambda as a Function as a Service to process contacts on a "on-demand" basis.

## Tech Stack

- severless
- golang
- aws lambda
- aws dynamodb

## Setting Up / Configuring

- install [golang v1.10](https://golang.org/doc/install)
- install [node and npm](https://nodejs.org/en/download/)
- install [serverless framework](https://serverless.com/framework/docs/getting-started/)
- install [aws cli](https://docs.aws.amazon.com/cli/latest/userguide/installing.html)
- setup [aws cli user](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html)
- install task go pkg: `go get -u -v github.com/go-task/task/cmd/task`

## Building

- `task build`

## Deploying

- `task deploy`
