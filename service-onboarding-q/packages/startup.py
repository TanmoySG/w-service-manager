from topics import Topics
from admin import KAdmin

# admin_client = KAdmin('./configuration.json')

class STARTUP:
    
    def __init__(self, configuration, admin_client=None) -> None:
        self.admin_client = admin_client or KAdmin('./configuration.json')
        self.configuration = configuration
        self.result = {
            "createTopic" : None
        }
        
    def execute(self) -> None:
        topics = self.configuration["create_topics"]
        self.result['createTopic'] = self.createTopicOnStartUp(topics)
        return self.result
    
    def createTopicOnStartUp(self, topics) -> None:
        t = Topics(self.admin_client)
        return t.create(topics=topics)
    
sampleConfig = {
    'create_topics' : [
        {
            "topic" : "testxyz0987"
        },
        {
            "topic" : "testxyz087654"
        }
    ]
}
s = STARTUP(sampleConfig)
print(s.execute())