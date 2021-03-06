// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AccessRights = {
    Manage: "Manage",
    Send: "Send",
    Listen: "Listen",
} as const;

export type AccessRights = (typeof AccessRights)[keyof typeof AccessRights];

export const Relaytype = {
    NetTcp: "NetTcp",
    Http: "Http",
} as const;

/**
 * WCFRelay Type.
 */
export type Relaytype = (typeof Relaytype)[keyof typeof Relaytype];

export const SkuName = {
    Standard: "Standard",
} as const;

/**
 * Name of this Sku
 */
export type SkuName = (typeof SkuName)[keyof typeof SkuName];

export const SkuTier = {
    Standard: "Standard",
} as const;

/**
 * The tier of this particular SKU
 */
export type SkuTier = (typeof SkuTier)[keyof typeof SkuTier];
