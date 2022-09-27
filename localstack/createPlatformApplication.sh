#!/bin/bash
set -x

awslocal sns create-platform-application \
    --name ios-voip-push-platform-application \
    --platform APNS \
    --attributes PlatformCredential=DAMMY

set +x


