// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DesktopVirtualization.V20201102Preview.Outputs
{

    [OutputType]
    public sealed class MsixPackageApplicationsResponse
    {
        /// <summary>
        /// Package Application Id, found in appxmanifest.xml.
        /// </summary>
        public readonly string? AppId;
        /// <summary>
        /// Used to activate Package Application. Consists of Package Name and ApplicationID. Found in appxmanifest.xml.
        /// </summary>
        public readonly string? AppUserModelID;
        /// <summary>
        /// Description of Package Application.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// User friendly name.
        /// </summary>
        public readonly string? FriendlyName;
        /// <summary>
        /// User friendly name.
        /// </summary>
        public readonly string? IconImageName;
        /// <summary>
        /// the icon a 64 bit string as a byte array.
        /// </summary>
        public readonly string? RawIcon;
        /// <summary>
        /// the icon a 64 bit string as a byte array.
        /// </summary>
        public readonly string? RawPng;

        [OutputConstructor]
        private MsixPackageApplicationsResponse(
            string? appId,

            string? appUserModelID,

            string? description,

            string? friendlyName,

            string? iconImageName,

            string? rawIcon,

            string? rawPng)
        {
            AppId = appId;
            AppUserModelID = appUserModelID;
            Description = description;
            FriendlyName = friendlyName;
            IconImageName = iconImageName;
            RawIcon = rawIcon;
            RawPng = rawPng;
        }
    }
}
