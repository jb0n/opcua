// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package uacp

import (
	"testing"

	"github.com/jb0n/opcua/ua"
)

func TestHello(t *testing.T) {
	cases := []CodecTestCase{
		{
			Struct: &Hello{
				Version:        0,
				ReceiveBufSize: 65280,
				SendBufSize:    65535,
				MaxMessageSize: 4000,
				MaxChunkCount:  1234,
				EndpointURL:    "opc.tcp://wow.its.easy:11111/UA/Server",
			},
			Bytes: []byte{ // Hello message
				// Version: 0
				0x00, 0x00, 0x00, 0x00,
				// ReceiveBufSize: 65280
				0x00, 0xff, 0x00, 0x00,
				// SendBufSize: 65535
				0xff, 0xff, 0x00, 0x00,
				// MaxMessageSize: 4000
				0xa0, 0x0f, 0x00, 0x00,
				// MaxChunkCount: 1234
				0xd2, 0x04, 0x00, 0x00,
				// EndPointURL
				0x26, 0x00, 0x00, 0x00, 0x6f, 0x70, 0x63, 0x2e,
				0x74, 0x63, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x6f,
				0x77, 0x2e, 0x69, 0x74, 0x73, 0x2e, 0x65, 0x61,
				0x73, 0x79, 0x3a, 0x31, 0x31, 0x31, 0x31, 0x31,
				0x2f, 0x55, 0x41, 0x2f, 0x53, 0x65, 0x72, 0x76,
				0x65, 0x72,
			},
		},
	}
	RunCodecTest(t, cases)
}

func TestAcknowledge(t *testing.T) {
	cases := []CodecTestCase{
		{
			Struct: &Acknowledge{
				Version:        0,
				ReceiveBufSize: 65280,
				SendBufSize:    65535,
				MaxMessageSize: 4000,
				MaxChunkCount:  1234,
			},
			Bytes: []byte{
				// Version: 0
				0x00, 0x00, 0x00, 0x00,
				// ReceiveBufSize: 65280
				0x00, 0xff, 0x00, 0x00,
				// SendBufSize: 65535
				0xff, 0xff, 0x00, 0x00,
				// MaxMessageSize: 4000
				0xa0, 0x0f, 0x00, 0x00,
				// MaxChunkCount: 1234
				0xd2, 0x04, 0x00, 0x00,
			},
		},
	}
	RunCodecTest(t, cases)
}

func TestReverseHello(t *testing.T) {
	cases := []CodecTestCase{
		{
			Name: "Normal",
			Struct: &ReverseHello{
				ServerURI:   "opc.tcp://wow.its.easy:11111/UA/Server",
				EndpointURL: "opc.tcp://wow.its.easy:11111/UA/Server",
			},
			Bytes: []byte{
				// ServerURI
				0x26, 0x00, 0x00, 0x00, 0x6f, 0x70, 0x63, 0x2e,
				0x74, 0x63, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x6f,
				0x77, 0x2e, 0x69, 0x74, 0x73, 0x2e, 0x65, 0x61,
				0x73, 0x79, 0x3a, 0x31, 0x31, 0x31, 0x31, 0x31,
				0x2f, 0x55, 0x41, 0x2f, 0x53, 0x65, 0x72, 0x76,
				0x65, 0x72,
				// EndPointURL
				0x26, 0x00, 0x00, 0x00, 0x6f, 0x70, 0x63, 0x2e,
				0x74, 0x63, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x6f,
				0x77, 0x2e, 0x69, 0x74, 0x73, 0x2e, 0x65, 0x61,
				0x73, 0x79, 0x3a, 0x31, 0x31, 0x31, 0x31, 0x31,
				0x2f, 0x55, 0x41, 0x2f, 0x53, 0x65, 0x72, 0x76,
				0x65, 0x72,
			},
		},
	}
	RunCodecTest(t, cases)
}

func TestError(t *testing.T) {
	cases := []CodecTestCase{
		{
			Struct: &Error{
				ErrorCode: uint32(ua.StatusBadSecureChannelClosed),
				Reason:    "foobar",
			},
			Bytes: []byte{
				// Error: BadSecureChannelClosed
				0x00, 0x00, 0x86, 0x80,
				// Reason: dummy
				0x06, 0x00, 0x00, 0x00, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72,
			},
		},
	}
	RunCodecTest(t, cases)
}
