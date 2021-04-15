### Benchmark experiment with Go

---
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz

------

__/sequential__

BenchmarkCountWords-4        	   21622	     55140 ns/op

BenchmarkCountWords1000-4    	   22098	     55124 ns/op

BenchmarkCountWords2000-4    	   12482	     98389 ns/op

BenchmarkCountWords3000-4    	    8599	    150289 ns/op

BenchmarkCountWords4000-4    	    5305	    237155 ns/op

BenchmarkCountWords5000-4    	    5586	    221677 ns/op

BenchmarkCountWords6000-4    	    4374	    282460 ns/op

BenchmarkCountWords7000-4    	    3681	    294271 ns/op

BenchmarkCountWords8000-4    	    3379	    325426 ns/op

BenchmarkCountWords9000-4    	    3034	    367464 ns/op

BenchmarkCountWords10000-4   	    2881	    410021 ns/op

----

__/concurrent__

BenchmarkCountWords1000-4    	   13071	     88327 ns/op

BenchmarkCountWords2000-4    	    7888	    163060 ns/op

BenchmarkCountWords3000-4    	    3980	    291070 ns/op

BenchmarkCountWords4000-4    	    3697	    320616 ns/op

BenchmarkCountWords5000-4    	    2944	    365759 ns/op

BenchmarkCountWords6000-4    	    2985	    400420 ns/op

BenchmarkCountWords7000-4    	    2959	    559283 ns/op

BenchmarkCountWords8000-4    	    2385	    510124 ns/op

BenchmarkCountWords9000-4    	    2140	    550706 ns/op

BenchmarkCountWords10000-4   	    1996	    610757 ns/op