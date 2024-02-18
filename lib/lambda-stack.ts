import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { GoFunction } from '@aws-cdk/aws-lambda-go-alpha';
import { aws_iam as iam } from "aws-cdk-lib";
import * as path from 'path';

import {BUCKET_NAME_INVOICES} from "./s3-stack";

export class LambdaStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const DB_NAME = "expenses";
    const DB_COLLECTION = "invoices";
    const DB_URI = "mongodb+srv://<username>:<password>d@amongold.xbxpmxt.mongodb.net/?retryWrites=true&w=majority";

    const lambdaSaveInvoice = new GoFunction(this, "LambdaSaveInvoice", {
      entry: path.join(__dirname, '../src/lambda/save-invoice.go'),
      functionName: LAMBDA_NAME_SAVE_INVOICE,
      environment: {
        BUCKET_NAME: BUCKET_NAME_INVOICES,
        DB_URI,
        DB_NAME,
        DB_COLLECTION,
        URL_S3_LOCAL: "http://s3.localhost.localstack.cloud:4566",
        ENVIRONMENT: "local"
      }
    });

    const bucketKeystorePolicy = new iam.PolicyStatement({
      effect: iam.Effect.ALLOW,
      actions: [
        's3:List',
        's3:PutObject',
      ],
      resources: [
        `arn:aws:s3:::${BUCKET_NAME_INVOICES}`,
        `arn:aws:s3:::${BUCKET_NAME_INVOICES}/*`
      ],
    });

    lambdaSaveInvoice.addToRolePolicy(bucketKeystorePolicy)

    new GoFunction(this, "LambdaSearchInvoice", {
      entry: path.join(__dirname, '../src/lambda/search-invoice.go'),
      functionName: LAMBDA_NAME_SEARCH_INVOICE,
      environment: {
        DB_URI,
        DB_NAME,
        DB_COLLECTION,
      }
    });

    new GoFunction(this, "LambdaDeleteInvoice", {
      entry: path.join(__dirname, '../src/lambda/delete-invoice.go'),
      functionName: LAMBDA_NAME_DELETE_INVOICE,
      environment: {
        DB_URI,
        DB_NAME,
        DB_COLLECTION,
      }
    });
  }
}

export const LAMBDA_NAME_SAVE_INVOICE = "SaveInvoice"
export const LAMBDA_NAME_SEARCH_INVOICE = "SearchInvoice"
export const LAMBDA_NAME_DELETE_INVOICE = "DeleteInvoice"
