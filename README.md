# Random_Forest

A Random Forest that predicts whether a tumor is malignant or benign, implemented from scratch.

I wanted to investigate the performance of a random forest in comparison with [a deep neural net I created](https://github.com/dfinnis/Multilayer_Perceptron). <br>
It turns out a random forest performs as well as a deep neural net for this task.

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


## test.sh

Prints for each *-d* depth the actual max & mean depth, accuracy for training & test sets, and mean training time. This shows the ideal forest depth to be around 5, where we reach peak test set accuracy (~96%). We start to overfit beyond a depth of 5, the training set accuracy continues to increase but the test set accuracy declines. The max depth never passes 9, by which point each leaf is pure (100% one diagnosis) and the training subset is classified perfectly. The mean depth never exceeds 5.5, by which point the majority of trees have classified their training subset perfectly.

```./test.sh```

<img src="https://github.com/dfinnis/Random_Forest/blob/master/img/test.png" width="400">
