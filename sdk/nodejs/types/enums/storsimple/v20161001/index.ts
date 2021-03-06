// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const CloudType = {
    Azure: "Azure",
    S3: "S3",
    S3_RRS: "S3_RRS",
    OpenStack: "OpenStack",
    HP: "HP",
} as const;

/**
 * The cloud service provider
 */
export type CloudType = (typeof CloudType)[keyof typeof CloudType];

export const DataPolicy = {
    Invalid: "Invalid",
    Local: "Local",
    Tiered: "Tiered",
    Cloud: "Cloud",
} as const;

/**
 * The data policy.
 */
export type DataPolicy = (typeof DataPolicy)[keyof typeof DataPolicy];

export const DiskStatus = {
    Online: "Online",
    Offline: "Offline",
} as const;

/**
 * The disk status.
 */
export type DiskStatus = (typeof DiskStatus)[keyof typeof DiskStatus];

export const EncryptionAlgorithm = {
    None: "None",
    AES256: "AES256",
    RSAES_PKCS1_v_1_5: "RSAES_PKCS1_v_1_5",
} as const;

/**
 * Algorithm used to encrypt "Value"
 */
export type EncryptionAlgorithm = (typeof EncryptionAlgorithm)[keyof typeof EncryptionAlgorithm];

export const EncryptionStatus = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * The encryption status "Enabled | Disabled".
 */
export type EncryptionStatus = (typeof EncryptionStatus)[keyof typeof EncryptionStatus];

export const ManagerSkuType = {
    Standard: "Standard",
} as const;

/**
 * Refers to the sku name which should be "Standard"
 */
export type ManagerSkuType = (typeof ManagerSkuType)[keyof typeof ManagerSkuType];

export const ManagerType = {
    GardaV1: "GardaV1",
    HelsinkiV1: "HelsinkiV1",
} as const;

/**
 * Refers to the type of the StorSimple Manager
 */
export type ManagerType = (typeof ManagerType)[keyof typeof ManagerType];

export const MonitoringStatus = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * The monitoring.
 */
export type MonitoringStatus = (typeof MonitoringStatus)[keyof typeof MonitoringStatus];

export const ShareStatus = {
    Online: "Online",
    Offline: "Offline",
} as const;

/**
 * The Share Status
 */
export type ShareStatus = (typeof ShareStatus)[keyof typeof ShareStatus];

export const SslStatus = {
    Enabled: "Enabled",
    Disabled: "Disabled",
} as const;

/**
 * SSL needs to be enabled or not
 */
export type SslStatus = (typeof SslStatus)[keyof typeof SslStatus];
