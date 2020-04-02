// 实现一个函数LazyMan

// LazyMan('hanz')   
/**
 * 输出 I'm hanz 
 */
// LazyMan('hanz').sleep(5).eat('dinner')
/**
 * 输出 I'm hanz 
 * 等待5s 
 * 输出 Wake Up After 5
 * 输出 eat dinner
 */
//   LazyMan('hanz').eat('supper').eat('dinner')
/**
 * 输出 I'm hanz 
 * 输出 eat supper
 * 输出 eat dinner
 */
// LazyMan('hanz').sleepFirst(4).eat('dinner')
/**
 * 等到4s
 * 输出 Wake Up After 4
 * 输出 I'm hanz 
 * 输出 eat dinner
 */

function sleep(time) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            console.log(`Wake Up After ${time}s`)
            resolve()
        }, ~~time * 1000)
    })
}

function LazyMan(name) {
    let normalActions = []
    let sleepFirstActions = []
    name && normalActions.push({ type: 'name', name })
    setTimeout(async () => {
        for (let i = 0; i < sleepFirstActions.length; i++) {
            await sleep(sleepFirstActions[i])
        }
        for (let i = 0; i < normalActions.length; i++) {
            switch (normalActions[i].type) {
                case 'name':
                    console.log(`I’m ${normalActions[i].name}`)
                    break
                case 'eat':
                    console.log(`eat ${normalActions[i].food}`)
                    break
                case 'sleep':
                    await sleep(normalActions[i].time)
                    break
            }
        }
    })
    return {
        eat(food) {
            normalActions.push({
                type: 'eat',
                food
            })
            return this
        },
        sleep(time) {
            normalActions.push({
                type: 'sleep',
                time
            })
            return this
        },
        sleepFirst(time) {
            sleepFirstActions.push(time)
            return this
        }
    }
}

// let a = LazyMan('hanz')
// LazyMan('hanz').sleep(5).eat('dinner')
LazyMan('hanz').eat('supper').eat('dinner')
  // LazyMan('hanz').sleepFirst(4).eat('dinner')
