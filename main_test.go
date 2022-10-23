package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "fmt"

    "github.com/stretchr/testify/assert"
  )

func TestMainHandler (t *testing.T){
	type args struct {
		expected string
	}
	tests := []struct {
		name 	string
		want 	string
	}{
		{
			name: "Happy scenario",
			want: "Hello from Chi Router !",
		},
		{
			name: "Unhappy scenario",
			want: "wrong endpoint message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil{
				fmt.Println(err)
			}

			w := httptest.NewRecorder()
			handler := http.HandlerFunc(MainHandler)
			handler.ServeHTTP(w, r)

			if status := w.Code; status != http.StatusOK{
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			if tt.want != w.Body.String(){
				assert.NotEqual(t, tt.want, w.Body.String(), "they should not be equal")
				return
			}

			assert.Equal(t, tt.want, w.Body.String(), "they should be equal")
		})

		
	}
}