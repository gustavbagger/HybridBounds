# Hybrid bounds for prime divisors
This library provides an algorithm for computing a non-trivial lower bound for a positive integer x, depending only on the amount of factors `\omega` of `x^n-1` for some fixed integer exponent `n>1`.
## Motivation
Bounds dependant on `\omega(x^n-1)` have applications to numerous existence results concerning primitive elements in `\FF_q` extensions. Trivially, one has that `x>=\Prod_{i=1}^\omega p_i` where `p_i` is the ith smallest prime. By applying our hybridised bounds, we can sometimes
obtain immediate improvements to previous known results. In general, sieves reliant on `\omega` are natural settings for these hybridised bounds. See [Hybrid bounds for prime divisors](https://arxiv.org/abs/2412.00010) for more details and applications.
## Quick Start
```bash
# install using Go toolchain
go install github.com/gustavbagger/HybridBounds

# run
HybridBounds <n (int)>  <omega (int)>
```

## Contributing

### Clone the repo

```bash
git clone https://github.com/gustavbagger/HybridBounds
cd HybridBounds
```

### Build the compiled binary
```bash
go build
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.