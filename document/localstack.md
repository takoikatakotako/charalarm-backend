
aws dynamodb list-tables --endpoint-url=http://localhost:4566

aws dynamodb create-table \
    --table-name user-table \
    --attribute-definitions AttributeName=userId,AttributeType=S \
    --key-schema AttributeName=userId,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --endpoint-url=http://localhost:4566

aws dynamodb describe-table --table-name user-table --endpoint-url=http://localhost:4566

aws dynamodb delete-table --table-name user-table --endpoint-url=http://localhost:4566

