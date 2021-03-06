// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20200601.Outputs
{

    [OutputType]
    public sealed class GoogleResponse
    {
        public readonly bool? Enabled;
        /// <summary>
        /// Resource Id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        public readonly Outputs.LoginScopesResponse? Login;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        public readonly Outputs.ClientRegistrationResponse? Registration;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        public readonly Outputs.AllowedAudiencesValidationResponse? Validation;

        [OutputConstructor]
        private GoogleResponse(
            bool? enabled,

            string id,

            string? kind,

            Outputs.LoginScopesResponse? login,

            string name,

            Outputs.ClientRegistrationResponse? registration,

            string type,

            Outputs.AllowedAudiencesValidationResponse? validation)
        {
            Enabled = enabled;
            Id = id;
            Kind = kind;
            Login = login;
            Name = name;
            Registration = registration;
            Type = type;
            Validation = validation;
        }
    }
}
