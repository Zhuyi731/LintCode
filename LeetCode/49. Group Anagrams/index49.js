// Given an array of strings, group anagrams together.

// Example:

// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]
// Note:

// All inputs will be in lowercase.
// The order of your output does not matter.

/**
 * @param {string[]} strs
 * @return {string[][]}
 */
var groupAnagrams = function(strs) {
    let hash = {}

    strs.forEach(str=>{
        let hashCode = str.split('').sort().join('')
        if(hash[hashCode]){
            hash[hashCode].push(str)
        }else{
            hash[hashCode] = [str]
        }
    })
    let ans = [] 
    for(let prop in hash){
        ans.push(hash[prop])
    }
    console.log(ans)
    return ans
};

groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"])