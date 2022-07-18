#

## examples
```bash

aws dynamodb query \
 --table-name PKSK \
 --index-name GSI \
 --key-condition-expression "GSI = :name" \
 --expression-attribute-values '{":name":{"S":"GSI-search"}}'

```

## Videos

https://www.youtube.com/watch?v=fiP2e-g-r4g
