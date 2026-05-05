#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib/core';
import { VoltorbFlipDotComStack } from '../lib/voltorb-flip-dot-com-stack';

const app = new cdk.App();

const commonProps = {
  env: { 
    account: '197408819919', 
    region: 'us-east-1' 
  },
};

// Beta Stack (beta.voltorbflip.com)
new VoltorbFlipDotComStack(app, 'VoltorbFlipStack-Beta', {
  ...commonProps,
  stage: 'beta',
});

// Production Stack (voltorbflip.com)
new VoltorbFlipDotComStack(app, 'VoltorbFlipStack-Prod', {
  ...commonProps,
  stage: 'prod',
});
