// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DesktopVirtualization.V20201019Preview
{
    public static class GetApplication
    {
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("azure-nextgen:desktopvirtualization/v20201019preview:getApplication", args ?? new GetApplicationArgs(), options.WithVersion());
    }


    public sealed class GetApplicationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the application group
        /// </summary>
        [Input("applicationGroupName", required: true)]
        public string ApplicationGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the application within the specified application group
        /// </summary>
        [Input("applicationName", required: true)]
        public string ApplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApplicationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// Resource Type of Application.
        /// </summary>
        public readonly string? ApplicationType;
        /// <summary>
        /// Command Line Arguments for Application.
        /// </summary>
        public readonly string? CommandLineArguments;
        /// <summary>
        /// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
        /// </summary>
        public readonly string CommandLineSetting;
        /// <summary>
        /// Description of Application.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Specifies a path for the executable file for the application.
        /// </summary>
        public readonly string? FilePath;
        /// <summary>
        /// Friendly name of Application.
        /// </summary>
        public readonly string? FriendlyName;
        /// <summary>
        /// the icon a 64 bit string as a byte array.
        /// </summary>
        public readonly string IconContent;
        /// <summary>
        /// Hash of the icon.
        /// </summary>
        public readonly string IconHash;
        /// <summary>
        /// Index of the icon.
        /// </summary>
        public readonly int? IconIndex;
        /// <summary>
        /// Path to icon.
        /// </summary>
        public readonly string? IconPath;
        /// <summary>
        /// Specifies the package application Id for MSIX applications
        /// </summary>
        public readonly string? MsixPackageApplicationId;
        /// <summary>
        /// Specifies the package family name for MSIX applications
        /// </summary>
        public readonly string? MsixPackageFamilyName;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies whether to show the RemoteApp program in the RD Web Access server.
        /// </summary>
        public readonly bool? ShowInPortal;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApplicationResult(
            string? applicationType,

            string? commandLineArguments,

            string commandLineSetting,

            string? description,

            string? filePath,

            string? friendlyName,

            string iconContent,

            string iconHash,

            int? iconIndex,

            string? iconPath,

            string? msixPackageApplicationId,

            string? msixPackageFamilyName,

            string name,

            bool? showInPortal,

            string type)
        {
            ApplicationType = applicationType;
            CommandLineArguments = commandLineArguments;
            CommandLineSetting = commandLineSetting;
            Description = description;
            FilePath = filePath;
            FriendlyName = friendlyName;
            IconContent = iconContent;
            IconHash = iconHash;
            IconIndex = iconIndex;
            IconPath = iconPath;
            MsixPackageApplicationId = msixPackageApplicationId;
            MsixPackageFamilyName = msixPackageFamilyName;
            Name = name;
            ShowInPortal = showInPortal;
            Type = type;
        }
    }
}
