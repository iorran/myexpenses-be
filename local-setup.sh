# Define variables with credentials and configs
AWS_ACCESS_KEY_ID="test"
AWS_SECRET_ACCESS_KEY="test"
AWS_REGION="us-east-1"
AWS_OUTPUT="json"
LOCALSTACK_ENDPOINT_URL="http://localhost.localstack.cloud:4566"

# Configure 'localstack' profile
aws configure set aws_access_key_id "$AWS_ACCESS_KEY_ID" --profile localstack
aws configure set aws_secret_access_key "$AWS_SECRET_ACCESS_KEY" --profile localstack
aws configure set region "$AWS_REGION" --profile localstack
aws configure set output "$AWS_OUTPUT" --profile localstack
aws configure set endpoint_url "$LOCALSTACK_ENDPOINT_URL" --profile localstack

export AWS_PROFILE=localstack
export AWS_ENDPOINT_URL=$LOCALSTACK_ENDPOINT_URL
