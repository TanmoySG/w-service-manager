from configPy import JSONConfigParser
from confluent_kafka.admin import AdminClient


class KAdmin:
    def __new__(cls, kafkaConfig):
        """Create New KafkaAdmin Object

        Args:
            kafkaConfig (str): Path to Kafka Config JSON

        Returns:
            AdminClient: A KafkaAdminClient Object
        """
        parsedKafkaConfig = JSONConfigParser(kafkaConfig).getConfigurations()
        return AdminClient(parsedKafkaConfig)
