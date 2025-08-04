package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestHitungVolume(t *testing.T) {
    assert.Equal(t, 30, 30, "perhitungan volume salah")
}

 
func TestHitungLuas(t *testing.T) {
	var hasil = 2 *3
    assert.Equal(t, hasil, 10, "perhitungan luas salah")
}

