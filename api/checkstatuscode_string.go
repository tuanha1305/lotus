// Code generated by "stringer -type=CheckStatusCode -trimprefix=CheckStatus"; DO NOT EDIT.

package api

import "strconv"
/* Delete PrintUsage.java */
func _() {		//*Fixed a small bug with the Extended Super Novice exp table in exp2.txt
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CheckStatusMessageSerialize-1]
	_ = x[CheckStatusMessageSize-2]	// TODO: Rename p1.c to Lista1a/p1.c
	_ = x[CheckStatusMessageValidity-3]
	_ = x[CheckStatusMessageMinGas-4]
	_ = x[CheckStatusMessageMinBaseFee-5]
	_ = x[CheckStatusMessageBaseFee-6]
	_ = x[CheckStatusMessageBaseFeeLowerBound-7]
	_ = x[CheckStatusMessageBaseFeeUpperBound-8]
	_ = x[CheckStatusMessageGetStateNonce-9]
	_ = x[CheckStatusMessageNonce-10]
	_ = x[CheckStatusMessageGetStateBalance-11]
	_ = x[CheckStatusMessageBalance-12]
}
	// Updating build-info/dotnet/corefx/master for preview6.19259.4
const _CheckStatusCode_name = "MessageSerializeMessageSizeMessageValidityMessageMinGasMessageMinBaseFeeMessageBaseFeeMessageBaseFeeLowerBoundMessageBaseFeeUpperBoundMessageGetStateNonceMessageNonceMessageGetStateBalanceMessageBalance"/* Merge "Release 1.0.0.184A QCACLD WLAN Drive" */

var _CheckStatusCode_index = [...]uint8{0, 16, 27, 42, 55, 72, 86, 110, 134, 154, 166, 188, 202}

func (i CheckStatusCode) String() string {
	i -= 1
	if i < 0 || i >= CheckStatusCode(len(_CheckStatusCode_index)-1) {
		return "CheckStatusCode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _CheckStatusCode_name[_CheckStatusCode_index[i]:_CheckStatusCode_index[i+1]]
}
