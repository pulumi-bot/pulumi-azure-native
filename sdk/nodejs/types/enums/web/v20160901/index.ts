// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AccessControlEntryAction = {
    Permit: "Permit",
    Deny: "Deny",
} as const;

/**
 * Action object.
 */
export type AccessControlEntryAction = (typeof AccessControlEntryAction)[keyof typeof AccessControlEntryAction];

export const ComputeModeOptions = {
    Shared: "Shared",
    Dedicated: "Dedicated",
    Dynamic: "Dynamic",
} as const;

/**
 * Shared or dedicated app hosting.
 */
export type ComputeModeOptions = (typeof ComputeModeOptions)[keyof typeof ComputeModeOptions];

export const InternalLoadBalancingMode = {
    None: "None",
    Web: "Web",
    Publishing: "Publishing",
} as const;

/**
 * Specifies which endpoints to serve internally in the Virtual Network for the App Service Environment.
 */
export type InternalLoadBalancingMode = (typeof InternalLoadBalancingMode)[keyof typeof InternalLoadBalancingMode];

export const RouteType = {
    DEFAULT: "DEFAULT",
    INHERITED: "INHERITED",
    STATIC: "STATIC",
} as const;

/**
 * The type of route this is:
 * DEFAULT - By default, every app has routes to the local address ranges specified by RFC1918
 * INHERITED - Routes inherited from the real Virtual Network routes
 * STATIC - Static route set on the app only
 *
 * These values will be used for syncing an app's routes with those from a Virtual Network.
 */
export type RouteType = (typeof RouteType)[keyof typeof RouteType];
