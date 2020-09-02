// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Solutions.V20200821Preview.Inputs
{

    /// <summary>
    /// Managed application locking policy.
    /// </summary>
    public sealed class ApplicationPackageLockingPolicyDefinitionArgs : Pulumi.ResourceArgs
    {
        [Input("allowedActions")]
        private InputList<string>? _allowedActions;

        /// <summary>
        /// The deny assignment excluded actions.
        /// </summary>
        public InputList<string> AllowedActions
        {
            get => _allowedActions ?? (_allowedActions = new InputList<string>());
            set => _allowedActions = value;
        }

        public ApplicationPackageLockingPolicyDefinitionArgs()
        {
        }
    }
}
