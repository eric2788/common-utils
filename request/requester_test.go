package request

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Product struct {
	Id                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              float64  `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float32  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

func createRequester(t *testing.T) *Requester {
	return New(
		WithBaseUrl("https://dummyjson.com"),
		AddHeader("Content-Type", "application/json"),
		AddHeader("Accept", "application/json"),
		WithTimeout(10*time.Second),
		AddRequestIntercepter(func(r *http.Request) error {
			// do something with request
			t.Logf("prepare to request: %s", r.URL.String())
			return nil
		}),
		AddResponseIntercepter(func(r *http.Response) error {
			// do something with response
			t.Logf("status code from %s: %d", r.Request.URL.String(), r.StatusCode)
			return nil
		}),
	)
}

func TestPost(t *testing.T) {

	r := createRequester(t)

	var resp Product
	_, err := r.Post("/products/add", &resp,
		Data(map[string]interface{}{
			"title":              "Product test",
			"description":        "This is a product test",
			"price":              100,
			"discountPercentage": 0,
			"rating":             0,
			"stock":              100,
			"brand":              "Test",
			"category":           "test",
			"thumbnail":          "https://dummyimage.com/600x400/000/fff",
			"images": []string{
				"https://dummyimage.com/600x400/000/fff",
				"https://dummyimage.com/600x400/000/fff",
				"https://dummyimage.com/600x400/000/fff",
			},
		}),
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("response: %+v", resp)

	assert.Equal(t, 101, resp.Id)
	assert.Equal(t, "Product test", resp.Title)
	assert.Equal(t, "This is a product test", resp.Description)
	assert.Equal(t, "Test", resp.Brand)
	assert.Equal(t, "test", resp.Category)
	assert.Equal(t, 100, resp.Stock)
	assert.Equal(t, float64(100), resp.Price)
	assert.Equal(t, float64(0), resp.DiscountPercentage)
}

func TestGet(t *testing.T) {

	r := createRequester(t)

	var product Product
	_, err := r.Get("/products/1", &product)
	if err != nil {
		t.Error(err)
	}

	t.Logf("product: %+v", product)

	assert.Equal(t, 1, product.Id)
	assert.Equal(t, "iPhone 9", product.Title)
	assert.Equal(t, "An apple mobile which is nothing like apple", product.Description)
	assert.Equal(t, "Apple", product.Brand)
	assert.Equal(t, "smartphones", product.Category)
	assert.Equal(t, 94, product.Stock)
	assert.Equal(t, float64(549), product.Price)
	assert.Equal(t, 12.96, product.DiscountPercentage)
}
