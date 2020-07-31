// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20180201
{
    public static class GetWebAppDomainOwnershipIdentifier
    {
        public static Task<GetWebAppDomainOwnershipIdentifierResult> InvokeAsync(GetWebAppDomainOwnershipIdentifierArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWebAppDomainOwnershipIdentifierResult>("azurerm:web/v20180201:getWebAppDomainOwnershipIdentifier", args ?? new GetWebAppDomainOwnershipIdentifierArgs(), options.WithVersion());
    }


    public sealed class GetWebAppDomainOwnershipIdentifierArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of domain ownership identifier.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetWebAppDomainOwnershipIdentifierArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWebAppDomainOwnershipIdentifierResult
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
        /// Identifier resource specific properties
        /// </summary>
        public readonly Outputs.IdentifierResponsePropertiesResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetWebAppDomainOwnershipIdentifierResult(
            string? kind,

            string name,

            Outputs.IdentifierResponsePropertiesResult properties,

            string type)
        {
            Kind = kind;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}