# submit_tools

## 项目由来
最近在做一些群内的人员管理工作，经常要下发表格，筛选出没提交的人。  
众所周知，懒是第一生产力，于是开发了一个可以根据人员信息表与实际提交信息，直接计算出未提交名单的工具。

## 简单介绍
### 人员信息表
csv 格式，每一行是一个人的信息。第一列是姓名，其余的都是别名列，列数任意。  
注意，每一列都应该能唯一标识该人，如QQ等，若提交信息出现了该QQ，则认为该人已提交。  
类似的，可以加上 Github ID、电话、邮箱等。不要使用非唯一别名。
### 提交信息
考虑到兼容性，采用了纯文本格式。主要一个人的任何一个别名出现在了提交信息字符串中，就认为该人已提交。
注：目前存在的 bug：如果一个人叫「张三」、另一个叫「张三丰」，那么只要后者提交了，前者会被误判为已提交。  
适当规避这个 bug 的方式是，使用较长的唯一 id。  

## 运行
* 使用 go mod, package 模式，选择 `desktop`  
* 在 `/desktop/static/list` 和 `/desktop/static/submission` 中分别建立人员信息文件与提交信息文件  
* 在 `/desktop/main.go` 中指定两个文件的文件名
* run

## features
* 结合 qq_bot
* 抽象出接口，增加可拓展性。为不同的信息提交分别实现接口，可以解决之前的误判 bug。