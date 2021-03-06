// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const BuildTaskStatus = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

/**
 * The current status of build task.
 */
export type BuildTaskStatus = (typeof BuildTaskStatus)[keyof typeof BuildTaskStatus];

export const OsType = {
    Windows: "Windows",
    Linux: "Linux",
} as const;

/**
 * The operating system type required for the build.
 */
export type OsType = (typeof OsType)[keyof typeof OsType];

export const SourceControlType = {
    Github: "Github",
    VisualStudioTeamService: "VisualStudioTeamService",
} as const;

/**
 * The type of source control service.
 */
export type SourceControlType = (typeof SourceControlType)[keyof typeof SourceControlType];

export const TokenType = {
    PAT: "PAT",
    OAuth: "OAuth",
} as const;

/**
 * The type of Auth token.
 */
export type TokenType = (typeof TokenType)[keyof typeof TokenType];
