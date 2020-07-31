// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DomainRegistration.V20180201
{
    public static class GetDomainOwnershipIdentifier
    {
        public static Task<GetDomainOwnershipIdentifierResult> InvokeAsync(GetDomainOwnershipIdentifierArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainOwnershipIdentifierResult>("azurerm:domainregistration/v20180201:getDomainOwnershipIdentifier", args ?? new GetDomainOwnershipIdentifierArgs(), options.WithVersion());
    }


    public sealed class GetDomainOwnershipIdentifierArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of domain.
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        /// <summary>
        /// Name of identifier.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDomainOwnershipIdentifierArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainOwnershipIdentifierResult
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// DomainOwnershipIdentifier resource specific properties
        /// </summary>
        public readonly Outputs.DomainOwnershipIdentifierResponsePropertiesResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDomainOwnershipIdentifierResult(
            string? kind,

            string name,

            Outputs.DomainOwnershipIdentifierResponsePropertiesResult properties,

            string type)
        {
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}