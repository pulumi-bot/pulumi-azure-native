// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getClientConfig";
export * from "./getClientToken";
export * from "./getManagementLockAtResourceGroupLevel";
export * from "./getManagementLockAtResourceLevel";
export * from "./getManagementLockAtSubscriptionLevel";
export * from "./getManagementLockByScope";
export * from "./getPolicyAssignment";
export * from "./getPolicyDefinition";
export * from "./getPolicySetDefinition";
export * from "./getRoleAssignment";
export * from "./getRoleDefinition";
export * from "./managementLockAtResourceGroupLevel";
export * from "./managementLockAtResourceLevel";
export * from "./managementLockAtSubscriptionLevel";
export * from "./managementLockByScope";
export * from "./policyAssignment";
export * from "./policyDefinition";
export * from "./policySetDefinition";
export * from "./roleAssignment";
export * from "./roleDefinition";

// Export enums:
export * from "../../types/enums/authorization/latest";

// Import resources to register:
import { ManagementLockAtResourceGroupLevel } from "./managementLockAtResourceGroupLevel";
import { ManagementLockAtResourceLevel } from "./managementLockAtResourceLevel";
import { ManagementLockAtSubscriptionLevel } from "./managementLockAtSubscriptionLevel";
import { ManagementLockByScope } from "./managementLockByScope";
import { PolicyAssignment } from "./policyAssignment";
import { PolicyDefinition } from "./policyDefinition";
import { PolicySetDefinition } from "./policySetDefinition";
import { RoleAssignment } from "./roleAssignment";
import { RoleDefinition } from "./roleDefinition";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-nextgen:authorization/latest:ManagementLockAtResourceGroupLevel":
                return new ManagementLockAtResourceGroupLevel(name, <any>undefined, { urn })
            case "azure-nextgen:authorization/latest:ManagementLockAtResourceLevel":
                return new ManagementLockAtResourceLevel(name, <any>undefined, { urn })
            case "azure-nextgen:authorization/latest:ManagementLockAtSubscriptionLevel":
                return new ManagementLockAtSubscriptionLevel(name, <any>undefined, { urn })
            case "azure-nextgen:authorization/latest:ManagementLockByScope":
                return new ManagementLockByScope(name, <any>undefined, { urn })
            case "azure-nextgen:authorization/latest:PolicyAssignment":
                return new PolicyAssignment(name, <any>undefined, { urn })
            case "azure-nextgen:authorization/latest:PolicyDefinition":
                return new PolicyDefinition(name, <any>undefined, { urn })
            case "azure-nextgen:authorization/latest:PolicySetDefinition":
                return new PolicySetDefinition(name, <any>undefined, { urn })
            case "azure-nextgen:authorization/latest:RoleAssignment":
                return new RoleAssignment(name, <any>undefined, { urn })
            case "azure-nextgen:authorization/latest:RoleDefinition":
                return new RoleDefinition(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-nextgen", "authorization/latest", _module)
