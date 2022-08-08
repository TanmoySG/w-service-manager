from confluent_kafka.admin import AdminClient, NewTopic
from packages.topics import Topics
from configPy import JSONConfigParser

kafkaConfig = JSONConfigParser(configFilePath="./configuration.json").getConfigurations()

admin_client = AdminClient(kafkaConfig)

topics = Topics(admin_client)
topics.create(topic='test-topic-3')
