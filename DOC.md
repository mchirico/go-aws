#

## examples
```bash

aws dynamodb query \
 --table-name PKSK \
 --index-name GSI \
 --key-condition-expression "GSI = :name" \
 --expression-attribute-values '{":name":{"S":"GSI-search"}}'


aws dynamodb update-item \
    --table-name pksk \
    --key '{"Id":{"N":"789"}}' \
    --update-expression "SET RelatedItems = :ri, ProductReviews = :pr" \
    --expression-attribute-values file://values.json \
    --return-values ALL_NEW
```

## Videos

https://www.youtube.com/watch?v=fiP2e-g-r4g
