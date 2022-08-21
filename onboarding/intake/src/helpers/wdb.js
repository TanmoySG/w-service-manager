import fetch from "node-fetch";

export default class WDB {
    constructor(configurations) {
        this.configurations = configurations
        this.database = configurations.database
        this.wdbEndpoint = `https://wdb.tanmoysg.com/connect?cluster=${configurations.cluster}&token=${configurations.token}`;
    }

    getData(collection, marker, callback) {
        fetch(this.wdbEndpoint, {
            method: "POST",
            body: JSON.stringify({
                "action": "view-data",
                "payload": {
                    "database": this.database,
                    "collection": collection,
                    "marker": `${marker.Key} : ${marker.Value}`
                }
            }),
            headers: { 'Content-Type': 'application/json' }
        }).then(function (response) {
            return response.json()
        }).then(function (response) {
            if(response.status_code === '1'){
                callback(response.response, undefined)
            }else{
                callback({}, "Not Found")
            }
        }).catch(function (error) {
            callback(error)
        })
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

    deleteData(collection, marker, callback) {
        fetch(this.wdbEndpoint, {
            method: "POST",
            body: JSON.stringify({
                "action": "delete-data",
                "payload": {
                    "database": this.database,
                    "collection": collection,
                    "marker": `${marker.Key} : ${marker.Value}`
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

    updateData(collection, marker, data,callback) {
        fetch(this.wdbEndpoint, {
            method: "POST",
            body: JSON.stringify({
                "action": "update-data",
                "payload": {
                    "database": this.database,
                    "collection": collection,
                    "marker": `${marker.Key} : ${marker.Value}`,
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

}