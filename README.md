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
BenchmarkPackCalculatorNegativeWithDefault-12       	19926258	        59.6 ns/op	      64 B/op	       2 allocs/op
BenchmarkPackCalculator0WithDefault-12              	19726498	        60.3 ns/op	      64 B/op	       2 allocs/op
BenchmarkPackCalculator1WithDefault-12              	 6409830	       176 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator250WithDefault-12            	 6720063	       179 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator500WithDefault-12            	 6687548	       177 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator1000WithDefault-12           	 6699860	       185 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator12001WithDefault-12          	 4577499	       254 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator12001WithLarge-12            	 6715616	       175 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator47501043056WithDefault-12    	 4263607	       281 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator47501043056WithLarge-12      	 5175243	       229 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator12001WithSmall-12            	 4733602	       236 ns/op	     240 B/op	       4 allocs/op
BenchmarkPackCalculator250WithAllNegatives-12       	 8281035	       142 ns/op	     112 B/op	       4 allocs/op
BenchmarkPackCalculator250WithSomeNegatives-12      	 6288799	       184 ns/op	     240 B/op	       4 allocs/op
PASS
ok  	pack-sales-calculator/calculator	24.407s
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
