package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
type TestCase struct {
	input    string
	expected string
}

func TestKeyGenH (t *testing.T) {
	for _, testCase := range []TestCase{
		{
			input:    "https://classify-web.herokuapp.com/api/keygen?length=32?symbols=1",
			expected: "https://classify-web.herokuapp.com/api/keygen?length=32?symbols=1",
		},
	}
}
t.run(TestCase.name, func(t *testing.T) {
	result, err := KeyGenH(testCase.input)
	if TestCase.expected == "" {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, testCase.expected, result)
	}	

} 


//assert := assert.new(t)
//var test = [] struct.
// Model es el dominio en nuestro problema/ejercicio, nuestro caso el modelo es la estructura de datos encryptData.