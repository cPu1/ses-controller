package configuration_set_event_destination

import (
	svcsdk "github.com/aws/aws-sdk-go/service/ses"

	svcapitypes "github.com/aws-controllers-k8s/ses-controller/apis/v1alpha1"
)

func getEventDestination(ko *svcapitypes.ConfigurationSetEventDestination, resp *svcsdk.DescribeConfigurationSetOutput) *svcapitypes.EventDestination {
	for _, eventDestination := range resp.EventDestinations {
		if *eventDestination.Name == *ko.Spec.EventDestination.Name {
			return setResourceEventDestination(resp.EventDestinations[0])
		}
	}
	return nil
}
