一个简单的活跃度计算方法
输入：
[]struct{
date time.Time
times int
}

问题：
1. sum的值域如何确定
2. 结果的合理性如何验证
3. 似乎曲线未必合理
4. 权重虽然简单，但很能反映本质

案例：
    1. 如果连续活跃了300天后，2个月都没活跃了，算比较活跃吧
    2. 如果刚刚天天活跃，连续持续了1个月，也不能算比较活跃吧
    3. 活跃的连续性需要考虑不，对于没活跃的那些天做减法？
        不能简单这么做，可以最近N天内的不活跃才做减法，N < 30 ; 实际上最近30天的积分应该比较高的，可达到30%+才行