from flask import Flask, request, jsonify

from packages.topics import Topics
from packages.admin import KAdmin

app = Flask(__name__)

admin_client = KAdmin('./configuration.json')

@app.route("/v1/api", methods=["GET"])
def api():
    return jsonify({
        "message": "Welcome to Service Onboarding Q API v1"
    })


@app.route("/v1/topic/create", methods=["POST"])
def topic_create():
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
