// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190601Preview.Inputs
{

    /// <summary>
    /// Defines a delta health policy used to evaluate the health of an application or one of its child entities when upgrading the cluster.
    /// </summary>
    public sealed class ApplicationDeltaHealthPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The delta health policy used by default to evaluate the health of a service type when upgrading the cluster.
        /// </summary>
        [Input("defaultServiceTypeDeltaHealthPolicy")]
        public Input<Inputs.ServiceTypeDeltaHealthPolicyArgs>? DefaultServiceTypeDeltaHealthPolicy { get; set; }

        [Input("serviceTypeDeltaHealthPolicies")]
        private InputMap<Inputs.ServiceTypeDeltaHealthPolicyArgs>? _serviceTypeDeltaHealthPolicies;

        /// <summary>
        /// The map with service type delta health policy per service type name. The map is empty by default.
        /// </summary>
        public InputMap<Inputs.ServiceTypeDeltaHealthPolicyArgs> ServiceTypeDeltaHealthPolicies
        {
            get => _serviceTypeDeltaHealthPolicies ?? (_serviceTypeDeltaHealthPolicies = new InputMap<Inputs.ServiceTypeDeltaHealthPolicyArgs>());
            set => _serviceTypeDeltaHealthPolicies = value;
        }

        public ApplicationDeltaHealthPolicyArgs()
        {
        }
    }
}
