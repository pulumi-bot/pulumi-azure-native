// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const EnforcementMode = {
    /**
     * The policy effect is enforced during resource creation or update.
     */
    Default: "Default",
    /**
     * The policy effect is not enforced during resource creation or update.
     */
    DoNotEnforce: "DoNotEnforce",
} as const;

/**
 * The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
 */
export type EnforcementMode = (typeof EnforcementMode)[keyof typeof EnforcementMode];

export const LockLevel = {
    NotSpecified: "NotSpecified",
    CanNotDelete: "CanNotDelete",
    ReadOnly: "ReadOnly",
} as const;

/**
 * The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
 */
export type LockLevel = (typeof LockLevel)[keyof typeof LockLevel];

export const ParameterType = {
    String: "String",
    Array: "Array",
    Object: "Object",
    Boolean: "Boolean",
    Integer: "Integer",
    Float: "Float",
    DateTime: "DateTime",
} as const;

/**
 * The data type of the parameter.
 */
export type ParameterType = (typeof ParameterType)[keyof typeof ParameterType];

export const PolicyType = {
    NotSpecified: "NotSpecified",
    BuiltIn: "BuiltIn",
    Custom: "Custom",
    Static: "Static",
} as const;

/**
 * The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
 */
export type PolicyType = (typeof PolicyType)[keyof typeof PolicyType];

export const ResourceIdentityType = {
    /**
     * Indicates that a system assigned identity is associated with the resource.
     */
    SystemAssigned: "SystemAssigned",
    /**
     * Indicates that no identity is associated with the resource or that the existing identity should be removed.
     */
    None: "None",
} as const;

/**
 * The identity type. This is the only required field when adding a system assigned identity to a resource.
 */
export type ResourceIdentityType = (typeof ResourceIdentityType)[keyof typeof ResourceIdentityType];
