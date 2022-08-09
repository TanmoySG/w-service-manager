from configPy import JSONConfigParser
from confluent_kafka.admin import AdminClient
from flask import Flask, request, jsonify
# from flask_cors import CORS

from packages.topics import Topics

app = Flask(__name__)
# CORS(app)

kafkaConfig = JSONConfigParser("./configuration.json").getConfigurations()
admin_client = AdminClient(kafkaConfig)


@app.route("/v1/api", methods=["GET"])
def api():
    return jsonify({
        "message": "Welcome to Service Onboarding Q API v1"
    })


@app.route("/v1/topic/create", methods=["POST"])
def topic_create():
    """
    GET /v1/topic/create
    {
        "topics" : [
            {"topic" : "xyz", "partitions" : 1, "replications" : 1}
        ]
    }
    """
    request_data = request.get_json(force=True)
    if request_data["topics"] == None:
        return {"error": "Topic Name Required"}
    topics_list = request_data.get("topics")
    topics = Topics(admin_client)
    response = topics.create(topics=topics_list)
    return jsonify({
        "message": response
    })


if __name__ == '__main__':
    app.run()
