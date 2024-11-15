package util

import (
	"context"
	"errors"
	"fmt"
	"strings"

	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	rtclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func ValidateImmutableResource[T interface{ RuntimeObject() rtclient.Object }](ctx context.Context, immutableFieldChanges []string, desired T) (t T, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	if len(immutableFieldChanges) > 0 {
		msg := fmt.Sprintf("Immutable Spec fields have been modified: %s", strings.Join(immutableFieldChanges, ","))
		return t, ackerr.NewTerminalError(errors.New(msg))
	}
	return desired, nil
}
