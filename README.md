# common-utils

本人在 golang 程序開發經常使用的工具

## Modules

- request - 仿 axios 的請求 API
- regex - 目前只有 utils 類
- datetime - 目前只有 utils 類
- array - 目前只有 utils 類
- str - 目前只有 utils 類
- stream - 仿 Functional Programming 的 collection API

utils 類別的 package 將不作介紹，可自行參考源碼

### Request

基本請求

```go
requester := request.New()
var resp Product
_, err := requester.Get("https://dummyjson.com/products/1", &resp)
if err != nil {
    panic(err)
}
```

原始請求

```go
requester := request.New().Raw()
resp, err := requester.Get("https://dummyjson.com/products/1")
if err != nil {
    panic(err)
}
r := resp.Resp // *http.Response
```

帶有 base url 的 requester

```go
requester := request.New(
    request.WithBaseUrl("https://dummyjson.com"),
    // headers（如有)
    request.WithHeaders(map[string]string{
        "Content-Type": "application/json"
    }),

    //cookies (如有)
    request.WithCookies(map[string]string{
        "token": "awhdiawhdiawhidahwi"
    }),

    // 自定義 http 客戶端
    request.WithClient(&http.Client{}),

    // 自定義逾時
    request.WithTimeout(10 * time.Second)
)

var products []Product
var product Product
_, err := requester.Get("/products", &products)
_, err = requester.Get("/product/1", &product)
```

帶有 query string 的請求

```go
requester.Get("https://dummyjson.com/products", request.Query(map[string]interface{}{
    "page": 1,
    "pageSize": 30
}))
```

帶有 payload 數據的請求

```go
// 默認是json請求
requester.Post("https://dummyjson.com/products", request.Data(map[string]interface{}{
    "title": "Dummy Product",
    "price": 234,
}))
```

改用 form url encode 作為 payload 數據

```go
requester.Post("https://dummyjson.com/products",
    request.Data(map[string]interface{}{
        "title": "Dummy Product",
        "price": 234,
    }),
    request.DataEncoder(request.FormUrlEncodedEncoder) // 庫內置的 encoder
)
```

自定義 encoder/decoder

```go

myEncoder := func(data map[string]interface{}) (io.Reader, error) {
    buffer := new(bytes.Buffer)
	enc := gob.NewEncoder(buffer)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

myDecoder := func(data []byte, res interface{}) error {
    buffer := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buffer)
	return dec.Decode(res)
}

// 設置為默認的 encoder/decoder
requester := request.New(
    request.WithBaseUrl("https://dummyjson.com"),
    request.WithDefaultEncoder(myEnoder),
    request.WithDefaultDecoder(myDecoder),
)

var products []Product
_, err := requester.Get("/products", &products)
var product Product
_, err = requester.Get("/products/1", &product, request.WithEncoder(request.JsonEncoder)) // 把此請求改用 JsonEncoder
```

請求/回應攔截

```go
requester := request.New(
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
```


### Stream

大致分為兩種

Stream[T] - 可以從 array / set 中初始化
MapStream[K, V] - 可以從 map 中初始化

懶，直接上 test function

**Steam[T]**

```go
  func TestStudentArr(t *testing.T) {

	students := []Student{
		{"Alice", 20},
		{"Bob", 18},
		{"Charlie", 19},
		{"David", 20},
		{"Eve", 18},
		{"Frank", 19},
		{"Grace", 20},
		{"Heidi", 18},
		{"Ivan", 19},
		{"Judy", 20},
		{"Kevin", 18},
		{"Lily", 19},
		{"Mallory", 20},
		{"Nate", 18},
		{"Oliver", 19},
		{"Peggy", 20},
		// give me some non-adult
		{"Quentin", 17},
		{"Romeo", 17},
		{"Steve", 17},
		{"Trent", 17},
		{"Uma", 17},
		{"Victor", 17},
		{"Walter", 17},
		{"Xavier", 17},
		{"Yvonne", 17},
		{"Zack", 17},
	}

	// filter

	adults := From(students).Filter(func(s Student) bool {
		return s.Age >= 18
	})

	t.Log("adults:", adults)

	// anyMatch

	any20 := adults.AnyMatch(func(s Student) bool {
		return s.Age == 20
	})

	t.Log("any20:", any20)

	// combine

	r := From(students).
		Filter(func(s Student) bool {
			return s.Age == 20
		}).
		AnyMatch(func(s Student) bool {
			return s.Name == "Alice"
		})

	t.Log("20 and Alice:", r)

	// map

	names := MapTo(From(students), func(s Student) string {
		return s.Name
	})

	t.Log("names:", names)


	s := From(students).Filter(func(s Student) bool {
		return s.Age > 17
	})

	m1 := ToMapStream(s, func(s Student)(string, int)  {
		return s.Name, s.Age
	})

	t.Logf("%+v", m1.ToMap())
	
	m2 := ToMapStream(s, func(s Student)(string, Student)  {
		return s.Name, s
	})

	t.Logf("%+v", m2.ToMap())

}
```


**MapStream[K, V]**

```go
func TestMapStream(t *testing.T) {
	m := map[string]string{
		"1": "a",
		"2": "b",
		"3": "c",
	}

	ms := FromMap(m)
	r := ms.Filter(func(k, v string) bool {
		return k == "1"
	})

	t.Log(r.ToMap())

	r = ms.Map(func(k, v string) (string, string) {
		return k, v + "!"
	})

	t.Log(r.ToMap())

	// map to entries -> Stream[MapEntry[K, V]]
	r2 := ms.Entries().Find(func(e MapEntry[string, string]) bool {
		return e.Key == "1"
	})
	
	t.Logf("%v: %v", r2.Key, r2.Value)
	
}
```