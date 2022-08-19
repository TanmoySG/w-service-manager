import * as filesystem from 'fs';
import { applyPatch, createPatch } from 'rfc6902'

export default class ConfigParser {
    constructor(filepath) {
        this.configurations = JSON.parse(
            filesystem.readFileSync(filepath)
        )
        this.patchedConfigurations = JSON.parse(
            filesystem.readFileSync(filepath)
        )
        this.patchFile = undefined
    }


    new(filepath) {
        this.configurations = JSON.parse(
            filesystem.readFileSync(filepath)
        )
        this.patchedConfigurations = JSON.parse(
            filesystem.readFileSync(filepath)
        )
    }

    loadPatchFile(patchFilepath) {
        if (patchFilepath != undefined) {
            const patchFile = JSON.parse(
                filesystem.readFileSync(patchFilepath)
            )
            this.patchFile = patchFile.patches
        }
        return this
    }

    patch(patch = undefined) {
        if (patch === undefined && this.patchFile != undefined) {
            patch = this.patchFile
        }
        if (Array.isArray(patch) == false) {
            patch = [patch]
        }
        applyPatch(this.patchedConfigurations, [...patch])
        return this
    }

    getBase() {
        return this.configurations
    }

    getPatched(){
        return this.patchedConfigurations
    }

    createPatch(base, target){
        return createPatch(base, target)
    }
}


// const cp = new ConfigParser('/Users/tanmoysg/Work/Projects/wunder/w-service-manager/onboarding/intake/config/secrets/wdb.config.secret.json')
// cp.loadPatchFile('/Users/tanmoysg/Work/Projects/wunder/w-service-manager/onboarding/intake/config/secrets/patches/config.contract.patch.secrets.json').patch()
// console.log(cp.get())

