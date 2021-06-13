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

<img src="https://github.com/dfinnis/Random_Forest/blob/master/img/default.png" width="400">


## Flags

### -d --depth

Provide addtional argument DEPTH integer, maximum tree depth.

Default depth is 42, in essence infinite as a tree categorizes samples perfectly (leaf gini mean = 0) before a depth of 10.
Random forests with infinite depth overfit the training set, as shown in the deafult example above with perfect accuracy on the training set.
To avoid overfitting, & increase accuracy for the test set, a lower depth can be set. For example:

```go run main.go -d 5```

<img src="https://github.com/dfinnis/Random_Forest/blob/master/img/d.png" width="400">


### -t --trees

Provide addtional argument SIZE integer, number of trees in the forest. The default is 100 trees. To create a forest with 1000 trees: <br>
```go run main.go -t 1000```


### -s --seed

Provide addtional argument SEED integer for randomization. for example: <br>
```go run main.go -s 42```

This seeds the pseudo-randomization of shuffling & splitting of data.
Thus a forest & set of predictions can be replicated exactly with a given seed.
The default seed is the current time.


### -f --forest

Print forest, all trees, recursively node by node. Let's see a simple example with 2 trees of depth 1:

```go run main.go -f -d 1 -t 2```

<img src="https://github.com/dfinnis/Random_Forest/blob/master/img/f.png" width="370">


### -q --quiet

Don't print seed or forest statistics.

```go run main.go -q```


### data.csv

Any non-flag argument will be read as data path. The default data path is data.csv.

```go run main.go data.csv```
