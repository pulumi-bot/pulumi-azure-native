// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerRegistry.V20180901.Inputs
{

    /// <summary>
    /// Describes the properties of a secret object value.
    /// </summary>
    public sealed class SecretObjectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of the secret object which determines how the value of the secret object has to be
        /// interpreted.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNextGen.ContainerRegistry.V20180901.SecretObjectType>? Type { get; set; }

        /// <summary>
        /// The value of the secret. The format of this value will be determined
        /// based on the type of the secret object. If the type is Opaque, the value will be
        /// used as is without any modification.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public SecretObjectArgs()
        {
        }
    }
}
