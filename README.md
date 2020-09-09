# prime-number-go
Prime number checker written in Go (concurrency, goroutine)

# Objectives
Excercise of concurrency and goroutines.

# Experiment and Results

## Experiment Settings

Compare two algorithms to check a prime number:
- sequential algorithm
- parallel algorithm with interprocess communication (ICP)

Target numbers:
- large prime 1
  - 100109100129100151
- large prime 2
  - 100109100129162907
- large non-prime 1
  - 100109100129100369
- large non-prime 2
  - 100109100129101027

Number of Parallel processes: 4

## Results

Parallel algorithm with ICP is faster than sequential one.

prime -> about 4 times faster

non-prime -> 6~10 times faster

||Sequential|IPC|
|--|--|--|
|large prime 1|1.17s|0.309s|
|large prime 2|1.16s|0.317s|
|large non-prime 1|0.362s|0.0655s|
|large non-prime 2|0.641s|0.0519s|

# Reference
- [High Performance Python (O'REILLY, Japanese translation of the 1st edition)](https://www.amazon.co.jp/dp/4873117402/ref=cm_sw_r_tw_dp_x_zPkvFbPPMC09T )
  - Mainly chapter 9 (multiprocessiong)
  - [Source codes](https://github.com/mynameisfiber/high_performance_python/tree/master/09_multiprocessing/prime_validation)
  - [2nd edition (English)](https://www.amazon.co.jp/dp/B087YTVL8F/ref=cm_sw_r_tw_dp_x_wqkvFbP0K1XHK)
