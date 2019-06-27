
實驗 fatal error: concurrent map iteration and map write 出現狀況
===

[參考issue sync: concurrent map iteration and map write on Map error](https://github.com/golang/go/issues/24112)


# 結論
1. 即使並非sync.Map，使用原本的map，當然也有concurrent map iteration and map write的問題發生。
1. 當有不同的gorutine，一個在對map做interation，一個在對map寫入，就會出事，當然兩邊都有加上lock可以避免錯誤，但是否是想要的請斟酌這樣使用。
