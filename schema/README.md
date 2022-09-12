# Schema

To generate schema run
```
sh prepareSchema.sh
```
A prompt comes up
```
Schema Name (format: schema.namespace.parent-namespace): test
Schema Filepath (format: path/to/schema.namespace.parent-namespace): ./
```

This creates the schema.

## CodeGen

We use the [quicktype](https://github.com/quicktype/quicktype) tool for code generation. To make it easy to use we've written a short shell script that can be used to generate code. 
```
chmod 775 ./codegen.sh
./codegen path/to/schema.json path/to/code.go 
```

### In work
As a part of standardized schema usage, a schema-standard validation is also in work and will be added sometime in the future.