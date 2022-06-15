package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name    string
	params  string
	returns string
}

func TestHelloName(t *testing.T) {
	result := HelloName("ebi")

	if result != "Halo ebi" {
		t.Fatal("return must be Hello ebi")
	}
}

func TestHelloNames(t *testing.T) {
	result := HelloName("ebi")
	assert.Equal(t, "Hello ebi", result, "salah")
}

func TestHelloSubtest(t *testing.T) {
	t.Run("params ebi", func(t *testing.T) {
		result := HelloName("ebi")
		assert.Equal(t, "Hello ebi", result, "salah")
	})

	t.Run("params viky", func(t *testing.T) {
		result := HelloName("vikcy")
		assert.Equal(t, "Hello vikcys", result, "salah")
	})

	t.Run("params tema", func(t *testing.T) {
		result := HelloName("tema")
		assert.Equal(t, "Hello tema", result, "salah")
	})
}

func TestHelloTable(t *testing.T) {
	var testTb = []testTable{
		{
			name:    "ebi",
			params:  "ebi",
			returns: "Hello ebi",
		},
		{
			name:    "devri",
			params:  "devri",
			returns: "Hello devri",
		},
	}

	for _, val := range testTb {
		t.Run(val.name, func(t *testing.T) {
			result := HelloName(val.params)
			assert.Equal(t, val.returns, result, "salah")
		})
	}
}
