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
    request_data = request.get_json(force=True)
    if request_data["topic"] == None:
        return {"error": "Topic Name Required"}
    topic_name = request_data.get("topic")
    topic_partitions = request_data.get("partitions", 1)
    topic_replications = request_data.get("replications", 1)
    topics = Topics(admin_client)
    response = topics.create(
        topic=topic_name,
        partitions=topic_partitions, 
        replication_factor=topic_replications
    )
    return jsonify({
        "message": response
    })


if __name__ == '__main__':
    app.run()
