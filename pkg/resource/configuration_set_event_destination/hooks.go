package configuration_set_event_destination

import (
	svcapitypes "github.com/aws-controllers-k8s/ses-controller/apis/v1alpha1"
	svcsdk "github.com/aws/aws-sdk-go/service/ses"
)

func setEventDestinations(ko *svcapitypes.ConfigurationSetEventDestination, resp *svcsdk.DescribeConfigurationSetOutput) {
	if len(resp.EventDestinations) != 1 {
		return
	}
	ko.Spec.EventDestination = setResourceEventDestination(resp.EventDestinations[0])
}
