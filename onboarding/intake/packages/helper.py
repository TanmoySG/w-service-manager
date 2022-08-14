import json

from configPy import JSONConfigParser
from confluent_kafka import Producer
from jsonschema import exceptions, validate

class JSONSchema:

    def __init__(self, schemaFilePath) -> None:
        self.schema = JSONConfigParser(schemaFilePath).getConfigurations()
        pass

    def validate(self, data):
        try:
            validate(instance=data, schema=self.schema)
        except exceptions.ValidationError:
            return False
        else:
            return True


def acked(err, msg):
    global delivered_records
    """
    Delivery report handler called on
    successful or failed delivery of message
    """
    if err is not None:
        print("Failed to deliver message: {}".format(err))
    else:
        delivered_records += 1
        print("Produced record to topic {} partition [{}] @ offset {}"
              .format(msg.topic(), msg.partition(), msg.offset()))


class Kafka:
    def __init__(self, configuration, topic) -> None:
        self.configuration = configuration
        self.topic = topic
        self.producer = Producer(**self.configuration)
        # self.consumer = Consumer(**self.configuration)
        pass

    def produce(self, key, data):
        # Produce line (without newline)
        self.producer.produce(self.topic, key=key, value=json.dumps(
            data).encode("utf-8"), on_delivery=acked)
        self.producer.poll(0)

    def consume(self):
        pass

# k = Kafka(configuration={
#     "bootstrap.servers": "localhost:9092",
#     }, topic="intake")
# k.produce(key="1", data={"test": "ok"})
