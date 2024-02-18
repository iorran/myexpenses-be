#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import {ApiGatewayStack} from "../lib/api-gateway-stack";
import {LambdaStack} from "../lib/lambda-stack";
import {S3Stack} from "../lib/s3-stack";

const app = new cdk.App();
new ApiGatewayStack(app, 'ApiGatewayStack', {});
new LambdaStack(app, 'LambdaStack', {});
new S3Stack(app, 'S3Stack', {});

