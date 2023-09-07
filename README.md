# Interest rates

Simple package to manage interest rates.


```go

    yearly, _ := interest.NewRate(0.1, 365)
    
    monthly, _, := yearly.Resample(30) 

    nominal, _ := yearly.NominalYearly()
```
