package receipt_filter

import (
	"context"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	svcapitypes "github.com/aws-controllers-k8s/ses-controller/apis/v1alpha1"
	svcsdk "github.com/aws/aws-sdk-go/service/ses"
)

func (rm *resourceManager) customFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if r.ko.Spec.Name == nil {
		return nil, ackerr.NotFound
	}

	var resp *svcsdk.ListReceiptFiltersOutput
	resp, err = rm.sdkapi.ListReceiptFiltersWithContext(ctx, &svcsdk.ListReceiptFiltersInput{})

	rm.metrics.RecordAPICall("READ_MANY", "ListReceiptFilters", err)
	if err != nil {
		if reqErr, ok := ackerr.AWSRequestFailure(err); ok && reqErr.StatusCode() == 404 {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	rm.setStatusDefaults(ko)
	setReceiptFilter(ko, resp.Filters)

	return &resource{ko}, nil
}

func setReceiptFilter(ko *svcapitypes.ReceiptFilter, receiptFilters []*svcsdk.ReceiptFilter) {
	for _, filter := range receiptFilters {
		if *filter.Name == *ko.Spec.Name {
			ko.Spec.Filter = setResourceReceiptFilter(filter)
		}
	}
}
