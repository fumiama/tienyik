package tienyik

import (
	"encoding/binary"
	"net/url"
	"testing"
)

func TestEUrlParams(t *testing.T) {
	const aesplain = "moduleCode=DESKTOP_MSGCENTER"
	var (
		rawkey = []uint32{
			2004378729, 1936745065, 1933079672, 1970627951,
			842425958, 1932686949, 1903374648, 1936290669,
		}
		key [32]byte
	)
	for i, k := range rawkey {
		binary.BigEndian.PutUint32(key[i*4:(i+1)*4], k)
	}
	t.Log(string(key[:])) // wxdispbis8txuueo26ffs2veqs18sism
	tya := NewAES(key[:])
	params := tya.EUrlParams(url.Values{
		"moduleCode": {"DESKTOP_MSGCENTER"},
	})
	t.Log(params)
	q, err := tya.ParseQuery(params)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range q {
		plainValue := k + "=" + v[0]
		if plainValue != aesplain {
			t.Fatal("expect", aesplain, "got", plainValue)
		}
	}
}

func TestEUrlParamsMultiple(t *testing.T) {
	var (
		rawkey = []uint32{
			2004378729, 1936745065, 1933079672, 1970627951,
			842425958, 1932686949, 1903374648, 1936290669,
		}
		key [32]byte
	)
	for i, k := range rawkey {
		binary.BigEndian.PutUint32(key[i*4:(i+1)*4], k)
	}
	tya := NewAES(key[:])

	testCases := []struct {
		name     string
		params   url.Values
		expected map[string]string
	}{
		{
			name: "single parameter",
			params: url.Values{
				"userId": {"12345"},
			},
			expected: map[string]string{
				"userId": "12345",
			},
		},
		{
			name: "multiple parameters",
			params: url.Values{
				"userId":   {"12345"},
				"userName": {"testUser"},
				"status":   {"active"},
			},
			expected: map[string]string{
				"userId":   "12345",
				"userName": "testUser",
				"status":   "active",
			},
		},
		{
			name: "special characters",
			params: url.Values{
				"email":   {"test@example.com"},
				"message": {"Hello World!"},
			},
			expected: map[string]string{
				"email":   "test@example.com",
				"message": "Hello World!",
			},
		},
		{
			name: "chinese characters",
			params: url.Values{
				"name": {"张三"},
				"city": {"北京"},
			},
			expected: map[string]string{
				"name": "张三",
				"city": "北京",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := tya.EUrlParams(tc.params)
			q, err := tya.ParseQuery(params)
			if err != nil {
				t.Fatal(err)
			}
			for key, expectedValue := range tc.expected {
				if vals, ok := q[key]; ok && len(vals) > 0 {
					if vals[0] != expectedValue {
						t.Fatalf("key %s: expect %s, got %s", key, expectedValue, vals[0])
					}
				} else {
					t.Fatalf("key %s not found in query", key)
				}
			}
		})
	}
}
