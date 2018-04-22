/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mockkms

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
)

type MockKMS struct {
}

func (_ *MockKMS) CreateKey(*kms.CreateKeyInput) (*kms.CreateKeyOutput, error) {
	return &kms.CreateKeyOutput{
		KeyMetadata: &kms.KeyMetadata{
			Arn:   aws.String("arn:aws:kms:us-east-1:fake-key"),
			KeyId: aws.String("arn:aws:kms:us-east-1:fake-key"),
		},
	}, nil
}
func (_ *MockKMS) CreateAlias(i *kms.CreateAliasInput) (*kms.CreateAliasOutput, error) {
	return &kms.CreateAliasOutput{}, nil
}

func (_ *MockKMS) DescribeKey(*kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error) {
	return &kms.DescribeKeyOutput{
		KeyMetadata: &kms.KeyMetadata{
			Arn:   aws.String("arn:aws:kms:us-east-1:fake-key"),
			KeyId: aws.String("arn:aws:kms:us-east-1:fake-key"),
		},
	}, nil
}

func (_ *MockKMS) DisableKey(*kms.DisableKeyInput) (*kms.DisableKeyOutput, error) {
	return &kms.DisableKeyOutput{}, nil
}

var _ kmsiface.KMSAPI = &MockKMS{}
