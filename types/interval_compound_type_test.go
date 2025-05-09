package types

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	proto "github.com/substrait-io/substrait-protobuf/go/substraitpb"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestNewIntervalCompoundType(t *testing.T) {
	allPossibleTimePrecision := []TimePrecision{PrecisionSeconds, PrecisionDeciSeconds, PrecisionCentiSeconds, PrecisionMilliSeconds,
		PrecisionEMinus4Seconds, PrecisionEMinus5Seconds, PrecisionMicroSeconds, PrecisionEMinus7Seconds, PrecisionEMinus8Seconds, PrecisionNanoSeconds}
	allPossibleNullability := []Nullability{NullabilityUnspecified, NullabilityNullable, NullabilityRequired}

	for _, precision := range allPossibleTimePrecision {
		for _, nullability := range allPossibleNullability {
			expectedIntervalCompoundType := IntervalCompoundType{precision: precision, nullability: nullability}
			expectedFormatString := fmt.Sprintf("%s<%d>", strNullable(expectedIntervalCompoundType), precision.ToProtoVal())

			parameters := expectedIntervalCompoundType.GetParameters()
			assert.Equal(t, parameters, []interface{}{precision})
			// verify IntervalCompoundType
			createdIntervalCompoundTypeIfc := NewIntervalCompoundType().WithPrecision(precision).WithTypeVariationRef(0).WithNullability(nullability)
			createdIntervalCompoundType := createdIntervalCompoundTypeIfc.(IntervalCompoundType)
			assert.True(t, createdIntervalCompoundType.Equals(expectedIntervalCompoundType))
			assert.Equal(t, expectedProtoValMap[precision], createdIntervalCompoundType.GetPrecisionProtoVal())
			assert.Equal(t, nullability, createdIntervalCompoundType.GetNullability())
			assert.Zero(t, createdIntervalCompoundType.GetTypeVariationReference())
			assert.Equal(t, fmt.Sprintf("interval_compound%s", expectedFormatString), createdIntervalCompoundType.String())
			assert.Equal(t, "icompound", createdIntervalCompoundType.ShortString())
			assertIntervalCompoundTypeProto(t, precision, nullability, createdIntervalCompoundType)
		}
	}
}

func assertIntervalCompoundTypeProto(t *testing.T, expectedPrecision TimePrecision, expectedNullability Nullability,
	toVerifyType IntervalCompoundType) {

	expectedTypeProto := &proto.Type{Kind: &proto.Type_IntervalCompound_{
		IntervalCompound: &proto.Type_IntervalCompound{
			Precision:   expectedPrecision.ToProtoVal(),
			Nullability: expectedNullability,
		},
	}}
	if diff := cmp.Diff(toVerifyType.ToProto(), expectedTypeProto, protocmp.Transform()); diff != "" {
		t.Errorf("IntervalCompoundType proto didn't match, diff:\n%v", diff)
	}

	expectedFuncArgProto := &proto.FunctionArgument{ArgType: &proto.FunctionArgument_Type{
		Type: expectedTypeProto,
	}}
	if diff := cmp.Diff(toVerifyType.ToProtoFuncArg(), expectedFuncArgProto, protocmp.Transform()); diff != "" {
		t.Errorf("IntervalCompoundType func arg proto didn't match, diff:\n%v", diff)
	}
}
