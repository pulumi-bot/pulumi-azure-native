// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getHyperVSite";
export * from "./getSite";
export * from "./hyperVSite";
export * from "./site";

// Import resources to register:
import { HyperVSite } from "./hyperVSite";
import { Site } from "./site";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:offazure/v20200101:HyperVSite":
                return new HyperVSite(name, <any>undefined, { urn })
            case "azure-native:offazure/v20200101:Site":
                return new Site(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "offazure/v20200101", _module)
