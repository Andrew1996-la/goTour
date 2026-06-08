package multiple

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMultiple(t *testing.T) {
	res := Multiply(2, 3)

	//assert и require
	assert.Equal(
		t,
		6,
		res,
	)
}

/*
популярные assert'ы: 
assert.Equal(t, 10, result)
assert.NotEqual(t, 10, result)

assert.Nil(t, err)
assert.NoError(t, err)
assert.Error(t, err)

assert.True(t, isValid)
assert.False(t, isValid)

users := []string{"Bob", "John"}
assert.Len(t, users, 2)

assert.Contains(
	t,
	"hello world",
	"world",
)
*/