

def main():
    a0 = 20480 * 250000000000           #6144 * 5 / 1.5 = 20480
    q = 0.9943    #8982456140.350817 + 10e8
    unit = 100000000.0
#    for  index in range(100000):
#        q +=   index/1000000.0
#        total = a0 *(1) /(1-q)
#        print(q, "\t", total/unit)

    balance = a0 *(1-q **855) /(1-q)
    print(q,"\t\t", balance/100000000.0)
    year1 = a0 *(1-q **17) /(1-q)
    print(year1/unit)
    total = a0 *(1) /(1-q)
    print(total/unit)

    QTestDc = 100/101.0
    BTestDc = 31.2 * 6144 * (1-QTestDc **17) /(1-QTestDc)
    print(BTestDc)
    print(BTestDc)    #=24e8

        
if __name__ == '__main__':
    main()

#预挖10 0000 0000
#总量 8982456140.350817 + 10e8 = 99.82亿
#平均90秒出一个块
#50年挖完
#衰减周期和TestDc保持一致
#等比衰减
#第一年出币8.3亿
#TestDc第一年出币301.3万，乘以500 大约是15亿。