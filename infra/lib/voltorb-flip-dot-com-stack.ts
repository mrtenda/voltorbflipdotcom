import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as s3 from 'aws-cdk-lib/aws-s3';
import * as cloudfront from 'aws-cdk-lib/aws-cloudfront';
import * as origins from 'aws-cdk-lib/aws-cloudfront-origins';
import * as apigwv2 from 'aws-cdk-lib/aws-apigatewayv2';
import { HttpLambdaIntegration } from 'aws-cdk-lib/aws-apigatewayv2-integrations';
import * as s3_deployment from 'aws-cdk-lib/aws-s3-deployment';
import * as acm from 'aws-cdk-lib/aws-certificatemanager';
import * as route53 from 'aws-cdk-lib/aws-route53';
import * as targets from 'aws-cdk-lib/aws-route53-targets';
import { GoFunction } from '@aws-cdk/aws-lambda-go-alpha';
import * as path from 'path';

interface VoltorbFlipDotComStackProps extends cdk.StackProps {
  readonly stage: 'beta' | 'prod';
}

export class VoltorbFlipDotComStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props: VoltorbFlipDotComStackProps) {
    super(scope, id, props);

    const baseDomain = 'voltorbflip.com';
    const domainName = props.stage === 'prod' ? baseDomain : `beta.${baseDomain}`;
    const subjectAlternativeNames = props.stage === 'prod' ? [`www.${baseDomain}`] : [];

    // 0. Route 53 Hosted Zone lookup (always against base domain)
    const zone = route53.PublicHostedZone.fromLookup(this, 'HostedZone', {
      domainName: baseDomain,
    });

    // 1. SSL Certificate (validated via DNS)
    const certificate = new acm.Certificate(this, 'SiteCertificate', {
      domainName: domainName,
      subjectAlternativeNames: subjectAlternativeNames,
      validation: acm.CertificateValidation.fromDns(zone),
    });

    // 2. S3 Bucket for Static Website
    const siteBucket = new s3.Bucket(this, 'SiteBucket', {
      publicReadAccess: false,
      blockPublicAccess: s3.BlockPublicAccess.BLOCK_ALL,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      autoDeleteObjects: true,
    });

    // 3. Origin Access Control for CloudFront
    const oac = new cloudfront.S3OriginAccessControl(this, 'SiteOAC', {
      signing: cloudfront.Signing.SIGV4_ALWAYS,
    });

    // 4. Go Lambda Backend
    const solverLambda = new GoFunction(this, 'SolverLambda', {
      entry: path.join(__dirname, '../../server'),
      runtime: cdk.aws_lambda.Runtime.PROVIDED_AL2023,
      architecture: cdk.aws_lambda.Architecture.ARM_64,
      timeout: cdk.Duration.seconds(30),
      memorySize: 256,
    });

    // 5. HTTP API (API Gateway v2)
    const httpApi = new apigwv2.HttpApi(this, 'SolverApi', {
      apiName: `VoltorbFlipApi-${props.stage}`,
    });

    httpApi.addRoutes({
      path: '/api/solve',
      methods: [apigwv2.HttpMethod.POST],
      integration: new HttpLambdaIntegration('SolverIntegration', solverLambda),
    });

    // 6. CloudFront Distribution
    const distribution = new cloudfront.Distribution(this, 'SiteDistribution', {
      certificate: certificate,
      domainNames: [domainName, ...subjectAlternativeNames],
      defaultRootObject: 'index.html',
      defaultBehavior: {
        origin: origins.S3BucketOrigin.withOriginAccessControl(siteBucket),
        viewerProtocolPolicy: cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
        cachePolicy: cloudfront.CachePolicy.CACHING_OPTIMIZED,
      },
      additionalBehaviors: {
        '/api/*': {
          origin: new origins.HttpOrigin(`${httpApi.apiId}.execute-api.${this.region}.${this.urlSuffix}`, {
            protocolPolicy: cloudfront.OriginProtocolPolicy.HTTPS_ONLY,
          }),
          viewerProtocolPolicy: cloudfront.ViewerProtocolPolicy.HTTPS_ONLY,
          cachePolicy: cloudfront.CachePolicy.CACHING_DISABLED,
          allowedMethods: cloudfront.AllowedMethods.ALLOW_ALL,
          originRequestPolicy: cloudfront.OriginRequestPolicy.ALL_VIEWER_EXCEPT_HOST_HEADER,
        },
      },
    });

    // 7. Route 53 A Records pointing to CloudFront
    new route53.ARecord(this, 'AliasRecord', {
      zone: zone,
      recordName: props.stage === 'prod' ? '' : 'beta',
      target: route53.RecordTarget.fromAlias(new targets.CloudFrontTarget(distribution)),
    });

    if (props.stage === 'prod') {
      new route53.ARecord(this, 'WwwAliasRecord', {
        zone: zone,
        recordName: 'www',
        target: route53.RecordTarget.fromAlias(new targets.CloudFrontTarget(distribution)),
      });
    }

    // 8. Deploy Jekyll Site
    new s3_deployment.BucketDeployment(this, 'DeployWebsite', {
      sources: [
        s3_deployment.Source.asset(path.join(__dirname, '../../jekyll-site'), {
          bundling: {
            image: cdk.DockerImage.fromRegistry('jekyll/jekyll:4.2.2'),
            command: [
              'sh', '-c',
              'rm -f Gemfile.lock && bundle install && jekyll build --destination /asset-output'
            ],
            user: 'root',
          },
        }),
      ],
      destinationBucket: siteBucket,
      distribution: distribution,
      distributionPaths: ['/*'],
    });

    new cdk.CfnOutput(this, 'DistributionDomainName', {
      value: distribution.distributionDomainName,
    });

    new cdk.CfnOutput(this, 'SiteUrl', {
      value: `https://${domainName}`,
    });
  }
}
