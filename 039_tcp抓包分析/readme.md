[toc]

# 使用 tcpdump 对 tcp 进行抓包分析

## 1：编写 tcp-server 代码

## 2：编写 tcp-client 代码

## 3：使用 tcpdump 进行抓包分析

- 现在，我们通过 tcpdump 来抓取上文 TCP 客户端与服务器通信全过程数据。

### 3.1 `sudo tcpdump -S -nn -vvv -i lo port 8000`

- `lo`是指定的网卡

### 3.2 tcp-server start

### 3.3 tcp-client start

### 3.4 tcpdum 抓包分析

- Flags []，其中 [S] 代表 SYN 包，[F] 代表 FIN，[.] 代表对应的 ACK 包。例如 [S.] 代表 SYN-ACK，[F.] 代表 FIN-ACK。
- 可以很明显看出 TCP 通信的全过程如下图所示:

```bash
# tcpdump 命令
$ sudo tcpdump -S -nn -vvv -i lo port 8000

# 开始抓包
tcpdump: listening on lo, link-type EN10MB (Ethernet), capture size 262144 bytes

# tcp 3次握手
10:45:31.713843 IP (tos 0x0, ttl 64, id 62915, offset 0, flags [DF], proto TCP (6), length 60)
    127.0.0.1.59322 > 127.0.0.1.8000: Flags [S], cksum 0xfe30 (incorrect -> 0x947e), seq 1960817469, win 65495, options [mss 65495,sackOK,TS val 1395669307 ecr 0,nop,wscale 7], length 0
10:45:31.713855 IP (tos 0x0, ttl 64, id 0, offset 0, flags [DF], proto TCP (6), length 60)
    127.0.0.1.8000 > 127.0.0.1.59322: Flags [S.], cksum 0xfe30 (incorrect -> 0xb72c), seq 645999200, ack 1960817470, win 65483, options [mss 65495,sackOK,TS val 1395669307 ecr 1395669307,nop,wscale 7], length 0
10:45:31.713863 IP (tos 0x0, ttl 64, id 62916, offset 0, flags [DF], proto TCP (6), length 52)
    127.0.0.1.59322 > 127.0.0.1.8000: Flags [.], cksum 0xfe28 (incorrect -> 0xdde8), seq 1960817470, ack 645999201, win 512, options [nop,nop,TS val 1395669307 ecr 1395669307], length 0

# client 发送数据
10:45:31.713962 IP (tos 0x0, ttl 64, id 62917, offset 0, flags [DF], proto TCP (6), length 90)
    127.0.0.1.59322 > 127.0.0.1.8000: Flags [P.], cksum 0xfe4e (incorrect -> 0x858b), seq 1960817470:1960817508, ack 645999201, win 512, options [nop,nop,TS val 1395669307 ecr 1395669307], length 38
10:45:31.713968 IP (tos 0x0, ttl 64, id 38181, offset 0, flags [DF], proto TCP (6), length 52)
    127.0.0.1.8000 > 127.0.0.1.59322: Flags [.], cksum 0xfe28 (incorrect -> 0xddc2), seq 645999201, ack 1960817508, win 512, options [nop,nop,TS val 1395669307 ecr 1395669307], length 0

# server 发送数据
10:45:31.717369 IP (tos 0x0, ttl 64, id 38182, offset 0, flags [DF], proto TCP (6), length 1076)
    127.0.0.1.8000 > 127.0.0.1.59322: Flags [P.], cksum 0x0229 (incorrect -> 0x8188), seq 645999201:646000225, ack 1960817508, win 512, options [nop,nop,TS val 1395669310 ecr 1395669307], length 1024
10:45:31.717421 IP (tos 0x0, ttl 64, id 62918, offset 0, flags [DF], proto TCP (6), length 52)
    127.0.0.1.59322 > 127.0.0.1.8000: Flags [.], cksum 0xfe28 (incorrect -> 0xd9c4), seq 1960817508, ack 646000225, win 504, options [nop,nop,TS val 1395669310 ecr 1395669310], length 0

# tcp  4次挥手
10:45:41.717759 IP (tos 0x0, ttl 64, id 62919, offset 0, flags [DF], proto TCP (6), length 52)
    127.0.0.1.59322 > 127.0.0.1.8000: Flags [F.], cksum 0xfe28 (incorrect -> 0xb2aa), seq 1960817508, ack 646000225, win 512, options [nop,nop,TS val 1395679311 ecr 1395669310], length 0
10:45:41.762182 IP (tos 0x0, ttl 64, id 38183, offset 0, flags [DF], proto TCP (6), length 52)
    127.0.0.1.8000 > 127.0.0.1.59322: Flags [.], cksum 0xfe28 (incorrect -> 0x8b6d), seq 646000225, ack 1960817509, win 512, options [nop,nop,TS val 1395679355 ecr 1395679311], length 0
10:45:56.772183 IP (tos 0x0, ttl 64, id 38184, offset 0, flags [DF], proto TCP (6), length 52)
    127.0.0.1.8000 > 127.0.0.1.59322: Flags [.], cksum 0xfe28 (incorrect -> 0x50cc), seq 646000224, ack 1960817509, win 512, options [nop,nop,TS val 1395694365 ecr 1395679311], length 0
10:45:56.772199 IP (tos 0x0, ttl 64, id 0, offset 0, flags [DF], proto TCP (6), length 52)
    127.0.0.1.59322 > 127.0.0.1.8000: Flags [.], cksum 0x509f (correct), seq 1960817509, ack 646000225, win 512, options [nop,nop,TS val 1395694365 ecr 1395679355], length 0


^C
17 packets captured
34 packets received by filter
0 packets dropped by kernel


```
