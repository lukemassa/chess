package chess

import "testing"

func TestNewLocation(t *testing.T) {

	testCases := []struct {
		locationString string
		expectedFile   uint8
		expectedRank   uint8
	}{
		{
			"A1",
			0,
			0,
		},
		{
			"C5",
			2,
			4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.locationString, func(t *testing.T) {
			location := NewLocation(tc.locationString)
			if location.file != tc.expectedFile {
				t.Errorf("Expected file %d, got file %d", tc.expectedFile, location.file)
			}
			if location.rank != tc.expectedRank {
				t.Errorf("Expected rank %d, got rank %d", tc.expectedRank, location.rank)
			}
			if location.String() != tc.locationString {
				t.Errorf("Expected format string of %s, got %s", location.String(), tc.locationString)
			}

		})
	}
}
