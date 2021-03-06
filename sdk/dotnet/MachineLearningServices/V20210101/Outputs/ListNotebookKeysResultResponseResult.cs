// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20210101.Outputs
{

    [OutputType]
    public sealed class ListNotebookKeysResultResponseResult
    {
        public readonly string PrimaryAccessKey;
        public readonly string SecondaryAccessKey;

        [OutputConstructor]
        private ListNotebookKeysResultResponseResult(
            string primaryAccessKey,

            string secondaryAccessKey)
        {
            PrimaryAccessKey = primaryAccessKey;
            SecondaryAccessKey = secondaryAccessKey;
        }
    }
}
