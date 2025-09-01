# Recommendation Algorithm in Go

Collaborative filtering recommendation system in Go:
- Ranking algorithm using likes/dislikes or numeric ratings.
- This package can be used in any Go project or module.
- MIT license. ***Feel free to use this project***. ***Leave a star :star: or make a fork !***  

*If you found this project useful, consider making a donation to support the developer.*  

[![paypal](https://www.paypalobjects.com/pt_BR/BR/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate?hosted_button_id=S7FBV5N6ZTRXQ)  
[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate?hosted_button_id=PPDESEV98R8KS)

[![](https://github.com/tigoCaval/images/blob/main/web/recommend.gif)](https://github.com/tigoCaval/recommendation-go)

### Getting started
Starting with Go modules:

1. Install Go (tested with 1.20+)
2. Download package:  
```bash
go get github.com/tigoCaval/recommendation-go
```
Algorithms
* Ranking
* Euclidean
* SlopeOne

## Introduction
Recommend products using collaborative filtering:
Example

Simple demonstration of collaborative filtering:

```go
table := []recommendation.Transaction{
    {ProductID: "A", Score: 1, UserID: "John"},
    {ProductID: "B", Score: 1, UserID: "John"},
    {ProductID: "A", Score: 1, UserID: "Mary"},
    {ProductID: "B", Score: 0, UserID: "Mary"},
    {ProductID: "C", Score: 1, UserID: "Mary"},
}

client := recommendation.NewRecommend()

fmt.Println(client.Ranking(table,"John"))    // map[C:1]
fmt.Println(client.Euclidean(table,"John"))  // map[C:1]
fmt.Println(client.SlopeOne(table,"John"))   // map[C:1]



```
