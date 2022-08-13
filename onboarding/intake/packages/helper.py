from jsonschema import validate
from configPy import JSONConfigParser

class JSONSchema:
        
    def __init__(self, schemaFilePath) -> None:
        self.schema = JSONConfigParser(schemaFilePath).getConfigurations()
        pass
        
    def validate(self, data):
        validate(instance=data, schema=self.schema)
        return