# testmod
just for test how to use go mod

url:
https://www.jianshu.com/p/f97b090a8d53

# 总结
对着上面url做了一下，之前一直不理解，在公司import包时路径和url是对不上的，但问题是go mod又可以正确地下载包。上面提到在包路径下加/v2标示版本也能正常运行，于是有点困惑，现在总结下。
go module会把文件下载在$GOPATH/mod下。调用的时候，其实package名字并不重要，重要的是import，import指示了如何在mod下寻找对应的文件夹，找到了文件夹就会把该文件夹下的所有符号都静态编译到import的模块里。我自己测试了下，把import路径下的库名改成别的字符串，也能正常运转。

找包流程：
- 编译器获取import，通过```go ls-remote -q origin```命令检查是否有这个远程仓库存在
- 不存在报错，存在则看$GOPATH/mod下是否有对应着import路径的目录（注意import的时候会有所不同，比如包名会加上版本）
- 然后编译器会像c那样，直接将该包下的文件都加入编译，并生成cache文件加速下一次查找
- 如果本地没有，则会从git上拉取。

(我为什么要花这么大力气搞懂这个流程 晕...)

