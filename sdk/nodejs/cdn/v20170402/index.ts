// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./customDomain";
export * from "./endpoint";
export * from "./getCustomDomain";
export * from "./getEndpoint";
export * from "./getProfile";
export * from "./getProfileSupportedOptimizationTypes";
export * from "./profile";

// Export enums:
export * from "../../types/enums/cdn/v20170402";

// Import resources to register:
import { CustomDomain } from "./customDomain";
import { Endpoint } from "./endpoint";
import { Profile } from "./profile";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:cdn/v20170402:CustomDomain":
                return new CustomDomain(name, <any>undefined, { urn })
            case "azure-native:cdn/v20170402:Endpoint":
                return new Endpoint(name, <any>undefined, { urn })
            case "azure-native:cdn/v20170402:Profile":
                return new Profile(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "cdn/v20170402", _module)
