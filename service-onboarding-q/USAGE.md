# Service Q API Usage Documentation

Service Q API Documentation for Usage Instructions.

## Topics API

Here are some of the Kafka Topic Related APIs and their usage.

### Create Topics
```
GET /v1/topic/create
    {
        "topics" : [
            {"topic" : "xyz", "partitions" : 1, "replications" : 1}
        ]
    }
```