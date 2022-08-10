import os

from configPy import JSONConfigParser
from flask import Flask, jsonify, request

from packages.admin import KAdmin
from packages.startup import StartUp
from packages.topics import Topics

app = Flask(__name__)

admin_client = KAdmin('./configuration.json')

def startup():
    startup_configuration_path = os.environ.get("SOQ_STARTUP_CONFIG", False)
    if startup_configuration_path != False:
        startup_configuration = JSONConfigParser(startup_configuration_path).getConfigurations()
        startup = StartUp(
            configuration=startup_configuration,
            admin_client=admin_client
        )
        result = startup.execute()
        print("| Service-Onboarding Q API Warm-Up |> \t {0}".format(result))

startup()

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


@app.route("/v1/topic/delete", methods=["POST"])
def topic_delete():
    request_data = request.get_json(force=True)
    if request_data["topics"] == None:
        return {"error": "Topic Name Required"}
    topics_list = request_data.get("topics")
    topics = Topics(admin_client)
    response = topics.delete(topics=topics_list)
    return jsonify({
        "message": response
    })


if __name__ == '__main__':
    app.run()
