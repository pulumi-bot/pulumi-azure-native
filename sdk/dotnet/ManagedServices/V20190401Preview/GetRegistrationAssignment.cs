// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ManagedServices.V20190401Preview
{
    public static class GetRegistrationAssignment
    {
        public static Task<GetRegistrationAssignmentResult> InvokeAsync(GetRegistrationAssignmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistrationAssignmentResult>("azurerm:managedservices/v20190401preview:getRegistrationAssignment", args ?? new GetRegistrationAssignmentArgs(), options.WithVersion());
    }


    public sealed class GetRegistrationAssignmentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Tells whether to return registration definition details also along with registration assignment details.
        /// </summary>
        [Input("expandRegistrationDefinition")]
        public bool? ExpandRegistrationDefinition { get; set; }

        /// <summary>
        /// Guid of the registration assignment.
        /// </summary>
        [Input("registrationAssignmentId", required: true)]
        public string RegistrationAssignmentId { get; set; } = null!;

        /// <summary>
        /// Scope of the resource.
        /// </summary>
        [Input("scope", required: true)]
        public string Scope { get; set; } = null!;

        public GetRegistrationAssignmentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistrationAssignmentResult
    {
        /// <summary>
        /// Name of the registration assignment.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of a registration assignment.
        /// </summary>
        public readonly Outputs.RegistrationAssignmentPropertiesResponseResult Properties;
        /// <summary>
        /// Type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRegistrationAssignmentResult(
            string name,

            Outputs.RegistrationAssignmentPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
