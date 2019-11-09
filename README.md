# Pack Sales Calculator

This is a go approach to solving the pack sales calculator challenge.

It provides both a web interface throguh AWS Lambda and the Serverless Framework, and a CLI interface by building the `cmd/main.go` file as a binary to run.

The algorithm used rounds up the order to the nearest number of the smallest packs that fit in (for example, an order of 12001 becomes 12250), then runs a greedy algorithm on the new order number to calculate the required packs.

As I have done this test before, I focused on making the algorithm as efficient as possible, and have included tests and benchmarks. The algorithm only makes a total of 4 memory allocations and is memory efficient.

```
go test -bench=. -benchmem ./...
```

These are the results received on my rather old 2013 macbook pro:

```
BenchmarkPackCalculator0WithDefault-4             	19853205	        59.3 ns/op	      48 B/op	       1 allocs/op
BenchmarkPackCalculator1WithDefault-4             	 3933694	       332 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator250WithDefault-4           	 3889850	       301 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator500WithDefault-4           	 3927804	       305 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator1000WithDefault-4          	 3771909	       310 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator12001WithDefault-4         	 3026665	       395 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator12001WithLarge-4           	 3491788	       311 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator47501043056WithDefault-4   	       8	 136161652 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator47501043056WithLarge-4     	     836	   1411947 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator12001WithSmall-4           	  274621	      4024 ns/op	     240 B/op	       4 allocs/op
PASS
ok  	pack-sales-calculator/calculator	14.993s
```

## CLI Usage

```
# The CLI requires both the order size, and the packs to be passed in.
# The first argument passed in is the order size, and the 2nd a comma seperated list of packs.
>> go run cmd/main.go 12001 250,500,1000,2000,5000
map[250:1 2000:1 5000:2]
```

## Lambda Usage

You can deploy the app by using the serverless framework and the below command:

```
serverless deploy -v --aws-s3-accelerate
```

The web service exposes one endpoint, a post endpoint that accpets 2 query params:

- `order` *REQUIRED* The order number. This must be a positive non-zero integer.
- `packs` *OPTIONAL* The packs to use. This must be a comma seperated list of integers. Default: 250,500,1000,2000,5000

Example With Default Packs:

```
curl -X POST "https://hubfys30s9.execute-api.eu-west-2.amazonaws.com/dev/calculate?order=12001"
```

Example With Custom Packs:

```
curl -X POST "https://hubfys30s9.execute-api.eu-west-2.amazonaws.com/dev/calculate?order=51&packs=5,10,20"
```