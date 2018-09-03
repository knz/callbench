# callbench

Microbenchmarks to measure the costs of calling conventions in Go and C++.

Included:

- `nargs`: measure the impact of adding more arguments to a
  function. Measure explict arguments vs. variable argument lists.

- `ret`: measure the impact of returning more values.

- `err`: compare error returns vs exceptions.
