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

### Startup Configuration Guide

Currently only Topic Creation and Deletion at start-up is supported. 
The structure of the JSON should be
```json
{
    "topic" : {
        "create" : <either false or list of topics>,
        "delete" : <either false or list of topics>
    }
}
```