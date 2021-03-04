# Supermarket

Supermarket provides basic checkout functionality which allows users to scan items and then get a total price. It also supports multibuy offers, such as buy 3 items for 130.

## Usage

Include Supermarket in your Go project by executing:

```
go get github.com/gspec/supermarket
```

Then create your supermarket and add some stock items:

```go
sm := new(supermarket.Store)
sm.LoadStock(map[rune]supermarket.Item{
	'A': {UnitPrice: 50},
	'B': {UnitPrice: 30},
	'C': {UnitPrice: 20},
	'D': {UnitPrice: 15},
})
```

Then add some multibuy offers:

```go
multibuyA, _ := supermarket.NewMultibuy('A', 3, 20)
multibuyB, _ := supermarket.NewMultibuy('B', 2, 15)
sm.LoadOffers([]supermarket.Discounter{multibuyA, multibuyB})
```

Create a checkout for your supermarket as follows:

```go
c := supermarket.NewCheckout(sm)
```

Then you're ready to scan some items and check the price:

```go
c.Scan('A')
c.Scan('B')
c.Scan('A')
c.Scan('A')
fmt.Printf("Total Price: %v", c.GetTotalPrice())
```
