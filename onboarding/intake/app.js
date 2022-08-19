import express from "express";
import Request from "./src/intake.js";
import ConfigParser from "./src/helpers/config.js";

const configurations = {
    base: 'onboarding/intake/config/secrets/config.secrets.json',
    contract: {
        patch: 'onboarding/intake/config/secrets/patches/config.contract.patch.secrets.json'
    },
    request: {
        patch: 'onboarding/intake/config/secrets/patches/config.request.patch.secrets.json'
    }
}

const app = express();
const PORT = process.env.PORT || 8080;

const config = new ConfigParser(configurations.base)
const contractConfig = config.loadPatchFile(configurations.contract.patch).patch().getPatched()
const requestConfig = config.loadPatchFile(configurations.request.patch).patch().getPatched()

const onboardingRequest = new Request(requestConfig)

app.use(express.json());

app.get("/", function (request, response) {
    response.send({
        message: "Welcome to Logsmith Monitor!",
        version: "logsmith-monitor v0.1.0-alpha"
    })
})

app.post("/v1/intake/request", function (request, response) {
    const requestData = request.body;
    onboardingRequest.new(requestData, onboardingRequest.id(), function (onboardingResponse){
        response.send(onboardingResponse)
    })
})

app.listen(PORT, function () {
    console.log("Running on http://localhost:" + PORT)
})
