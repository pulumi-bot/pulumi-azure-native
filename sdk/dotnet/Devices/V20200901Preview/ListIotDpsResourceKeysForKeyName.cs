// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20200901Preview
{
    public static class ListIotDpsResourceKeysForKeyName
    {
        public static Task<ListIotDpsResourceKeysForKeyNameResult> InvokeAsync(ListIotDpsResourceKeysForKeyNameArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListIotDpsResourceKeysForKeyNameResult>("azurerm:devices/v20200901preview:listIotDpsResourceKeysForKeyName", args ?? new ListIotDpsResourceKeysForKeyNameArgs(), options.WithVersion());
    }


    public sealed class ListIotDpsResourceKeysForKeyNameArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Logical key name to get key-values for.
        /// </summary>
        [Input("keyName", required: true)]
        public string KeyName { get; set; } = null!;

        /// <summary>
        /// Name of the provisioning service.
        /// </summary>
        [Input("provisioningServiceName", required: true)]
        public string ProvisioningServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the provisioning service.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListIotDpsResourceKeysForKeyNameArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListIotDpsResourceKeysForKeyNameResult
    {
        /// <summary>
        /// Name of the key.
        /// </summary>
        public readonly string KeyName;
        /// <summary>
        /// Primary SAS key value.
        /// </summary>
        public readonly string? PrimaryKey;
        /// <summary>
        /// Rights that this key has.
        /// </summary>
        public readonly string Rights;
        /// <summary>
        /// Secondary SAS key value.
        /// </summary>
        public readonly string? SecondaryKey;

        [OutputConstructor]
        private ListIotDpsResourceKeysForKeyNameResult(
            string keyName,

            string? primaryKey,

            string rights,

            string? secondaryKey)
        {
            KeyName = keyName;
            PrimaryKey = primaryKey;
            Rights = rights;
            SecondaryKey = secondaryKey;
        }
    }
}
