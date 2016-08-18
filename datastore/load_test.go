// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datastore

import (
	"reflect"
	"testing"

	pb "google.golang.org/genproto/googleapis/datastore/v1"
)

type A0 struct {
	I int64
}

type AT struct {
	I int64 `datastore:"II"`
}

type ANT struct {
	A AT `datastore:"AA"`
}

type A0N struct {
	A []A0
}

type A1 struct {
	S  string
	SS string
}

type AA struct {
	A0
	X string
}

type AN struct {
	A A0
	B []byte
	I int
}

type AN1 struct {
	A A0
	X string
}

type AN2X struct {
	AN AN
	A  A1
	S  string
}

type NS0 struct {
	A         A0
	ARef      *A0
	ASlice    []A0
	ARefSlice []*A0
}

type NS1 struct {
	AAnonym   AA
	AUnexport A1
	ANested   AN

	AAnonymRef *AA
	ANestedRef *AN

	AAnonymSlice []AA
	ANestedSlice []AN

	AAnonymRefSlice []*AA
	ANestedRefSlice []*AN
}

var (
	a00 = A0{3}
	a01 = A0{4}

	aa0 = AA{a00, "X"}
	aa1 = AA{a01, "Y"}

	a10 = A1{"S", "s"}

	an0 = AN{a00, []byte("xx"), 1}
	an1 = AN{a01, []byte("yy"), 2}
)

func TestLoadEntityNested(t *testing.T) {
	testCases := []struct {
		desc string
		src  interface{}
	}{
		{
			"nested basic",
			&AN{
				A: a00,
				B: []byte("a00"),
				I: 10,
			},
		},
		{
			"nested with struct tags",
			&ANT{
				A: AT{1},
			},
		},
		{
			"nested 2x",
			&AN2X{
				AN: an0,
				A:  a10,
				S:  "SS",
			},
		},
		{
			"nested anonymous",
			&AA{
				a00,
				"SomeX",
			},
		},
		{
			"nested with ptrs",
			&NS0{
				A:         a00,
				ARef:      &a00,
				ASlice:    []A0{a00, a01},
				ARefSlice: []*A0{&a00, &a01},
			},
		},
		{
			"nested simple",
			&A0N{
				A: []A0{a00, a01},
			},
		},
		{
			"nested complex",
			&NS1{
				AAnonym:   aa0,
				AUnexport: a10,
				ANested:   an0,

				AAnonymRef: &aa0,
				ANestedRef: &an0,

				AAnonymSlice: []AA{aa0, aa1},
				ANestedSlice: []AN{an0, an1},

				AAnonymRefSlice: []*AA{&aa0, &aa1},
				ANestedRefSlice: []*AN{&an0, &an1},
			},
		},
	}

	for _, tc := range testCases {
		entity, err := saveEntity(testKey0, tc.src)
		if err != nil {
			t.Errorf("saveEntity: %s: %v", tc.desc, err)
			continue
		}

		dst := reflect.New(reflect.TypeOf(tc.src).Elem()).Interface()
		err = loadEntity(dst, entity)
		if err != nil {
			t.Errorf("loadEntity: %s: %v", tc.desc, err)
			continue
		}

		if !reflect.DeepEqual(tc.src, dst) {
			t.Errorf("%s: compare:\ngot: %#v\nwant: %#v", tc.desc, dst, tc.src)
		}
	}
}

func TestLoadEntityNestedLegacy(t *testing.T) {
	testCases := []struct {
		desc string
		src  *pb.Entity
		want interface{}
	}{
		{
			"nested",
			&pb.Entity{
				Key: keyToProto(testKey0),
				Properties: map[string]*pb.Value{
					"X":   {ValueType: &pb.Value_StringValue{"two"}},
					"A.I": {ValueType: &pb.Value_IntegerValue{2}},
				},
			},
			&AN1{
				A: A0{I: 2},
				X: "two",
			},
		},
		{
			"nested with tag",
			&pb.Entity{
				Key: keyToProto(testKey0),
				Properties: map[string]*pb.Value{
					"AA.II": {ValueType: &pb.Value_IntegerValue{2}},
				},
			},
			&ANT{
				A: AT{I: 2},
			},
		},
		{
			"nested with anonymous struct field",
			&pb.Entity{
				Key: keyToProto(testKey0),
				Properties: map[string]*pb.Value{
					"X": {ValueType: &pb.Value_StringValue{"two"}},
					"I": {ValueType: &pb.Value_IntegerValue{2}},
				},
			},
			&AA{
				A0: A0{I: 2},
				X:  "two",
			},
		},
		// TODO: test anonymous struct fields' fields promoted to top level
	}

	for _, tc := range testCases {
		dst := reflect.New(reflect.TypeOf(tc.want).Elem()).Interface()
		err := loadEntity(dst, tc.src)
		if err != nil {
			t.Errorf("loadEntity: %s: %v", tc.desc, err)
			continue
		}

		if !reflect.DeepEqual(tc.want, dst) {
			t.Errorf("%s: compare:\ngot: %#v\nwant: %#v", tc.desc, dst, tc.src)
		}
	}
}
