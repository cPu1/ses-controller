	_ = resp
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == svcsdk.svcsdk.ErrCodeRuleDoesNotExistException {
			rm.metrics.RecordAPICall("READ_ONE", "DescribeReceiptRule", err)
			return nil, ackerr.NotFound
		}
	}
