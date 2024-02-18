import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { aws_s3 as s3} from 'aws-cdk-lib';


export class S3Stack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    new s3.Bucket(this, "InvoiceBucket", {
      bucketName: BUCKET_NAME_INVOICES,
      versioned: false,
      publicReadAccess: true,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
    })
  }
}

export const BUCKET_NAME_INVOICES = "my-expenses-iorran"
