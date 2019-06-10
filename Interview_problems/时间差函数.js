'use strict';
const assert = require('assert')
/*
编写timeAgo(t1,t2) 函数
* 不用考虑闰年情况
* 考虑输入可能是数值型或者日期型
* 请尽可能少用if-else方式完成
* 完成以下应该的输出

1.考察代码的书写规范
2.考察解决问题的逻辑
3.考察代码质量
4.考察综合解决问题及延展的能力

*/
const timeAgo = function (t1, t2) { //两个时间差 中文显示函数
    t1 = new Date(t1).getTime()
    t2 = new Date(t2).getTime()
    let temp = (t2 - t1) / 1000, prefix, unit
    const unitMap = [[1, '秒'], [60, '分钟'], [60 * 60, '小时'], [60 * 60 * 24, '天'], [60 * 60 * 24 * 30, '个月'], [60 * 60 * 24 * 365, '年']]
    unit = unitMap.findIndex(el => el[0] > Math.abs(temp))
    unit = unit == -1 ? 5 : unit - 1
    prefix = Math.abs(~~(temp / unitMap[unit][0]))
    return `${prefix}${unitMap[unit][1]}${temp > 0 ? '前' : '后'}`
};



assert.strictEqual(timeAgo('2016-1-1', '2017-2-1'), '1年前')
assert.strictEqual(timeAgo('2016-1-1', '2016-3-1'), '2个月前')
assert.strictEqual(timeAgo('2016-1-1', '2016-1-16'), '15天前')
assert.strictEqual(timeAgo('2016-1-1', '2016-1-1 1:13:01'), '1小时前')
assert.strictEqual(timeAgo('2016-1-1', '2016-1-1 0:13:01'), '13分钟前')
assert.strictEqual(timeAgo('2016-1-1', '2016-1-1 0:0:50'), '50秒前')
assert.strictEqual(timeAgo('2018-1-1', '2017-1-1'), '1年后')
assert.strictEqual(timeAgo('2016-3-1', '2016-1-1'), '2个月后')
assert.strictEqual(timeAgo('2016-1-16', '2016-1-1'), '15天后')
assert.strictEqual(timeAgo('2016-1-1 1:13:01', '2016-1-1'), '1小时后')
assert.strictEqual(timeAgo('2016-1-1 0:13:01', '2016-1-1'), '13分钟后')
assert.strictEqual(timeAgo('2016-1-1 0:0:50', '2016-1-1'), '50秒后')
assert.strictEqual(timeAgo(new Date('2016-1-1'), new Date('2016-1-1 0:0:50')), '50秒前')
assert.strictEqual(timeAgo(+new Date('2016-1-1'), +new Date('2016-1-1 0:0:50')), '50秒前')
assert.strictEqual(timeAgo((new Date('2016-1-1 0:0:50')).toISOString(), +new Date('2016-1-1')), '50秒后')

// /*
// 以下是单元测试不需要修改
// */
// QUnit.module('时间差函数测试', function () {
//     QUnit.test('timeAgo', function (assert) {
//         assert.strictEqual(timeAgo('2016-1-1', '2017-2-1'), '1年前')
//         assert.strictEqual(timeAgo('2016-1-1', '2016-3-1'), '2个月前')
//         assert.strictEqual(timeAgo('2016-1-1', '2016-1-16'), '15天前')
//         assert.strictEqual(timeAgo('2016-1-1', '2016-1-1 1:13:01'), '1小时前')
//         assert.strictEqual(timeAgo('2016-1-1', '2016-1-1 0:13:01'), '13分钟前')
//         assert.strictEqual(timeAgo('2016-1-1', '2016-1-1 0:0:50'), '50秒前')
//         assert.strictEqual(timeAgo('2018-1-1', '2017-1-1'), '1年后')
//         assert.strictEqual(timeAgo('2016-3-1', '2016-1-1'), '2个月后')
//         assert.strictEqual(timeAgo('2016-1-16', '2016-1-1'), '15天后')
//         assert.strictEqual(timeAgo('2016-1-1 1:13:01', '2016-1-1'), '1小时后')
//         assert.strictEqual(timeAgo('2016-1-1 0:13:01', '2016-1-1'), '13分钟后')
//         assert.strictEqual(timeAgo('2016-1-1 0:0:50', '2016-1-1'), '50秒后')
//         assert.strictEqual(timeAgo(new Date('2016-1-1'), new Date('2016-1-1 0:0:50')), '50秒前')
//         assert.strictEqual(timeAgo(+new Date('2016-1-1'), +new Date('2016-1-1 0:0:50')), '50秒前')
//         assert.strictEqual(timeAgo((new Date('2016-1-1 0:0:50')).toISOString(), +new Date('2016-1-1')), '50秒后')
//     })
// })