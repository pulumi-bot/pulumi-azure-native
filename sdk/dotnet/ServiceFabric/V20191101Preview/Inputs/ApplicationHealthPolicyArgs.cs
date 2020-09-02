// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20191101Preview.Inputs
{

    /// <summary>
    /// Defines a health policy used to evaluate the health of an application or one of its children entities.
    /// </summary>
    public sealed class ApplicationHealthPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The health policy used by default to evaluate the health of a service type.
        /// </summary>
        [Input("defaultServiceTypeHealthPolicy")]
        public Input<Inputs.ServiceTypeHealthPolicyArgs>? DefaultServiceTypeHealthPolicy { get; set; }

        [Input("serviceTypeHealthPolicies")]
        private InputMap<Inputs.ServiceTypeHealthPolicyArgs>? _serviceTypeHealthPolicies;

        /// <summary>
        /// The map with service type health policy per service type name. The map is empty by default.
        /// </summary>
        public InputMap<Inputs.ServiceTypeHealthPolicyArgs> ServiceTypeHealthPolicies
        {
            get => _serviceTypeHealthPolicies ?? (_serviceTypeHealthPolicies = new InputMap<Inputs.ServiceTypeHealthPolicyArgs>());
            set => _serviceTypeHealthPolicies = value;
        }

        public ApplicationHealthPolicyArgs()
        {
        }
    }
}
