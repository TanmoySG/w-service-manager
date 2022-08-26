read -p "Schema Name (format: schema.namespace.parent-namespace): " SCHEMAID
read -p "Schema Filepath (format: path/to/schema.namespace.parent-namespace): " SCHEMAPATH

SCHEMAFULLPATH=$SCHEMAPATH/$SCHEMAID.schema.json

TEMPFILE=$(mktemp)

cat ./schema.template.json | jq 'del(.properties.complexField) | del(.required[1]) | .definitions={} | .' > $TEMPFILE
jq --arg SCHEMAID "${SCHEMAID}"  '.SchemaID=$SCHEMAID' < $TEMPFILE > $SCHEMAFULLPATH;

cat $SCHEMAFULLPATH > $TEMPFILE
jq --arg SCHEMAID "${SCHEMAID}"  '.properties.kind.default=$SCHEMAID | .properties.kind.enum[0]=$SCHEMAID' < $TEMPFILE > $SCHEMAFULLPATH

rm -rf $TEMPFILE