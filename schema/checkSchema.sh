SCHEMAFILEPATH=$1

HAS_SchemaID=$(jq 'has(".SchemaID")' < $SCHEMAFILEPATH)
if [ $HAS_SchemaID == 'false' ]
then 
    echo "\033[00;31m× SchemaID Missing"
else
    echo "\033[00;31m✔️ SchemaID Found"
fi


# HAS_Type=$(jq 'has(".type")' < $SCHEMAFILEPATH)
HAS_Properties=$(jq 'has(".properties")' < $SCHEMAFILEPATH)
if [ $HAS_Properties == false ]
then 
    echo "\033[00;31m× HAS_Properties Missing"
fi