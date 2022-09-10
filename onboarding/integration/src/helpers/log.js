import Logsmith from "logsmithjs";

export default class Log {
    constructor(configuration) {
        this.log = new Logsmith({});
        this.log.fetchConfigFromFile(configuration)
    }

    request(request, response) {
        const loggable = {}
        loggable["service"] = request["service_name"]
        if (response.error == undefined) {
            loggable["requestID"] = response["request-details"]["request_id"]
            loggable["status"] = response["request-details"]["request_status"]
            loggable["timestamp"] = response["request-details"]["timestamp"]
            this.log.SUCCESS(loggable)
            return loggable
        } else {
            loggable["error"] = response.error
            this.log.FAILURE(loggable)
            return loggable
        }
    }

    contract(request, response) {
        const loggable = {}
        loggable["service"] = request["service"]["name"]
        if (response.error === undefined) {
            loggable["contractID"] = response["contract_id"]
            loggable["requestID"] = response["request_id"]
            loggable["status"] = response["request_status"]
            loggable["timestamp"] = response["timestamp"]
            this.log.SUCCESS(loggable)
            return loggable
        } else {
            loggable["error"] = response.error
            this.log.FAILURE(loggable)
            return loggable
        }
    }
}