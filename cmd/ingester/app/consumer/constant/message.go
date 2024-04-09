package constant

import "github.com/Shopify/sarama"

// Copyright (c) 2018 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Message contains the parts of a sarama ConsumerMessage that we care about.
type Message interface {
	Key() []byte
	Value() []byte
	Topic() string
	Partition() int32
	Offset() int64
	Headers() []*sarama.RecordHeader
}

type SaramaMessageWrapper struct {
	*sarama.ConsumerMessage
}

func (m SaramaMessageWrapper) Key() []byte {
	return m.ConsumerMessage.Key
}

func (m SaramaMessageWrapper) Value() []byte {
	return m.ConsumerMessage.Value
}

func (m SaramaMessageWrapper) Topic() string {
	return m.ConsumerMessage.Topic
}

func (m SaramaMessageWrapper) Partition() int32 {
	return m.ConsumerMessage.Partition
}

func (m SaramaMessageWrapper) Offset() int64 {
	return m.ConsumerMessage.Offset
}

func (m SaramaMessageWrapper) Headers() []*sarama.RecordHeader {
	return m.ConsumerMessage.Headers
}
