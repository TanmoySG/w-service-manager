from confluent_kafka.admin import AdminClient, NewTopic


class Topics:

    def __init__(self, kafkaClient: AdminClient) -> None:
        self.kafkaClient: AdminClient = kafkaClient

    def create(self, topic: str, partitions: int = 1, replication_factor: int = 1) -> None:
        newTopic = [NewTopic(topic, partitions, replication_factor)]
        topics = self.kafkaClient.create_topics(
            new_topics=newTopic,
            validate_only=False
        )

        for topic, f in topics.items():
            try:
                f.result()  # The result itself is None
                print("Topic {} created".format(topic))
            except Exception as e:
                print("Failed to create topic {}: {}".format(topic, e))

