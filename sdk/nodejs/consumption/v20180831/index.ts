// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./budget";
export * from "./budgetByResourceGroupName";
export * from "./getBudget";
export * from "./getBudgetByResourceGroupName";

// Export enums:
export * from "../../types/enums/consumption/v20180831";

// Import resources to register:
import { Budget } from "./budget";
import { BudgetByResourceGroupName } from "./budgetByResourceGroupName";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:consumption/v20180831:Budget":
                return new Budget(name, <any>undefined, { urn })
            case "azure-native:consumption/v20180831:BudgetByResourceGroupName":
                return new BudgetByResourceGroupName(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "consumption/v20180831", _module)
