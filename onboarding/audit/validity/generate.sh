SCHEMAFILE=$1
GOFILE=$2

TMPFILE=$(mktemp)
TMPOAPIFILE=$(mktemp)

echo '{"components":{"schemas":null}}' > $TMPFILE

SCHEMA=$(cat ${SCHEMAFILE} | jq '.components.schemas')
jq --argjson SCHEMA "${SCHEMA}"  '.components.schemas=$SCHEMA' < $TMPFILE > $TMPOAPIFILE

oapi-codegen -package spec -o $GOFILE -generate types,spec,skip-prune --old-config-style $TMPOAPIFILE

rm $TMPFILE
rm $TMPOAPIFILE

