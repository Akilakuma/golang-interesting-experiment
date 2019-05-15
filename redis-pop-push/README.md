測試golang對於redis的push 和 pop處理是否有異常
====

＃＃說明
1. 本地需要安裝redis，6379 port
2. 先執行 ./rtest push
3. 後執行 ./rtest pop
4. 測試的迴圈，加入『思考時間』
5. 若有連續兩次pop出同樣的內容則有問題


＃＃結果
執行1750萬次，未出現錯誤

