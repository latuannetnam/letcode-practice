// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 2. Add Two Numbers
// Letcode link: https://leetcode.com/problems/add-two-numbers/
// Level: medium
// Topics: Link list, Math, Recursion
// Problem detail:
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
use std::convert::TryInto;
use std::collections::HashMap;
pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut seen = HashMap::new();
    for (i, num) in nums.iter().enumerate() {
        let complement = target - num;
        // let complement: i32 = i32::from(complement);
        if seen.contains_key(&complement) {
            // return vec![i, seen[&complement]];
            return vec![i.try_into().unwrap(), seen[&complement]];
        } else {
            // seen.insert(num, i);
            seen.insert(num, i.try_into().unwrap());
        }
    }
    return vec![];
}

pub fn two_sum_2(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut map = std::collections::HashMap::new();
    for (i, &num) in nums.iter().enumerate() {
        if let Some(&j) = map.get(&(target - num)) {
            return vec![j as i32, i as i32];
        }
        map.insert(num, i);
    }
    vec![]
}


pub fn test_two_sum() {
    let nums = vec![2, 7, 11, 15];
    let target = 9;
    let indices = two_sum_2(nums, target);
    println!("{:?}", indices);
}
