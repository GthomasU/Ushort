package tests

import (
	"Ushort/services"
	"reflect"
	"testing"
)

var serviceShortener services.ServiceShortener

func TestCreateShortUrl(t *testing.T) {
	testCases := []struct {
		Name          string
		OriginalUrl   string
		ExpectedError error
		ExpectedUrl   string
		sslActive     bool
	}{
		{
			Name:          "Generar url exitosamente",
			OriginalUrl:   "https://www.wikipedia.org",
			ExpectedError: nil,
			ExpectedUrl:   "https://localhost:3000/r/abcdefghi",
			sslActive:     true,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			serviceShortener.SetSslActive(tc.sslActive)
			shortenedUrl, err := serviceShortener.CreateShortUrl(tc.OriginalUrl)
			if tc.ExpectedError == nil {
				if *shortenedUrl != tc.ExpectedUrl {
					t.Errorf("expected %v, got %v", tc.ExpectedUrl, *shortenedUrl)
				}
			} else {
				if reflect.TypeOf(err) != reflect.TypeOf(tc.ExpectedError) {
					t.Errorf("expected error %v, got error: %v", tc.ExpectedError.Error(), err.Error())
				}
			}
		})
	}
}

func TestGetOriginalUrl(t *testing.T) {
	testCases := []struct {
		Name          string
		urlId         string
		ExpectedError error
		ExpectedUrl   string
	}{
		{
			Name:          "Obtener url correctamente",
			urlId:         "abcdefghi",
			ExpectedError: nil,
			ExpectedUrl:   "https://www.wikipedia.org",
		},
		{
			Name:          "Obtener url que no existe",
			urlId:         "abcd",
			ExpectedError: services.UrlNotFound{},
			ExpectedUrl:   "https://www.wikipedia.org",
		},
		{
			Name:          "Obtener url con urlId vacía",
			urlId:         "",
			ExpectedError: services.InvalidUrlId{},
			ExpectedUrl:   "https://www.wikipedia.org",
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			originalUrl, err := serviceShortener.GetOriginalUrl(tc.urlId)
			if err == nil {
				if originalUrl != tc.ExpectedUrl {
					t.Errorf("expected url: %s got %s", tc.ExpectedUrl, originalUrl)
				}
			} else {
				if reflect.TypeOf(err) != reflect.TypeOf(tc.ExpectedError) {
					t.Errorf("expected error %v, got error: %v", tc.ExpectedError.Error(), err.Error())
				}
			}
		})
	}
}

func TestRemoveOriginalUrl(t *testing.T) {

	testClases := []struct {
		name           string
		urlId          string
		expectedError  error
		expectedResult bool
	}{
		{
			name:           "Remover urlId correctamente",
			urlId:          "abcdefghi",
			expectedError:  nil,
			expectedResult: true,
		},
		{
			name:           "Remover urlId invalida",
			urlId:          "",
			expectedError:  services.InvalidUrlId{},
			expectedResult: false,
		},
		{
			name:           "Remover urlId que no existe",
			urlId:          "abcd",
			expectedError:  nil,
			expectedResult: false,
		},
	}

	for i := range testClases {
		tc := testClases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := serviceShortener.RemoveOriginalUrl(tc.urlId)
			if tc.expectedError == nil {
				if tc.expectedResult != result {
					t.Errorf("expected result: %t got %t", tc.expectedResult, result)

				}
			} else {
				if reflect.TypeOf(err) != reflect.TypeOf(tc.expectedError) {
					t.Errorf("expected error %v, got error: %v", tc.expectedError.Error(), err.Error())

				}
			}

		})

	}
}

func TestUpdateOriginalUrl(t *testing.T) {
	testClases := []struct {
		name           string
		urlId          string
		originalUrl    string
		expectedResult bool
		expectedError  error
	}{
		{
			name:           "Update url correctamente",
			urlId:          "abcdefghi",
			originalUrl:    "https://www.youtube.com",
			expectedResult: true,
			expectedError:  nil,
		},
		{
			name:           "Update url con urlId invalida",
			urlId:          "",
			originalUrl:    "https://www.youtube.com",
			expectedResult: false,
			expectedError:  services.InvalidUrlId{},
		},
	}
	for i := range testClases {
		tc := testClases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := serviceShortener.UpdateOriginalUrl(tc.urlId, tc.originalUrl)
			if tc.expectedError == nil {
				if result != tc.expectedResult {
					t.Errorf("expected result: %t got %t", tc.expectedResult, result)
				}
			} else {
				if reflect.TypeOf(err) != reflect.TypeOf(tc.expectedError) {
					t.Errorf("expected error %v, got error: %v", tc.expectedError.Error(), err.Error())
				}
			}
		})
	}
}
