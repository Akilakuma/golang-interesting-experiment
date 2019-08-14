開大量 gorutine 測試DB連線
===

## 說明
為了探查偶爾發生的bad connection 問題

若業務上會根據不同gorutine而使用DB連線>> 可能同時間大量同時連線DB

db.SetConnMaxLifetime() 若設太短絕對會發生慘案，建議是不設，除非是非常非常建議每個request的有效期限

但是當你會用開不同gorutine，分別處理DB連線的狀況，那麼基本上這些goruitne處理的東西基本上彼此獨立，不注重前後關係。

那麼這個SetConnMaxLifetime，就非常沒必要

## 備註
db.SetConnMaxLifetime() 設到3s，在百萬次query就有測試到3次bad connection的狀況。

當連線達到max的時候，因為一但空出的連線可以使用的request是隨機挑選，難免有雖鬼出現。
