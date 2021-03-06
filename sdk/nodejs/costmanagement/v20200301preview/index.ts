// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./costAllocationRule";
export * from "./getCostAllocationRule";

// Export enums:
export * from "../../types/enums/costmanagement/v20200301preview";

// Import resources to register:
import { CostAllocationRule } from "./costAllocationRule";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:costmanagement/v20200301preview:CostAllocationRule":
                return new CostAllocationRule(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "costmanagement/v20200301preview", _module)
