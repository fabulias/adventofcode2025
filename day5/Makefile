build-full:
	clang++ -std=c++17 -O0 -g -Wall -Wextra \
      -fsanitize=address,undefined \
      problem.cpp -o problem

build:
	clang++ -std=c++17 problem.cpp -o problem