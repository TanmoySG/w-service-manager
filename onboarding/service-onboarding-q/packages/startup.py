from packages.topics import Topics

class StartUp:

    def __init__(self, configuration, admin_client) -> None:
        self.admin_client = admin_client
        self.configuration = configuration
        self.result = {
            "topics": {
                "create" : None,
                "delete" : None
            }
        }
        
    def createTopicOnStartUp(self, topics) -> None:
        t = Topics(self.admin_client)
        return t.create(topics=topics)

    def deleteTopicOnStartUp(self, topics) -> None:
        t = Topics(self.admin_client)
        return t.delete(topics=topics)
        
    def topicRunner(self, topicConfig):
        if topicConfig.get("create", False) != False:
            self.result["topics"]["create"] = self.createTopicOnStartUp(topics=topicConfig["create"])
            
        if topicConfig.get("delete", False) != False:
            self.result["topics"]["delete"] = self.deleteTopicOnStartUp(topics=topicConfig["delete"])

        return self.result

    def execute(self) -> None:
        if self.configuration.get("topic", False) != False:
            topicsStartupConfig = self.configuration["topic"]
            self.result= self.topicRunner(topicsStartupConfig)

        return self.result
