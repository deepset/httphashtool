package script

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetHashFromURL(t *testing.T) {

	var tests = []struct {
		url  string
		want string
	}{
		{"www.google.com", "http://www.google.com"},
		{"twitter.com", "http://twitter.com"},
		{"xyzadjust.com", ""},
	}

	tCount, pass, fail := 0, 0, 0
	for _, tc := range tests {

		tCount++
		testname := tc.url

		t.Run(testname, func(t *testing.T) {
			str, _ := GetHashFromURL(tc.url)
			if strings.Contains(str, tc.want) == false {
				fail++
				t.Errorf("got %s, want %s", str, tc.want)
			} else {
				pass++
			}
		})
	}
	t.Logf("\nResults:\nTotal test cases executed : %d \nTest Passed : %d \nTest Failed : %d\n", tCount, pass, fail)
}

func TestCreateWorkers(t *testing.T) {

	noOfWorkers := 2
	var tests = []struct { //Array of structure with the initialized member
		serverURL []string
		want      error
	}{
		{[]string{"adjust.com", "google.com", "facebook.com", "yahoo.com", "yandex.com"}, nil},
		{[]string{"reddit.com/r/funny", "reddit.com/r/notfunny", "baroquemusiclibrary.com"}, nil},
	}

	tCount, pass, fail := 0, 0, 0
	for idx, tc := range tests {

		tCount++
		testname := fmt.Sprintf("%d", idx)

		t.Run(testname, func(t *testing.T) {
			result := CreateWorkers(tc.serverURL, noOfWorkers)
			if result != nil {
				fail++
				t.Errorf("got %s, want %s", result, tc.want)
			} else {
				pass++
			}
		})
	}
	t.Logf("\nResults:\nTotal test cases executed : %d \nTest Passed : %d \nTest Failed : %d\n", tCount, pass, fail)

}
