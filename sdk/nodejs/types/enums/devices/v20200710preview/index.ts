// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AccessRights = {
    RegistryRead: "RegistryRead",
    RegistryWrite: "RegistryWrite",
    ServiceConnect: "ServiceConnect",
    DeviceConnect: "DeviceConnect",
    RegistryRead_RegistryWrite: "RegistryRead, RegistryWrite",
    RegistryRead_ServiceConnect: "RegistryRead, ServiceConnect",
    RegistryRead_DeviceConnect: "RegistryRead, DeviceConnect",
    RegistryWrite_ServiceConnect: "RegistryWrite, ServiceConnect",
    RegistryWrite_DeviceConnect: "RegistryWrite, DeviceConnect",
    ServiceConnect_DeviceConnect: "ServiceConnect, DeviceConnect",
    RegistryRead_RegistryWrite_ServiceConnect: "RegistryRead, RegistryWrite, ServiceConnect",
    RegistryRead_RegistryWrite_DeviceConnect: "RegistryRead, RegistryWrite, DeviceConnect",
    RegistryRead_ServiceConnect_DeviceConnect: "RegistryRead, ServiceConnect, DeviceConnect",
    RegistryWrite_ServiceConnect_DeviceConnect: "RegistryWrite, ServiceConnect, DeviceConnect",
    RegistryRead_RegistryWrite_ServiceConnect_DeviceConnect: "RegistryRead, RegistryWrite, ServiceConnect, DeviceConnect",
} as const;

/**
 * The permissions assigned to the shared access policy.
 */
export type AccessRights = (typeof AccessRights)[keyof typeof AccessRights];

export const AuthenticationType = {
    KeyBased: "keyBased",
    IdentityBased: "identityBased",
} as const;

/**
 * Specifies authentication type being used for connecting to the storage account.
 */
export type AuthenticationType = (typeof AuthenticationType)[keyof typeof AuthenticationType];

export const Capabilities = {
    None: "None",
    DeviceManagement: "DeviceManagement",
} as const;

/**
 * The capabilities and features enabled for the IoT hub.
 */
export type Capabilities = (typeof Capabilities)[keyof typeof Capabilities];

export const IotHubSku = {
    F1: "F1",
    S1: "S1",
    S2: "S2",
    S3: "S3",
    B1: "B1",
    B2: "B2",
    B3: "B3",
} as const;

/**
 * The name of the SKU.
 */
export type IotHubSku = (typeof IotHubSku)[keyof typeof IotHubSku];

export const IpFilterActionType = {
    Accept: "Accept",
    Reject: "Reject",
} as const;

/**
 * The desired action for requests captured by this rule.
 */
export type IpFilterActionType = (typeof IpFilterActionType)[keyof typeof IpFilterActionType];

export const PrivateLinkServiceConnectionStatus = {
    Pending: "Pending",
    Approved: "Approved",
    Rejected: "Rejected",
    Disconnected: "Disconnected",
} as const;

/**
 * The status of a private endpoint connection
 */
export type PrivateLinkServiceConnectionStatus = (typeof PrivateLinkServiceConnectionStatus)[keyof typeof PrivateLinkServiceConnectionStatus];

export const PublicNetworkAccess = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * Whether requests from Public Network are allowed
 */
export type PublicNetworkAccess = (typeof PublicNetworkAccess)[keyof typeof PublicNetworkAccess];

export const RoutingSource = {
    Invalid: "Invalid",
    DeviceMessages: "DeviceMessages",
    TwinChangeEvents: "TwinChangeEvents",
    DeviceLifecycleEvents: "DeviceLifecycleEvents",
    DeviceJobLifecycleEvents: "DeviceJobLifecycleEvents",
    DigitalTwinChangeEvents: "DigitalTwinChangeEvents",
} as const;

/**
 * The source that the routing rule is to be applied to, such as DeviceMessages.
 */
export type RoutingSource = (typeof RoutingSource)[keyof typeof RoutingSource];
