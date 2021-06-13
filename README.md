# Random_Forest

A Random Forest that predicts whether a tumor is malignant or benign, implemented from scratch.

Based on the Wisconsin breast cancer diagnosis dataset.


## Getting Started

First you need to have your golang workspace set up on your machine.
Then clone this repo into your go-workspace/src/ folder. <br>
```git clone https://github.com/dfinnis/Random_Forest.git; cd Random_Forest```

Download dependencies. <br>
```go get -d ./...```

To run. <br>
```go run main.go```

Alternatively, build & run the binary. <br>
```go build; ./Random_Forest```

Default behaviour is to split the data into training & test sets, train a random forest on the training set, & show metrics for training & test sets.


## Flags

### -d --depth

```go run main.go -d 5```

Provide addtional argument DEPTH integer, maximum tree depth.
The default is 42, in essence infinite as the tree categorizes the training set perfectly (leaf gini mean = 0) before a depth of 10.


### -t --trees

```go run main.go -t 1000```

Provide addtional argument SIZE integer, number of trees in the forest. The default is 100 trees.


### -s --seed

```go run main.go -s 42```

Provide addtional argument SEED integer for randomization.

This seeds the pseudo-randomization of shuffling & splitting of data.
Thus a forest & set of predictions can be replicated exactly with a given seed.
The default seed is the current time.


### -f --forest

```go run main.go -f -d 1 -s 2```

Print forest, trees, node by node.


### -q --quiet

```go run main.go -q```

Don't print seed or forest statistics.


### data.csv

```go run main.go data.csv```
