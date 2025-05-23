// Copyright The OpenTelemetry Authors
// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2018 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package zipkin

import (
	"context"
	"testing"

	"github.com/jaegertracing/jaeger-idl/thrift-gen/zipkincore"
	"github.com/stretchr/testify/require"
)

func TestDeserializeWithBadListStart(t *testing.T) {
	ctx := context.Background()
	spanBytes, err := SerializeThrift(ctx, []*zipkincore.Span{{}})
	require.NoError(t, err)
	_, err = DeserializeThrift(ctx, append([]byte{0, 255, 255}, spanBytes...))
	require.Error(t, err)
}

func TestDeserializeWithCorruptedList(t *testing.T) {
	ctx := context.Background()
	spanBytes, err := SerializeThrift(ctx, []*zipkincore.Span{{}})
	require.NoError(t, err)
	spanBytes[2] = 255
	_, err = DeserializeThrift(ctx, spanBytes)
	require.Error(t, err)
}

func TestDeserialize(t *testing.T) {
	ctx := context.Background()
	spanBytes, err := SerializeThrift(ctx, []*zipkincore.Span{{}})
	require.NoError(t, err)
	_, err = DeserializeThrift(ctx, spanBytes)
	require.NoError(t, err)
}
