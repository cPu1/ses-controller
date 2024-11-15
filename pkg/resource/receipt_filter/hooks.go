package receipt_filter

import (
	"context"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	svcsdk "github.com/aws/aws-sdk-go/service/ses"

	"github.com/aws-controllers-k8s/ses-controller/pkg/util"
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
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	var receiptFilter *svcsdk.ReceiptFilter
	for _, filter := range resp.Filters {
		if *filter.Name == *r.ko.Spec.Name {
			receiptFilter = filter
			break
		}
	}
	if receiptFilter == nil {
		return nil, ackerr.NotFound
	}
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)
	ko.Spec.Filter = setResourceReceiptFilter(receiptFilter)
	return &resource{ko}, nil
}

func (rm *resourceManager) customUpdate(
	ctx context.Context,
	desired *resource,
	_ *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	return util.ValidateImmutableResource(ctx, rm.getImmutableFieldChanges(delta), desired)
}
