// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.V20190601
{
    public static class GetManagementPolicy
    {
        public static Task<GetManagementPolicyResult> InvokeAsync(GetManagementPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagementPolicyResult>("azure-nextgen:storage/v20190601:getManagementPolicy", args ?? new GetManagementPolicyArgs(), options.WithVersion());
    }


    public sealed class GetManagementPolicyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Storage Account Management Policy. It should always be 'default'
        /// </summary>
        [Input("managementPolicyName", required: true)]
        public string ManagementPolicyName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetManagementPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetManagementPolicyResult
    {
        /// <summary>
        /// Returns the date and time the ManagementPolicies was last modified.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Storage Account ManagementPolicy, in JSON format. See more details in: https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
        /// </summary>
        public readonly Outputs.ManagementPolicySchemaResponse Policy;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetManagementPolicyResult(
            string lastModifiedTime,

            string name,

            Outputs.ManagementPolicySchemaResponse policy,

            string type)
        {
            LastModifiedTime = lastModifiedTime;
            Name = name;
            Policy = policy;
            Type = type;
        }
    }
}
