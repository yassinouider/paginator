package paginator_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/yassinouider/paginator"
)

func Example() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		p := paginator.New(r)

		customers, total, _ := getFakeDataStore().FindCustomers(p.Limit(), p.Offset())

		res := map[string]interface{}{
			"data":       customers,
			"pagination": p.SetCount(len(customers)).SetTotal(total),
		}

		b, _ := json.MarshalIndent(res, "", " ")

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}

	req := httptest.NewRequest("GET", "http://example.com/customers?page=3&per_page=2", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	// Output:
	//{
	//  "data": [
	//   "User 5",
	//   "User 6"
	//  ],
	//  "pagination": {
	//   "total": 7,
	//   "count": 2,
	//   "per_page": 2,
	//   "current_page": 3,
	//   "total_page": 4
	//  }
	//}
}

func getFakeDataStore() *fakeDataStore {
	return &fakeDataStore{}
}

type fakeDataStore struct{}

func (s *fakeDataStore) FindCustomers(limit, offset int) ([]string, int, error) {
	customers := []string{
		"User 1",
		"User 2",
		"User 3",
		"User 4",
		"User 5",
		"User 6",
		"User 7",
	}

	var start = offset
	var end = offset + limit

	if start > len(customers) {
		start = len(customers)
	}

	if end > len(customers) {
		end = len(customers)
	}

	return customers[start:end], len(customers), nil
}

func ExampleOffset() {
	page := 3
	limit := 10

	result := paginator.Offset(page, limit)

	fmt.Printf("%d", result)

	// Output: 20
}
