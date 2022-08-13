from flask import Flask

app = Flask(__name__)

@app.route("/api/v1", methods=["GET"])
def apiv1():
    return {
        "message" : "Welcome to Service Manager API v1",
        "service" : "Onboarding API - Intake v1",
        "version" : "v1[1.0.0]"
    }
    

@app.route("/api/v1/", methods=["GET"])
def apiv1():
    return {
        "message": "Welcome to Service Manager API v1",
        "service": "Onboarding API - Intake v1",
        "version": "v1[1.0.0]"
    }

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8081)