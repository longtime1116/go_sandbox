* Go Module を使って開発していく
  * https://github.com/golang/go/wiki/Modules
  * Go Module だとローカルのパッケージをvscodeで補完できない！
    * gopls(Go の Language Server)を導入すれば良い
    * https://mattn.kaoriya.net/software/lang/go/20181217000056.htm
    * https://qiita.com/ryysud/items/1cf66ee4363aec22394a
    * https://github.com/golang/tools/blob/gopls/v0.3.1/gopls/doc/vscode.md
* ↓で擬似的にホットリロード的な機能を実現できる...
  * watch -n1 ./main
  * watch -n1 go build -o main
