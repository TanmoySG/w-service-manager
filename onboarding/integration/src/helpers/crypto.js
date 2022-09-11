import * as crypto from "crypto";

export default class RSA {
    constructor(){}

    getKeys(){
        return crypto.generateKeyPairSync("rsa", {
            // The standard secure default length for RSA keys is 2048 bits
            modulusLength: 2048,
        })
    }
}