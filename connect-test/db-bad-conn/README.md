開大量 gorutine 測試DB連線
===

## 說明
為了探查偶爾發生的bad connection 問題

若業務上會根據不同gorutine而使用DB連線>> 可能同時間大量同時連線DB

db.SetConnMaxLifetime() 若設太短絕對會發生慘案，建議是不設，除非是非常非常建議每個request的有效期限

但是當你會用開不同gorutine，分別處理DB連線的狀況，那麼基本上這些goruitne處理的東西基本上彼此獨立，不注重前後關係。



## 備註
db.SetConnMaxLifetime() 設到3s，在百萬次query就有測試到3次bad connection的狀況。

當連線達到max的時候，後續的request都是等待既有的連線回來，連線是隨機給到哪個等待的request，難免有衰鬼出現，一直拿不到。

##
SetConnMaxLifetime()可以設，300s之類的，避免被DB server因為timeout關閉。
db.SetMaxIdleConns()設個30個，50個就很多了，設太多被server不明原因大量關閉時會很麻煩。
