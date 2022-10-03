#!/bin/bash
set -x

awslocal sqs create-queue \
  --queue-name voip-push-queue.fifo \
  --attributes FifoQueue=true

set +x
