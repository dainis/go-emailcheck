package emailcheck

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestRegex(t *testing.T) {
	c := CheckRegex("test@test.com")
	assert.True(t, c)

	c = CheckRegex("yolo@trolololo")
	assert.False(t, c)
}

func TestRecords(t *testing.T) {
	r, e := CheckRecords("test@gmail.com")

	assert.Nil(t, e)
	assert.True(t, r)

	n := strconv.FormatInt(time.Now().Unix(), 16)

	r, e = CheckRecords("test@somethingsomethingdangerous" + n + ".com")

	assert.Nil(t, e)
	assert.False(t, r)
}

func TestConnectivity(t *testing.T) {
	r, e := CheckConnectivity("test@gmail.com")

	assert.Nil(t, e)
	assert.True(t, r)
}

func TestCheck(t *testing.T) {
	r, e := Check("test@gmail.com")

	assert.Nil(t, e)
	assert.True(t, r)
}
