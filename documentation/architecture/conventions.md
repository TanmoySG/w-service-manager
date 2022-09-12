## Conventions

These conventions should be followed while developing, but also should be revisited frequently as they might change until a solid set of conventions can be formulated.

### Directory Structure

The different components should have their own directory. For eg.
```
- w-service-manager
  - onboarding
    - ...
  - management
    - ...
  - provision
    - ...
```

### Schema Definations

Schemas Definations contain the schemas required to perform any valid request. Certain conventions to be followed (subject to changes)
- Schema Files should be defined in JSON format preferably, though we need YAML Schema as well in future.
- Schema Files should be stored in [`schema` directory](../../schema/)
- Schema Files should be named in the following conventions
```
<schema-name>.schema.json
<schema-name>.schema.yaml

# schema groups can be collective namespace for schemas of simillar kind or origin.
<schema-name>.<schema-group>.schema.json 
...
```
- Defining a Schema should follow - [TBD]

Do not add the following fields in JSON Schema File
```
    "$schema": "https://json-schema.org/draft/2020-12/schema",
    "$id": "https://github.com/TanmoySG/w-service-manager/blob/service-onboarding/schema/service-onboarding/request.intake.schema.json",
    "title": "New Service Onboarding Request",
    "description": "Service Onboarding Request Schema",
```

#### Schema Mapping
Schemas defined in the schema directory needs to be mapped in a schema.mapping.json file.
- The Schema Mapping JSON has the schema names, schema groups 
- These are mapped to their respective schema JSONs in the same directory
- Schema Groups are mapped to all their schemas and the schemas mapped to the files.

Example
```json
{
    "shopper" : "schema/shopper.schema.json",
    "product": {
        "new-product" : "schema/new-product.product.schema.json",
        "product-price" : "schema/product-price.product.schema.json"
    }
}
```

#### Open Questions

- Need to figure out defining a standard for schema definition
- Might need to create a "custom" template to define the schema
- The Standard Schema Definition should also be interchangeable with JSON Schema or OpenAPI Spec
- Should support Code Generation (reqd. for Go, etc.)
- **Require a Standard** - a task that can be carried parallely with [#6](https://github.com/TanmoySG/w-service-manager/issues/6)

### JSON Patching

A Patch Helper Class with `replace` , `remove` and other JSON Patching Methods.
- https://jsonpatch.com/
- Constructor uses - 
  - `base` configuration for initial JSON, no patching
  - `patched` configuration for patched JSON
- `mergePatch` method merges patch to base
- `reset` method resets the patched variable to base
- All Operations have a method. [Ref.](https://jsonpatch.com/#operations)