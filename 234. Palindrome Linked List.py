# Letcode problem: https://leetcode.com/problems/palindrome-linked-list/
# Level: easy
# Given the head of a singly linked list, return true if it is a palindrome.

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

    
class Solution:
    def head2number(self, head) -> int:
        node = head
        number = node.val
        print("number:", number)
        while node is not None:
            print("node:", node)            
            next_node = node.next
            if next_node is not None:
                number = number*10 + next_node.val
                print("number:", number)
                node = next_node
            else:
                break
        print("total number:", number)
        return number
    
    def reverse_number(self, num: int) -> int:
        reversed_num = 0
        while num != 0:
            digit = num % 10
            reversed_num = reversed_num * 10 + digit
            num //= 10
        return reversed_num
    
    def node2list(self, head):
        nums = []
        node = head
        while node is not None:
            nums.append(node.val)
            node = node.next
        return nums
    
    def isPalindrome(self, head) -> bool:
        nums =self.node2list(head)
        reversed_nums = nums[::-1]
        print('nums:', nums)
        print('r_nums:', reversed_nums)
        if nums == reversed_nums:
            return True
        
        
        return False

if __name__ == '__main__':
    nums = [4,8,2,0,5,0,1,6,3,0,9,0,6,5,3,4,2,7,3,8,8,1,0,6,0,1,0,4,0,6,3,2,6,0,2,0,9,0,5,5,5,7,9,7,7,0,6,2,8,0,6,7,4,1,7,5,0,8,1,2,3,9,1,5,9,8,3,9,2,0,0,4,0,2,8,7,4,9,6,8,9,3,8,8,3,9,8,6,9,4,7,8,2,0,4,0,0,2,9,3,8,9,5,1,9,3,2,1,8,0,5,7,1,4,7,6,0,8,2,6,0,7,7,9,7,5,5,5,0,9,0,2,0,6,2,3,6,0,4,0,1,0,6,0,1,8,8,3,7,2,4,3,5,6,0,9,0,3,6,1,0,5,0,2,8,4,0]
    head = ListNode(0)
    for num in nums:
        last_node = head.last_node()
        last_node.next = ListNode(num)

    solution = Solution()
    print(solution.isPalindrome(head))
    