package services

import (
	"os"
	"reflect"
	"testing"
)

var urlGenerator UrlGenerator

func TestMain(m *testing.M) {
	urlGenerator = NewUrlGenerator()
	code := m.Run()
	os.Exit(code)
}

func TestCreateRandomString(t *testing.T) {
	testCases := []struct {
		Name          string
		N             int
		ErrorExpected error
	}{
		{
			Name:          "Generar string de 5 caracteres",
			N:             5,
			ErrorExpected: nil,
		},
		{
			Name:          "Generar string con longitud invalida",
			N:             0,
			ErrorExpected: InvalidLength{Length: 0},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			randomString, err := urlGenerator.CreateRandomString(tc.N)
			randomStringLength := len(*randomString)
			if tc.ErrorExpected == nil {
				if randomStringLength != tc.N {
					t.Errorf("expected %v, got %v", tc.N, randomStringLength)
				}
			} else {
				if reflect.TypeOf(err) != reflect.TypeOf(tc.ErrorExpected) {
					t.Errorf("expected error %v, got error: %v", tc.ErrorExpected.Error(), err.Error())
				}
			}
		})
	}
}
