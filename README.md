# Welcome to your CDK Expenses project


## Pre requirements

* `brew install awscli`                           Installation of the official Amazon AWS command-line interface
* `brew install awscli-local`                     AWS-cli local - Thin wrapper around the `aws` command-line interface for use with LocalStack
* `brew install localstack/tap/localstack-cl`     Local Stack installation
* `npm install -g aws-cdk`                        CDK Installation
* `npm install -g aws-cdk-local`                  CDK Local installation

### Before start working, run this command:

This command will make sure you are looking for localstack profile. For each new terminal tab, run this command.

* `chmod +x local-setup.sh && . ./local-setup.sh` Setup the localstack profile and make it the current one


## Useful commands

* `awslocal sts get-caller-identity`                      Shows the current AWS account
* `npm run local:start`                                   Start AWS services locally
* `npm run local:services-status`                         Shows the status of the services running in the LocalStack environment
* `awslocal apigateway get-rest-apis | grep -E '\"id\"'`  Shows the api gateway id
* `local:deploy:first-time`                               Execute when deploying CDK stacks for the first time
* `local:deploy:all`                                      Deploy all CDK stack together

#### Run these commands only if you have run `local:deploy:first-time` at least once
* `local:deploy:api-gateway`                              Deploy only Api Gateway CDK stack.
* `local:deploy:lambda`                                   Deploy only Lambda CDK stack.
* `local:deploy:s3`                                       Deploy only S3 CDK stack.



