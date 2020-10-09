别问 问就是不知道用来干啥的
用法:

```
获取ip
▶ cat domain.txt | domain2ip | awk '{print $2}' >> ips.txt
```

```
计算ip范围
▶ for i in $(cat ips.txt);do ipcalc -n $i/24 | grep Network |awk '{print ($2)}' | grep -v "^10."|grep -v "192.168."| grep -v "172." >> ips_range.txt;done
```

```
测试
▶ for i in $(cat ips.txt);do ipcalc -n $i/24 | grep Network |awk '{print ($2)}' | grep -v "^10."|grep -v "192.168."| grep -v "172." ;done
```
安装:

```
▶ go get -u -v github.com/flag007/domain2ip
```
