# Pack Sales Calculator

This is a go approach to solving the pack sales calculator challenge.

It provides both a web interface throguh AWS Lambda and the Serverless Framework, and a CLI interface by building the `cmd/main.go` file as a binary to run.

The algorithm used rounds up the order to the nearest number of the smallest packs that fit in (for example, an order of 12001 becomes 12250), then runs a greedy algorithm on the new order number to calculate the required packs.

As I have done this test before, I focused on making the algorithm as efficient as possible, and have included tests and benchmarks. The algorithm only makes a total of 4 memory allocations and is memory efficient.

```
go test -bench=. -benchmem ./...
```

These are the results received are:

```
goos: darwin
goarch: amd64
pkg: pack-sales-calculator/calculator
BenchmarkPackCalculatorNegativeWithDefault-8      	14820410	        79.3 ns/op	      64 B/op	       2 allocs/op
BenchmarkPackCalculator0WithDefault-8             	14581399	        79.3 ns/op	      64 B/op	       2 allocs/op
BenchmarkPackCalculator1WithDefault-8             	 4616272	       242 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator250WithDefault-8           	 4864417	       243 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator500WithDefault-8           	 4673331	       247 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator1000WithDefault-8          	 4915204	       246 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator12001WithDefault-8         	 3785590	       317 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator12001WithLarge-8           	 4900058	       243 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator47501043056WithDefault-8   	       9	 114751263 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator47501043056WithLarge-8     	    1081	   1133706 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator12001WithSmall-8           	  347701	      3158 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator250WithAllNegatives-8      	 6015590	       226 ns/op	     112 B/op	       4 allocs/op
BenchmarkPackCalculator250WithSomeNegatives-8     	 4249892	       353 ns/op	     240 B/op	       4 allocs/op
PASS
ok  	pack-sales-calculator/calculator	18.411s
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
curl -X POST "https://rc8nn4w5fb.execute-api.eu-west-2.amazonaws.com/dev/calculate?order=12001"
```

Example With Custom Packs:

```
curl -X POST "https://rc8nn4w5fb.execute-api.eu-west-2.amazonaws.com/dev/calculate?order=51&packs=5,10,20"
```