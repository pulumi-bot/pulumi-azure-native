// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getSpatialAnchorsAccount";
export * from "./listSpatialAnchorsAccountKeys";
export * from "./spatialAnchorsAccount";

// Export enums:
export * from "../../types/enums/mixedreality/v20200501";

// Import resources to register:
import { SpatialAnchorsAccount } from "./spatialAnchorsAccount";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:mixedreality/v20200501:SpatialAnchorsAccount":
                return new SpatialAnchorsAccount(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "mixedreality/v20200501", _module)
