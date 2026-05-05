import * as cdk from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import { VoltorbFlipDotComStack } from '../lib/voltorb-flip-dot-com-stack';

test('Beta Stack Resources Created', () => {
  const app = new cdk.App();
  const stack = new VoltorbFlipDotComStack(app, 'BetaStack', {
    env: { account: '123456789012', region: 'us-east-1' },
    stage: 'beta'
  });
  const template = Template.fromStack(stack);

  template.resourceCountIs('AWS::S3::Bucket', 1);
  template.resourceCountIs('AWS::CloudFront::Distribution', 1);
  template.hasResource('AWS::Lambda::Function', {});
});

test('Prod Stack Resources Created', () => {
  const app = new cdk.App();
  const stack = new VoltorbFlipDotComStack(app, 'ProdStack', {
    env: { account: '123456789012', region: 'us-east-1' },
    stage: 'prod'
  });
  const template = Template.fromStack(stack);

  template.resourceCountIs('AWS::S3::Bucket', 1);
  template.resourceCountIs('AWS::CloudFront::Distribution', 1);
  template.hasResource('AWS::Lambda::Function', {});
});
