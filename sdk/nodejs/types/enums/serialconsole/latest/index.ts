// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const SerialPortState = {
    Enabled: "enabled",
    Disabled: "disabled",
} as const;

/**
 * Specifies whether the port is enabled for a serial console connection.
 */
export type SerialPortState = (typeof SerialPortState)[keyof typeof SerialPortState];
