// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20170901Preview.Inputs
{

    /// <summary>
    /// Identity properties of the factory resource.
    /// </summary>
    public sealed class FactoryIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity type. Currently the only supported type is 'SystemAssigned'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FactoryIdentityArgs()
        {
        }
    }
}
