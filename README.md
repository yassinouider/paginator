# paginator [![GoDoc](http://godoc.org/github.com/yassinouider/paginator?status.svg)](http://godoc.org/github.com/yassinouider/paginator)

**A simple package to paginate your data in Go.**

## Installation

```bash
$ go get github.com/yassinouider/paginator
```

## Usage

In your handler
```go
//This will automatically read the URL query parameter "page" and "per_page" in the request r, 
//if they are not present the default values will be chosen
p := paginator.New(r)
```


Example Server

```go
package main 

import "github.com/yassinouider/paginator"

func main() {
	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		p := paginator.New(r)

		customers, total, _ := getFakeDataStore().FindCustomers(p.Limit(), p.Offset())

		res := map[string]interface{}{
			"data":       customers,
			"pagination": p.SetCount(len(customers)).SetTotal(total),
		}

		b, _ := json.MarshalIndent(res, "", " ")

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	http.ListenAndServe(":8080", nil)
}
```

Example Client

```go
resp, err := http.Get("http://localhost:8080/customers?page=2&per_page=2")
if err != nil {
	panic(err)
}

defer resp.Body.Close()

body, err := ioutil.ReadAll(resp.Body)
if err != nil {
	panic(err)
}

fmt.Println(string(body))
```

Output

```json
{
  "data": [
    "User 5",
    "User 6"
  ],
  "pagination": {
    "total": 7,
    "count": 2,
    "per_page": 2,
    "current_page": 3,
    "total_page": 4
  }
}
```


You can change the default values just like this

```go
//default value = 50
paginator.PerPage = 20

//default value = 1000
paginator.PerPageMax = 100 

//default value = "page"
paginator.PageName = "p"// http://localhost:8080/customers?p=3

//default value = "per_page"
paginator.PerPageName = "limit" // http://localhost:8080/customers?limit=5
```
