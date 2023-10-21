### FXMONEY

for now this package will more focused on "stuff" related to fx market not the money itself.  
will start with this structure:

```go
type Price struct {
Amount   int
Currency string

coefficient int
}
```

`Amount` is the amount of money in the POINTS (not pips) - smallest unit of the currency at fx market.
for example for USD will be 1/100000 part of dollar, for JPY is 1/1000 part of yen.

`Currency` is the currency of the price.

`coefficient` is the coefficient for the currency. will be set automatically based on the currency.
