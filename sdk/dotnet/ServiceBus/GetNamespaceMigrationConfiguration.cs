// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceBus
{
    public static class GetNamespaceMigrationConfiguration
    {
        public static Task<GetNamespaceMigrationConfigurationResult> InvokeAsync(GetNamespaceMigrationConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNamespaceMigrationConfigurationResult>("azurerm:servicebus:getNamespaceMigrationConfiguration", args ?? new GetNamespaceMigrationConfigurationArgs(), options.WithVersion());
    }


    public sealed class GetNamespaceMigrationConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The configuration name. Should always be "$default".
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNamespaceMigrationConfigurationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNamespaceMigrationConfigurationResult
    {
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties required to the Create Migration Configuration
        /// </summary>
        public readonly Outputs.MigrationConfigPropertiesResponsePropertiesResult Properties;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetNamespaceMigrationConfigurationResult(
            string name,

            Outputs.MigrationConfigPropertiesResponsePropertiesResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}