// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Insights.Latest
{
    [Obsolete(@"The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:insights:getComponentCurrentBillingFeature'.")]
    public static class GetComponentCurrentBillingFeature
    {
        /// <summary>
        /// An Application Insights component billing features
        /// Latest API Version: 2015-05-01.
        /// </summary>
        public static Task<GetComponentCurrentBillingFeatureResult> InvokeAsync(GetComponentCurrentBillingFeatureArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetComponentCurrentBillingFeatureResult>("azure-native:insights/latest:getComponentCurrentBillingFeature", args ?? new GetComponentCurrentBillingFeatureArgs(), options.WithVersion());
    }


    public sealed class GetComponentCurrentBillingFeatureArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Application Insights component resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetComponentCurrentBillingFeatureArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetComponentCurrentBillingFeatureResult
    {
        /// <summary>
        /// Current enabled pricing plan. When the component is in the Enterprise plan, this will list both 'Basic' and 'Application Insights Enterprise'.
        /// </summary>
        public readonly ImmutableArray<string> CurrentBillingFeatures;
        /// <summary>
        /// An Application Insights component daily data volume cap
        /// </summary>
        public readonly Outputs.ApplicationInsightsComponentDataVolumeCapResponse? DataVolumeCap;

        [OutputConstructor]
        private GetComponentCurrentBillingFeatureResult(
            ImmutableArray<string> currentBillingFeatures,

            Outputs.ApplicationInsightsComponentDataVolumeCapResponse? dataVolumeCap)
        {
            CurrentBillingFeatures = currentBillingFeatures;
            DataVolumeCap = dataVolumeCap;
        }
    }
}
