RPC 在本機環境測試實驗
======

## 說明
    線上環境有很小很小機率遇到連線建立，耗費超過預期之時間。

## 結論
    該實驗方式無法模擬出線上狀況，雖有突然建立連線會花費較多時間的狀況，但是另外的問題造成。

## 額外的實驗結果
    同時連線數量打爆server端的話，request不會進到server端程式，是由系統自動拒絕。
    一但有這樣的拒絕出現，會導致清場的現象，也許是系統保護機制，有一大段時間，server完全無收受任何request進來。

## 實驗手法
    挑選grpc、jsonrpc、rpc、包裝的zrpc 交叉在client和server的角色做使用。
    client 端以gorutine的方式，短時間大量發request到server 端
    到vendor/google.golang.org/grpc/server.go偷塞一些log程式碼做觀察


## server 收到的時候中間有大幅的真空狀況
![](https://i.imgur.com/6noRSaK.png)

## 錯誤訊息
tcp 127.0.0.1:7777: connect: connection reset by peer
