TMPFILE=$(mktemp)
TMPOAPIFILE=$(mktemp)

echo '{"components":{"schemas":null}}' > $TMPFILE

SCHEMA=$(cat ../../../schema/service-onboarding/validity.audit.schema.json | jq '.components.schemas')
jq --argjson SCHEMA "${SCHEMA}"  '.components.schemas=$SCHEMA' < $TMPFILE > $TMPOAPIFILE

oapi-codegen -package spec -o spec.go -generate types,spec,skip-prune --old-config-style $TMPOAPIFILE

rm $TMPFILE
rm $TMPOAPIFILE

