# Letcode problem: https://leetcode.com/problems/add-two-numbers/
# Level: medium
# You are given two non-empty linked lists representing two non-negative integers. 
# The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
    def last_node(self):
        node = self
        while node is not None:
            next_node = node.next
            if next_node is not None:
                node = next_node
            else:
                break
        return node
    def nums2node(self, nums):
        first = True
        for num in nums:
            if first == True:
                self.val = num
                self.next = None
            else:
                node =self.last_node()
                node.next = ListNode(num)
    def num2node(self, num):
        node = self.next
        while num > 0:
            digit = num % 10
            # print('num:', num, " digit:", digit)
            if node is None:
                self.val = digit
                node = self
            else:
                node =self.last_node()
                node.next = ListNode(digit)
            num = num // 10
    
    def node2list(self):
        nums = [self.val]
        node = self.next
        while node is not None:
            nums.append(node.val)
            node = node.next
        return nums


class Solution:

    def reverse_number(self, num):
        reversed_num = 0
        while num != 0:
            digit = num % 10
            reversed_num = reversed_num * 10 + digit
            num //= 10
        return reversed_num

    def node2num(self, node):
        last_node = self.last_node(node)
        num = last_node.val
        node_bf = self.node_before(node, last_node)
        while node_bf is not None:
            num = num*10 + node_bf.val
            last_node = node_bf
            node_bf = self.node_before(node, last_node)
        return num
    
    def node_before(self, root , node):
        node_bf = root
        while node_bf.next is not None:
            if node_bf.next == node:
                return node_bf
            node_bf = node_bf.next
        return None
        
    def last_node(self, node):
        last_node = node
        while last_node.next is not None:
            last_node = last_node.next
        return last_node   

    def num2node(self, num):
        node = None
        if num == 0:
            node = ListNode(0)
        while num > 0:
            digit = num % 10
            if node is None:
                node = ListNode(digit)
            else:
                last_node =self.last_node(node)
                last_node.next = ListNode(digit)
            num = num // 10
        return node

    def print_node(self, node):
        nums = [node.val]
        while node.next is not None:
            nums.append(node.next.val)
            node = node.next
        print("node num:", nums)

    def addTwoNumbers(self, l1, l2):
        num1 = self.node2num(l1)
        num2 = self.node2num(l2)
        print('num1:', num1, num2)
        total = num1 + num2
        l3 = self.num2node(total)
        print('Total:', total)
        self.print_node(l3)
        return l3

if __name__ == '__main__':
    head_1 = ListNode(0)
    # nums_11 = 342
    # [0,8,6,5,6,8,3,5,7]
# [6,7,8,0,8,5,8,9,7]
    nums_11 = 753865680
    head_1.num2node(nums_11)
    nums_12 = head_1.node2list()
    # nums_21 = 465
    nums_21 = 798580876
    head_2 = ListNode(0)
    head_2.num2node(nums_21)
    nums_22 = head_2.node2list()
    print('head_1:', nums_11, nums_12)
    print('head_2:', nums_21, nums_22)
    solution = Solution()
    solution.addTwoNumbers(head_1, head_2)


