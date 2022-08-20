import express from "express";
import Request from "./src/intake.js";
import Contract from "./src/contract.js";
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

const contractConfig = new ConfigParser(configurations.base).loadPatchFile(configurations.contract.patch).patch().getPatched()
const requestConfig = new ConfigParser(configurations.base).loadPatchFile(configurations.request.patch).patch().getPatched()

const onboardingRequest = new Request(requestConfig)
const onboardingContract = new Contract(contractConfig)


app.use(express.json());

app.get("/", function (request, response) {
    response.send({
        message: "Welcome to Logsmith Monitor!",
        version: "logsmith-monitor v0.1.0-alpha"
    })
})

app.post("/v1/intake/request", function (request, response) {
    const requestData = request.body;
    onboardingRequest.new(requestData, onboardingRequest.id(), function (onboardingResponse) {
        response.send(onboardingResponse)
    })
})


app.post("/v1/intake/contract", function (request, response) {
    const contract = request.body;
    onboardingContract.register(contract, onboardingContract.id(), function (onboardingResponse) {
        response.json(onboardingResponse)
    })
})

app.listen(PORT, function () {
    console.log("Running on http://localhost:" + PORT)
})
