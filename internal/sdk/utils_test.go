// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package go_npm_proxy_manager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroStruct(t *testing.T) {

	type NestedObject struct {
		F4 *[]string
		F5 string
		F6 []int
	}

	type Object struct {
		F1 []string
		F2 *string
		F3 *[]int
		NestedObject
	}

	t.Run("Test zeroing nil slice pointers", func(t *testing.T) {

		obj := Object{}

		setZeroSlices(&obj)

		assert.Equal(t, []string{}, obj.F1)  // The nil []string should be empty
		assert.Nil(t, obj.F2)                // The string shouldn't be changed
		assert.Equal(t, []int{}, *obj.F3)    // The nil []int pointer should point to an empty slice
		assert.Equal(t, []string{}, *obj.F4) // The nil nested []string pointer should point to an empty slice
		assert.Equal(t, "", obj.F5)          // The empty nested string should be unchanged
		assert.Equal(t, []int{}, obj.F6)     // The nil nested []int should be an empty slice

	})

}

func TestUtils(t *testing.T) {

	t.Run("Test returnPointer", func(t *testing.T) {

		assert.Equal(t, "test", *returnPointer("test"))
		assert.Equal(t, 1234, *returnPointer(1234))
		assert.Equal(t, []string{"hello"}, *returnPointer([]string{"hello"}))
		assert.Equal(t, true, *returnPointer(true))

	})

	// t.Run("Test deepCompare", func(t *testing.T) {

	// 	type nestedObj struct {
	// 		ValN1 string
	// 		Val2 []bool
	// 	}
	// 	type obj struct{
	// 		Val1 *string
	// 		Val2 int
	// 		Val3 *nestedObj
	// 	}

	// 	obj1 := obj{
	// 		Val1: returnPointer("Hello"),
	// 		Val2: 1234,
	// 		Val3: returnPointer(nestedObj{
	// 			"Yes",
	// 			[]bool{true, false, true, true, false},
	// 		}),
	// 	}

	// 	obj2 := obj{
	// 		Val1: returnPointer("Hello"),
	// 		Val2: 1234,
	// 		Val3: returnPointer(nestedObj{
	// 			"Yes",
	// 			[]bool{true, false, true, true, false},
	// 		}),
	// 	}

	// 	obj3 := obj{
	// 		Val1: returnPointer("Helloo"),
	// 		Val2: 1234,
	// 		Val3: returnPointer(nestedObj{
	// 			"Yes",
	// 			[]bool{true, false, true, true, false},
	// 		}),
	// 	}

	// 	assert.True(t, deepCompare(obj1, obj2))
	// 	assert.False(t, deepCompare(obj1, obj3))

	// })

}
