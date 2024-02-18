import * as cdk from 'aws-cdk-lib';
import {aws_apigateway as apigateway} from 'aws-cdk-lib';
import {Construct} from 'constructs';
import {LAMBDA_NAME_DELETE_INVOICE, LAMBDA_NAME_SAVE_INVOICE, LAMBDA_NAME_SEARCH_INVOICE} from "./lambda-stack";
import {GoFunction} from "@aws-cdk/aws-lambda-go-alpha";
import {AuthorizationType} from "aws-cdk-lib/aws-apigateway";


export class ApiGatewayStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const api = new apigateway.RestApi(this, 'InvoiceApi', {
      restApiName: 'Invoice',
    });

    const defaultCorsPreflightOptions = {
      allowOrigins: apigateway.Cors.ALL_ORIGINS,
      allowMethods: apigateway.Cors.ALL_METHODS,
    }
    const apiOptions = { apiKeyRequired: false, authorizationType: AuthorizationType.NONE}
    const invoices = api.root.addResource("invoices", { defaultCorsPreflightOptions })

    const lambdaSaveInvoice = GoFunction.fromFunctionName(this, "LambdaSaveInvoice", LAMBDA_NAME_SAVE_INVOICE)
    invoices.addMethod("POST", new apigateway.LambdaIntegration(lambdaSaveInvoice), apiOptions)

    const lambdaSearchInvoice = GoFunction.fromFunctionName(this, "LambdaSearchInvoice", LAMBDA_NAME_SEARCH_INVOICE)
    invoices.addMethod("GET", new apigateway.LambdaIntegration(lambdaSearchInvoice), apiOptions)

    const idResource = invoices.addResource("{id}")
    const lambdaDeleteInvoice = GoFunction.fromFunctionName(this, "LambdaDeleteInvoice", LAMBDA_NAME_DELETE_INVOICE)
    idResource.addMethod("DELETE", new apigateway.LambdaIntegration(lambdaDeleteInvoice), apiOptions)

    new cdk.CfnOutput(this, 'RestApiID', {
      value: api.restApiId,
      description: "Rest API ID"
    })
  }
}
