package isg_test

import (
	"context"
	"github.com/derbylock/isg-go/isg-go-lib/pkg/isg"
	"testing"
	"time"
)

func TestContextKeeper(t *testing.T) {
	testTime := time.Now()
	keeper := isg.NewContextValueContextKeeper()

	testCases := []struct {
		name     string
		ctx      context.Context
		inbound  *isg.InboundContext
		expected *isg.InboundContext
	}{
		{
			name:     "Keep and Extract InboundContext",
			ctx:      context.Background(),
			inbound:  isg.NewInboundContext("service1", "component1", "type1", "id1", testTime),
			expected: isg.NewInboundContext("service1", "component1", "type1", "id1", testTime),
		},
		{
			name:     "Extract nil when no InboundContext",
			ctx:      context.Background(),
			inbound:  nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := keeper.KeepInboundContext(tc.ctx, tc.inbound)
			actual := keeper.ExtractInboundContext(ctx)

			if actual == nil && tc.expected != nil {
				t.Errorf("Expected %v, got nil", tc.expected)
			} else if actual != nil && tc.expected == nil {
				t.Errorf("Expected nil, got %v", actual)
			} else if actual != nil && tc.expected != nil && *actual != *tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
