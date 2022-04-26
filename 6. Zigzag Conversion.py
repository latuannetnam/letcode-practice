# Author: Le Anh Tuan (latuannetnam@gmail.com)
# Letcode problem: https://leetcode.com/problems/zigzag-conversion/
# Level: medium
# The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: 
# (you may want to display this pattern in a fixed font for better legibility)

# Row = 1
# Convert: "PAYPALISHIRING"

# Row = 2
# P Y A I H R N
# A P L S I I G
# Convert: "PYAIHRNAPLSIIG"

#  Row = 3
# P   A   H   N
# A P L S I I G
# Y   I   R
# Convert: "PAHNAPLSIIGYIR"

# Row = 4
# P     I    N
# A   L S  I G
# Y A   H R
# P     I
# Convert: "PINALSIGYAHRPI"
            
# Write the code that will take a string and make this conversion given a number of rows:

class Solution:
    def convert_2_rows(self, s: str, numRows: int, s_convert_arr):
        # Place char in row, col
        r_idx = 0
        c_idx = 0
        for char in s:
            s_convert_arr[r_idx][c_idx] = char
            r_idx += 1
            if r_idx == numRows:
                r_idx = 0
                c_idx +=1

    def convert(self, s: str, numRows: int) -> str:
        # Return orginal string if row=1
        if numRows == 1:
            return s

        # Init 2D list to store char of string
        numCols = len(s)//2 + 1
        if numCols < numRows:
            numCols = numRows

        s_convert = ""
        s_convert_arr = [['' for i in range(numCols)] for j in range(numRows)]
        if numRows == 2:
            self.convert_2_rows(s, numRows, s_convert_arr)
        else:
            # Place char in row, col
            r_idx = 0
            c_idx = 0
            full_row = True
            for char in s:
                print("row:{} col: {} {}".format(r_idx, c_idx, char))
                s_convert_arr[r_idx][c_idx] = char
                if full_row:
                    r_idx += 1
                    if r_idx == numRows:
                        full_row = False
                        r_idx -= 2
                        c_idx +=1
                else:
                    r_idx -= 1
                    c_idx +=1
                    if r_idx == 0:
                        full_row = True

        # Convert 2D list to string
        for row in s_convert_arr:
                print(row)
                s_convert += ''.join(row)
        
        
        return s_convert
        


if __name__ == '__main__':
    # s = "PAYPALISHIRING"
    s = "A"
    numRows = 2
    solution = Solution()
    print("s: {} rows:{} convert:{}".format(s, numRows, solution.convert(s, numRows)))
    
