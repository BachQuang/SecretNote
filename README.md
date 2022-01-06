# SecretNote

Please visit this link for demo: http://a606cd567c6334074a5404393d41f48c-1169708641.us-east-2.elb.amazonaws.com

This repository contains the codes of the SecretNote service.

The repository includes design, develop, and deploy a complete system using React, Golang, PostgreSQL, Docker(AWS ECR) and K8s(AWS EKS).


## Secret Note service

The service that I am going to build is a Secret Note.:

Features:
- 1. Login the system via third party: Google Authenticator.
- 2. After login people can read their note privately.
- 3. After login people can create their note privately
- 4. Automatically run CI/CD and deploy to AWSECR
- 5. Setup .yaml files for deploy to AWS EKS(K8s)
## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    brew install golang-migrate
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    brew install sqlc
    ```

### Setup infrastructure

- Start postgres container:

    ```bash
    make postgres
    ```

- Create database:

    ```bash
    make createdb
    ```

- Run db migration up all versions:

    ```bash
    make migrateup
    ```

- Run db migration down all versions:

    ```bash
    make migratedown
    ```


### How to generate code

- Generate SQL CRUD with sqlc:

    ```bash
    make sqlc
    ```

- Generate DB mock with gomock:

    ```bash
    make mock
    ```

- Create a new db migration:

    ```bash
    migrate create -ext sql -dir db/migration -seq <migration_name>
    ```

### How to run locally

- Run client:

    ```bash
    cd frontend
    ```
    ```bash
    nano .env
    ```
    ```bash
    REACT_APP_API_URL = 'http://localhost:8080'
    ```
    ```bash
    yarn start
    ```
- Run server:
    ```bash
    make server
    ```

- Run test:

    ```bash
    make test
    ```

## Push the image to AWS ECR

- Create an AWS account
- Go to IAM in AWS, setup a access key id, and access secret key id
- Store it in github account via secret variable
- Setup secret manager key on AWS : https://docs.aws.amazon.com/secretsmanager/latest/userguide/tutorials_basic.html
- Push the code to your account. It will automatically build an image on your AWS ECR via the file .github/workflows/deploy.yaml

## Deploy to kubernetes cluster

- Setup kubectl and k9s: https://kubernetes.io/docs/tasks/tools/
- Create an AWS EKS and login in via kubectl
- 
    ```bash
    cd backend
    ```
    ```bash
    cd eks
    ```
    ```bash
    nano deployment.yaml
    ```
    ```bash
    Replace the image uri with your docker image that pushed to your AWS ECR before
    ```
    ```bash
    kubectl apply -f backend/eks/deployment.yaml
    ```
    ```bash
    kubectl apply -f backend/eks/service.yaml
    ```
## To do in the future
- Write the mock test via gomock
- Setup fully auto deployment to AWS EKS
