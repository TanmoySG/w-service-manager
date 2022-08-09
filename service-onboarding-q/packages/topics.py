from confluent_kafka.admin import AdminClient, NewTopic


class Topics:

    def __init__(self, kafkaClient: AdminClient) -> None:
        self.kafkaClient: AdminClient = kafkaClient

    def create(self, topics) -> None:
        if type(topics) == str:
            topics = [{"topic" : topics}]
        elif type(topics) != list:
            topics = [topics]

        newTopicsList = []
        for topic in topics:
            newTopicsList.append(
                NewTopic(
                    topic['topic'], 
                    topic.get("partitions", 1),
                    topic.get("replications", 1)
                )
            )

        topics = self.kafkaClient.create_topics(
            new_topics=newTopicsList,
            validate_only=False
        )

        result = []

        for topic, f in topics.items():
            try:
                f.result()  # The result itself is None
                result.append("Topic {} created".format(topic))
            except Exception as e:
                result.append("Failed to create topic {}: {}".format(topic, e))

        return result
    
    def delete(self, topics) -> None:
        if type(topics) != list:
            topics = [topics]
            
        deletedTopics = self.kafkaClient.delete_topics(topics, operation_timeout=30)
        
        result = []
        for topic in deletedTopics.items():
            try:
                f.result()  # The result itself is None
                result.append("Topic {} deleted".format(topic))
            except Exception as e:
                result.append("Failed to delete topic {}: {}".format(topic, e))

