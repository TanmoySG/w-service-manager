import fetch from "node-fetch";

export default class WDB {
    constructor(configurations) {
        this.configurations = configurations
        this.database = configurations.database
        this.wdbEndpoint = `https://wdb.tanmoysg.com/connect?cluster=${configurations.cluster}&token=${configurations.token}`;
    }

    getData(collection, callback) {
        fetch(this.wdbEndpoint, {
            method: "POST",
            body: JSON.stringify({
                "action": "get-data",
                "payload": {
                    "database": this.database,
                    "collection": collection
                }
            }),
            headers: { 'Content-Type': 'application/json' }
        }).then(function (response) {
            return response.json()
        }).then(function (response) {
            callback(response.data, response.schema)
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

    deleteData(collection, maker, callback) {
        fetch(this.wdbEndpoint, {
            method: "POST",
            body: JSON.stringify({
                "action": "delete-data",
                "payload": {
                    "database": this.database,
                    "collection": collection,
                    "marker": `${maker.Key} : ${maker.Value}`
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

    updateData(collection, maker, data,callback) {
        fetch(this.wdbEndpoint, {
            method: "POST",
            body: JSON.stringify({
                "action": "update-data",
                "payload": {
                    "database": this.database,
                    "collection": collection,
                    "marker": `${maker.Key} : ${maker.Value}`,
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