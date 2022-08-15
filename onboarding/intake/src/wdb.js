import fetch from "node-fetch";

export default class WDB {
    constructor(configurations) {
        this.configurations = configurations
        this.database = configurations.database
        this.wdbEndpoint = `https://wdb.tanmoysg.com/connect?cluster=${configurations.cluster}&token=${configurations.token}`;
    }

    addData(collection, data, callback) {
        fetch(this.wdbEndpoint, {
            method: "POST",
            body: JSON.stringify({
                "action": "add-data",
                "payload": {
                    "database": this.database,
                    "collection": collection,
                    "data": data
                }
            }),
            headers: { 'Content-Type': 'application/json' }
        }).then(function (response) {
            return response.json()
        }).then(function (response) {
            callback(response)
        }).catch(function (error) {
            callback(error)
        })
    }

    deleteData(collection, data, callback) {

    }
}
// const wdb = new WDB({
//     "cluster": "",
//     "token": "",
//     "database": "",
//     "colllection": ""
// })

// const sampledata = {
//     "requestID": "1234",
//     "requestType": "create",
//     "serviceName": "xyz",
//     "serviceOwners": ["t", "r"],
//     "status": "created",
//     "timestamp": "123456"
// }

// wdb.addData("IntakeRequest-Stage", sampledata, function (resp) {
//     console.log(resp)
// })