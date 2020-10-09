别问 问就是不知道用来干啥的
用法:

```
获取ip
▶ cat domain.txt | domain2ip | awk '{print $2}' | sort -u | tee ips.txt
```

```
计算ip范围
▶ for i in $(cat ips.txt);do ipcalc -n $i/24 | grep Network |awk '{print ($2)}' | grep -v "127.0.0.0" |grep -v "^10."|grep -v "192.168."| grep -v "172." ;done | sort -u | tee ips_range.txt
```

安装:

```
▶ go get -u -v github.com/flag007/domain2ip
```
