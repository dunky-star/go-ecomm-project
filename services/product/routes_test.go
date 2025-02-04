package product

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dunky-star/ecomm-proj/payloads"
	"github.com/gorilla/mux"
)

func TestProductServiceHandlers(t *testing.T) {
	productStore := &mockProductStore{}
	userStore := &mockUserStore{}
	handler := NewHandler(productStore, userStore)

	t.Run("should handle get products", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleGetProducts).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should fail if the product ID is not a number", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products/abc", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products/{productID}", handler.handleGetProduct).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should handle get product by ID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/products/42", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products/{productID}", handler.handleGetProduct).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("should fail creating a product if the payload is missing", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/products", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleCreateProduct).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should handle creating a product", func(t *testing.T) {
		payload := payloads.CreateProductPayload{
			Name:        "test",
			Price:       100,
			Image:       "test.jpg",
			Description: "test description",
			Quantity:    10,
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleCreateProduct).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockProductStore struct{}

func (m *mockProductStore) GetProductByID(productID int) (*payloads.Product, error) {
	return &payloads.Product{}, nil
}

func (m *mockProductStore) GetProducts() ([]*payloads.Product, error) {
	return []*payloads.Product{}, nil
}

func (m *mockProductStore) CreateProduct(product payloads.CreateProductPayload) error {
	return nil
}

func (m *mockProductStore) UpdateProduct(product payloads.Product) error {
	return nil
}

func (m *mockProductStore) GetProductsByID(ids []int) ([]payloads.Product, error) {
	return []payloads.Product{}, nil
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByID(userID int) (*payloads.User, error) {
	return &payloads.User{}, nil
}

func (m *mockUserStore) CreateUser(user payloads.User) error {
	return nil
}

func (m *mockUserStore) GetUserByEmail(email string) (*payloads.User, error) {
	return &payloads.User{}, nil
}