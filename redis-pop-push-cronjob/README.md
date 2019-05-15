測試golang對於redis的push 和 pop處理是否有異常
====

＃＃說明
1. 本地需要安裝redis，6379 port
2. 先執行 ./rtest push
3. 後執行 ./rtest pop
4. 實際的環境因為是結合crontab，定時啟動處理，因此本次時間加入cronjob的部分
5. 若有連續兩次pop出同樣的內容則有問題


＃＃結果
每秒可試驗5000次，每分鐘 = 15萬次, 1小時 = 900萬次

