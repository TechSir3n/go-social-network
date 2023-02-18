package v1

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"social_network/internal/api/v1/models"
	"testing"
)

func TestLogin(t *testing.T) {
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
	

	handler.ServeHTTP(res, req)
	defer res.Result().Body.Close()

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Couldn't read response: %v", err)
	}

}

func TestSignUp(t *testing.T) {
	req, err := http.NewRequest("GET", "/registration", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(SignUp)

	handler.ServeHTTP(res, req)
	defer res.Result().Body.Close()

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	user := models.User{
		Name:     "Anatoli",
		Password: "Anatoli12345",
		Email:    "Anatoli@mail.ru",
		ConfirmPassword: "Anatoli2345",
	}
	
	_, err = db.CreateUser(context.Background(), user)
	if err != nil {
		t.Errorf("Unable create user: %v ", err)
	}
}

func TestVerifyEmail(t *testing.T) {
	req, err := http.NewRequest("GET", "/verify", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(VerifyEmail)

	handler.ServeHTTP(res, req)
	defer res.Result().Body.Close()

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestResetPassword(t *testing.T) {
	req, err := http.NewRequest("GET", "/reset/password", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(ResetPassword)

	handler.ServeHTTP(res, req)
	defer res.Result().Body.Close()

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestAccessAdmin(t *testing.T) {
	req, err := http.NewRequest("GET", "/access/admin", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(AccessAdmin)

	handler.ServeHTTP(res, req)
	defer res.Result().Body.Close()
	
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
