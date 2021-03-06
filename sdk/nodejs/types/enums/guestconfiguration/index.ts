// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as latest from "./latest";
import * as v20180630preview from "./v20180630preview";
import * as v20181120 from "./v20181120";
import * as v20200625 from "./v20200625";

export {
    latest,
    v20180630preview,
    v20181120,
    v20200625,
};

export const ActionAfterReboot = {
    ContinueConfiguration: "ContinueConfiguration",
    StopConfiguration: "StopConfiguration",
} as const;

/**
 * Specifies what happens after a reboot during the application of a configuration. The possible values are ContinueConfiguration and StopConfiguration
 */
export type ActionAfterReboot = (typeof ActionAfterReboot)[keyof typeof ActionAfterReboot];

export const AllowModuleOverwrite = {
    True: "True",
    False: "False",
} as const;

/**
 * If true - new configurations downloaded from the pull service are allowed to overwrite the old ones on the target node. Otherwise, false
 */
export type AllowModuleOverwrite = (typeof AllowModuleOverwrite)[keyof typeof AllowModuleOverwrite];

export const ConfigurationMode = {
    ApplyOnly: "ApplyOnly",
    ApplyAndMonitor: "ApplyAndMonitor",
    ApplyAndAutoCorrect: "ApplyAndAutoCorrect",
} as const;

/**
 * Specifies how the LCM(Local Configuration Manager) actually applies the configuration to the target nodes. Possible values are ApplyOnly, ApplyAndMonitor, and ApplyAndAutoCorrect.
 */
export type ConfigurationMode = (typeof ConfigurationMode)[keyof typeof ConfigurationMode];

export const Kind = {
    DSC: "DSC",
} as const;

/**
 * Kind of the guest configuration. For example:DSC
 */
export type Kind = (typeof Kind)[keyof typeof Kind];

export const RebootIfNeeded = {
    True: "True",
    False: "False",
} as const;

/**
 * Set this to true to automatically reboot the node after a configuration that requires reboot is applied. Otherwise, you will have to manually reboot the node for any configuration that requires it. The default value is false. To use this setting when a reboot condition is enacted by something other than DSC (such as Windows Installer), combine this setting with the xPendingReboot module.
 */
export type RebootIfNeeded = (typeof RebootIfNeeded)[keyof typeof RebootIfNeeded];
