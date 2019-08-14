

def main():
    for index in range(10000):
        q = 0.91 + index/10000.0

        balance = 50 * 6144 * 5 * (1.0-q**855)/(1.0-q)
        print(q,"\t\t", balance)

        
if __name__ == '__main__':
    main()

#预挖210 0000
#pow2100 0000  总量 2310 0000
#平均一分钟出一个块
#50年挖完
#衰减周期和现在保持一致
#等比衰减