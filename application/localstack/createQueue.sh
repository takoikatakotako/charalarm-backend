#!/bin/bash
set -x
awslocal sqs create-queue --queue-name voip-push-queue.fifo --region ap-northeast-1
set +x
