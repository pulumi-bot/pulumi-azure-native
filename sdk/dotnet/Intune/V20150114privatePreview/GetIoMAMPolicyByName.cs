// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Intune.V20150114privatePreview
{
    public static class GetIoMAMPolicyByName
    {
        public static Task<GetIoMAMPolicyByNameResult> InvokeAsync(GetIoMAMPolicyByNameArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIoMAMPolicyByNameResult>("azurerm:intune/v20150114privatepreview:getIoMAMPolicyByName", args ?? new GetIoMAMPolicyByNameArgs(), options.WithVersion());
    }


    public sealed class GetIoMAMPolicyByNameArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Location hostName for the tenant
        /// </summary>
        [Input("hostName", required: true)]
        public string HostName { get; set; } = null!;

        /// <summary>
        /// Unique name for the policy
        /// </summary>
        [Input("policyName", required: true)]
        public string PolicyName { get; set; } = null!;

        /// <summary>
        /// select specific fields in entity.
        /// </summary>
        [Input("select")]
        public string? Select { get; set; }

        public GetIoMAMPolicyByNameArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIoMAMPolicyByNameResult
    {
        public readonly string? AccessRecheckOfflineTimeout;
        public readonly string? AccessRecheckOnlineTimeout;
        public readonly string? AppSharingFromLevel;
        public readonly string? AppSharingToLevel;
        public readonly string? Authentication;
        public readonly string? ClipboardSharingLevel;
        public readonly string? DataBackup;
        public readonly string? Description;
        public readonly string? DeviceCompliance;
        public readonly string? FileEncryptionLevel;
        public readonly string? FileSharingSaveAs;
        public readonly string FriendlyName;
        public readonly string GroupStatus;
        public readonly string LastModifiedTime;
        /// <summary>
        /// Resource Location
        /// </summary>
        public readonly string? Location;
        public readonly string? ManagedBrowser;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        public readonly int NumOfApps;
        public readonly string? OfflineWipeTimeout;
        public readonly string? Pin;
        public readonly int? PinNumRetry;
        /// <summary>
        /// Resource Tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        public readonly string? TouchId;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetIoMAMPolicyByNameResult(
            string? accessRecheckOfflineTimeout,

            string? accessRecheckOnlineTimeout,

            string? appSharingFromLevel,

            string? appSharingToLevel,

            string? authentication,

            string? clipboardSharingLevel,

            string? dataBackup,

            string? description,

            string? deviceCompliance,

            string? fileEncryptionLevel,

            string? fileSharingSaveAs,

            string friendlyName,

            string groupStatus,

            string lastModifiedTime,

            string? location,

            string? managedBrowser,

            string name,

            int numOfApps,

            string? offlineWipeTimeout,

            string? pin,

            int? pinNumRetry,

            ImmutableDictionary<string, string>? tags,

            string? touchId,

            string type)
        {
            AccessRecheckOfflineTimeout = accessRecheckOfflineTimeout;
            AccessRecheckOnlineTimeout = accessRecheckOnlineTimeout;
            AppSharingFromLevel = appSharingFromLevel;
            AppSharingToLevel = appSharingToLevel;
            Authentication = authentication;
            ClipboardSharingLevel = clipboardSharingLevel;
            DataBackup = dataBackup;
            Description = description;
            DeviceCompliance = deviceCompliance;
            FileEncryptionLevel = fileEncryptionLevel;
            FileSharingSaveAs = fileSharingSaveAs;
            FriendlyName = friendlyName;
            GroupStatus = groupStatus;
            LastModifiedTime = lastModifiedTime;
            Location = location;
            ManagedBrowser = managedBrowser;
            Name = name;
            NumOfApps = numOfApps;
            OfflineWipeTimeout = offlineWipeTimeout;
            Pin = pin;
            PinNumRetry = pinNumRetry;
            Tags = tags;
            TouchId = touchId;
            Type = type;
        }
    }
}
